//******************************************************************************************************
//  SubscriberConnector.go - Gbtc
//
//  Copyright © 2021, Grid Protection Alliance.  All Rights Reserved.
//
//  Licensed to the Grid Protection Alliance (GPA) under one or more contributor license agreements. See
//  the NOTICE file distributed with this work for additional information regarding copyright ownership.
//  The GPA licenses this file to you under the MIT License (MIT), the "License"; you may not use this
//  file except in compliance with the License. You may obtain a copy of the License at:
//
//      http://opensource.org/licenses/MIT
//
//  Unless agreed to in writing, the subject software distributed under the License is distributed on an
//  "AS-IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. Refer to the
//  License for the specific language governing permissions and limitations.
//
//  Code Modification History:
//  ----------------------------------------------------------------------------------------------------
//  09/09/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package transport

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sttp/goapi/sttp/thread"
	"github.com/tevino/abool/v2"
)

// SubscriberConnector represents a connector that will establish or automatically
// reestablish a connection from a DataSubscriber to a DataPublisher.
type SubscriberConnector struct {
	// ErrorMessageCallback is called when an error message should be logged.
	ErrorMessageCallback func(string)

	// ReconnectCallback is called when SubscriberConnector attempts to reconnect.
	ReconnectCallback func(*DataSubscriber)

	// Hostname is the DataPublisher DNS name or IP.
	Hostname string

	// Port it the TCP/IP listening port of the DataPublisher.
	Port uint16

	// MaxRetries defines the maximum number of times to retry a connection.
	// Set value to -1 to retry infinitely.
	MaxRetries int32

	// RetryInterval defines the base retry interval, in milliseconds. Retries will
	// exponentially back-off starting from this interval.
	RetryInterval int32

	// MaxRetryInterval defines the maximum retry interval, in milliseconds.
	MaxRetryInterval int32

	// AutoReconnect defines flag that determines if connections should be
	// automatically reattempted.
	AutoReconnect bool

	connectAttempt       int32
	connectionRefused    abool.AtomicBool
	cancel               abool.AtomicBool
	reconnectThread      *thread.Thread
	reconnectThreadMutex sync.Mutex
	waitTimer            *time.Timer
	waitTimerMutex       sync.Mutex

	assigningHandlerMutex sync.RWMutex
}

func (ds *DataSubscriber) autoReconnect() {
	sc := ds.connector

	if sc.cancel.IsSet() || ds.disposing.IsSet() {
		return
	}

	// Make sure to wait on any running reconnect to complete...
	sc.reconnectThreadMutex.Lock()
	reconnectThread := sc.reconnectThread
	sc.reconnectThreadMutex.Unlock()

	if reconnectThread != nil {
		reconnectThread.Join()
	}

	reconnectThread = thread.NewThread(func() {
		// Reset connection attempt counter if last attempt was not refused
		if sc.connectionRefused.IsNotSet() {
			sc.ResetConnection()
		}

		if sc.MaxRetries != -1 && sc.connectAttempt >= sc.MaxRetries {
			sc.dispatchErrorMessage("Maximum connection retries attempted. Auto-reconnect canceled.")
			return
		}

		sc.waitForRetry(ds)

		if sc.cancel.IsSet() || ds.disposing.IsSet() {
			return
		}

		if sc.connect(ds, true) == ConnectStatus.Canceled {
			return
		}

		// Notify the user that reconnect attempt was completed.
		sc.BeginCallbackSync()

		if sc.cancel.IsNotSet() && sc.ReconnectCallback != nil {
			sc.ReconnectCallback(ds)
		}

		sc.EndCallbackSync()
	})

	sc.reconnectThreadMutex.Lock()
	sc.reconnectThread = reconnectThread
	sc.reconnectThreadMutex.Unlock()

	reconnectThread.Start()
}

