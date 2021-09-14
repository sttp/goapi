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
	"time"

	"github.com/sttp/goapi/sttp/thread"
)

// SubscriberConnector represents a connector that will establish to reestablish
// a connection from a DataSubscriber to a DataPublisher.
type SubscriberConnector struct {
	// ErrorMessageCallback is called when an error message should be logged.
	ErrorMessageCallback func(*DataSubscriber, string)

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

	connectAttempt    int32
	connectionRefused bool
	cancel            bool
	reconnectThread   *thread.Thread
}

// ConnectSuccess defines that a connection succeeded.
const ConnectSuccess int = 1

// ConnectFailed defines that a connection failed.
const ConnectFailed int = 0

// ConnectCanceled defines that a connection was cancelled.
const ConnectCanceled int = -1

func autoReconnect(subscriber *DataSubscriber) {
	connector := subscriber.GetSubscriberConnector()

	if connector.cancel || subscriber.disposing {
		return
	}

	// Make sure to wait on any running reconnect to complete...
	if connector.reconnectThread != nil {
		connector.reconnectThread.Join()
	}

	connector.reconnectThread = thread.NewThread(func() {
		// Reset connection attempt counter if last attempt was not refused
		if !connector.connectionRefused {
			connector.ResetConnection()
		}

		if connector.MaxRetries != -1 && connector.connectAttempt >= connector.MaxRetries {
			if connector.ErrorMessageCallback != nil {
				connector.ErrorMessageCallback(subscriber, "Maximum connection retries attempted. Auto-reconnect canceled.")
			}

			return
		}

		// Apply exponential back-off algorithm for retry attempt delays
		var exponent float64

		if connector.connectAttempt > 13 {
			exponent = 12
		} else {
			exponent = float64(connector.connectAttempt - 1)
		}

		var retryInterval int32

		if connector.connectAttempt > 0 {
			retryInterval = connector.RetryInterval * int32(math.Pow(2, exponent))
		}

		if retryInterval > connector.MaxRetryInterval {
			retryInterval = connector.MaxRetryInterval
		}

		// Notify the user that we are attempting to reconnect.
		if connector.ErrorMessageCallback != nil {
			var message strings.Builder

			message.WriteString("Connection")

			if connector.connectAttempt > 0 {
				message.WriteString(" attempt ")
				message.WriteString(strconv.Itoa(int(connector.connectAttempt + 1)))
			}

			message.WriteString(" to \"")
			message.WriteString(connector.Hostname)
			message.WriteString(":")
			message.WriteString(strconv.Itoa(int(connector.Port)))
			message.WriteString("\" was terminated. ")

			if retryInterval > 0 {
				message.WriteString("Attempting to reconnect in ")
				message.WriteString(fmt.Sprintf("%.2f", float64(connector.RetryInterval)/1000.0))
				message.WriteString(" seconds...")
			} else {
				message.WriteString("Attempting to reconnect...")
			}

			connector.ErrorMessageCallback(subscriber, message.String())
		}

		time.Sleep(time.Duration(retryInterval) * time.Millisecond)

		if connector.cancel || subscriber.disposing {
			return
		}

		if connector.connect(subscriber, true) == ConnectCanceled {
			return
		}

		// Notify the user that reconnect attempt was completed.
		if !connector.cancel && connector.ReconnectCallback != nil {
			connector.ReconnectCallback(subscriber)
		}
	})
}

// Connect initiates a connection sequence for a DataSubscriber for the specified SubscriptionInfo.
func (sc *SubscriberConnector) Connect(subscriber *DataSubscriber, info SubscriptionInfo) int {
	if sc.cancel {
		return ConnectCanceled
	}

	subscriber.SetSubscriptionInfo(info)
	return sc.connect(subscriber, false)
}

func (sc *SubscriberConnector) connect(subscriber *DataSubscriber, autoReconnecting bool) int {
	if sc.AutoReconnect {
		subscriber.AutoReconnectCallback = autoReconnect
	}

	sc.cancel = false

	//for !subscriber.disposing {

	//}

	return ConnectSuccess
}

// ResetConnection resets SubscriberConnector for a new connection.
func (sc *SubscriberConnector) ResetConnection() {
	sc.connectAttempt = 0
	sc.cancel = false
}
