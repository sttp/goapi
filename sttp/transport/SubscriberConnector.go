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
	"time"
)

type ErrorMessageCallback func(*DataSubscriber, string)
type ReconnectCallback func(*DataSubscriber)

type SubscriberConnector struct {
	errorMessageCallback ErrorMessageCallback
	reconnectCallback    ReconnectCallback

	hostname string
	port     uint16
	timer    *time.Timer

	maxRetries        int32
	retryInterval     int32
	maxRetryInterval  int32
	connectAttempt    int32
	connectionRefused bool
	autoReconnect     bool
	cancel            bool
}

const ConnectSuccess int = 1
const ConnectFailed int = 0
const ConnectCanceled int = -1

// Registers a callback to provide error messages each time
// the subscriber fails to connect during a connection sequence.
func (sc *SubscriberConnector) RegisterErrorMessageCallback(errorMessageCallback ErrorMessageCallback) {
	sc.errorMessageCallback = errorMessageCallback
}

// Registers a callback to notify after an automatic reconnection attempt has been made.
// This callback will be called whether the connection was successful or not, so it is
// recommended to check the connected state of the subscriber using the IsConnected() method.
func (sc *SubscriberConnector) RegisterReconnectCallback(reconnectCallback ReconnectCallback) {
	sc.reconnectCallback = reconnectCallback
}

func (sc *SubscriberConnector) Connect(subscriber *DataSubscriber, info SubscriptionInfo) int {
	if sc.cancel {
		return ConnectCanceled
	}

	subscriber.SetSubscriptionInfo(info)
	return sc.connect(subscriber, false)
}

func (sc *SubscriberConnector) connect(subscriber *DataSubscriber, autoReconnecting bool) int {
	return ConnectSuccess
}