func (sc *SubscriberConnector) waitForRetry(subscriber *DataSubscriber) {
	// Apply exponential back-off algorithm for retry attempt delays
	var exponent float64

	if sc.connectAttempt > 13 {
		exponent = 12
	} else {
		exponent = float64(sc.connectAttempt - 1)
	}

	var retryInterval int32

	if sc.connectAttempt > 0 {
		retryInterval = sc.RetryInterval * int32(math.Pow(2, exponent))
	}

	if retryInterval > sc.MaxRetryInterval {
		retryInterval = sc.MaxRetryInterval
	}

	// Notify the user that we are attempting to reconnect.
	var message strings.Builder

	message.WriteString("Connection")

	if sc.connectAttempt > 0 {
		message.WriteString(" attempt ")
		message.WriteString(strconv.Itoa(int(sc.connectAttempt + 1)))
	}

	message.WriteString(" to \"")
	message.WriteString(sc.Hostname)
	message.WriteString(":")
	message.WriteString(strconv.Itoa(int(sc.Port)))
	message.WriteString("\" was terminated. ")

	if retryInterval > 0 {
		message.WriteString("Attempting to reconnect in ")
		message.WriteString(fmt.Sprintf("%.2f", float64(retryInterval)/1000.0))
		message.WriteString(" seconds...")
	} else {
		message.WriteString("Attempting to reconnect...")
	}

	sc.dispatchErrorMessage(message.String())

	waitTimer := time.NewTimer(time.Duration(retryInterval) * time.Millisecond)

	sc.waitTimerMutex.Lock()
	sc.waitTimer = waitTimer
	sc.waitTimerMutex.Unlock()

	<-waitTimer.C
}

// Connect initiates a connection sequence for a DataSubscriber
func (sc *SubscriberConnector) Connect(ds *DataSubscriber) ConnectStatusEnum {
	if sc.cancel.IsSet() {
		return ConnectStatus.Canceled
	}

	return sc.connect(ds, false)
}

func (sc *SubscriberConnector) connect(ds *DataSubscriber, autoReconnecting bool) ConnectStatusEnum {
	if sc.AutoReconnect {
		ds.AutoReconnectCallback = ds.autoReconnect
	}

	sc.cancel.UnSet()

	for ds.disposing.IsNotSet() {
		if sc.MaxRetries != -1 && sc.connectAttempt >= sc.MaxRetries {
			sc.dispatchErrorMessage("Maximum connection retries attempted. Auto-reconnect canceled.")
			break
		}

		sc.connectAttempt++

		if ds.disposing.IsSet() {
			return ConnectStatus.Canceled
		}

		err := ds.connect(sc.Hostname, sc.Port, autoReconnecting)

		if err == nil {
			break
		}

		if ds.disposing.IsNotSet() && sc.RetryInterval > 0 {
			autoReconnecting = true
			sc.waitForRetry(ds)

			if sc.cancel.IsSet() {
				return ConnectStatus.Canceled
			}
		}
	}

	if ds.disposing.IsSet() {
		return ConnectStatus.Canceled
	}

	if ds.IsConnected() {
		return ConnectStatus.Success
	}

	return ConnectStatus.Failed
}

// Cancel stops all current and future connection sequences.
func (sc *SubscriberConnector) Cancel() {
	sc.cancel.Set()

	sc.waitTimerMutex.Lock()
	waitTimer := sc.waitTimer
	sc.waitTimerMutex.Unlock()

	if waitTimer != nil {
		waitTimer.Stop()
	}

	sc.reconnectThreadMutex.Lock()
	reconnectThread := sc.reconnectThread
	sc.reconnectThreadMutex.Unlock()

	if reconnectThread != nil {
		reconnectThread.Join()
	}
}

// ResetConnection resets SubscriberConnector for a new connection.
func (sc *SubscriberConnector) ResetConnection() {
	sc.connectAttempt = 0
	sc.cancel.UnSet()
}

func (sc *SubscriberConnector) dispatchErrorMessage(message string) {
	sc.BeginCallbackSync()

	if sc.ErrorMessageCallback != nil {
		go sc.ErrorMessageCallback(message)
	}

	sc.EndCallbackSync()
}

// BeginCallbackAssignment informs SubscriberConnector that a callback change has been initiated.
func (sc *SubscriberConnector) BeginCallbackAssignment() {
	sc.assigningHandlerMutex.Lock()
}

// BeginCallbackSync begins a callback synchronization operation.
func (sc *SubscriberConnector) BeginCallbackSync() {
	sc.assigningHandlerMutex.RLock()
}

// EndCallbackSync ends a callback synchronization operation.
func (sc *SubscriberConnector) EndCallbackSync() {
	sc.assigningHandlerMutex.RUnlock()
}

// EndCallbackAssignment informs SubscriberConnector that a callback change has been completed.
func (sc *SubscriberConnector) EndCallbackAssignment() {
	sc.assigningHandlerMutex.Unlock()
}
