//******************************************************************************************************
//  DataSubscriber.go - Gbtc
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
	"net"
	"strconv"
)

// DataSubscriber represents a client subscription for an STTP connection.
type DataSubscriber struct {
	subscriptionInfo SubscriptionInfo
	encoding         OperationalEncodingEnum
	connector        SubscriberConnector
	connected        bool
	connection       net.Conn

	// ConnectionTerminatedCallback is called when DataSubscriber terminates its connection.
	ConnectionTerminatedCallback func(*DataSubscriber)

	// AutoReconnectCallback is called when DataSubscriber automatically reconnects.
	AutoReconnectCallback func(*DataSubscriber)

	disposing bool
}

// SetSubscriptionInfo assigns the desired SubscriptionInfo for a DataSubscriber.
func (ds *DataSubscriber) SetSubscriptionInfo(info SubscriptionInfo) {
	ds.subscriptionInfo = info
}

// GetSubscriberConnector gets the SubscriberConnector for a DataSubscriber.
func (ds *DataSubscriber) GetSubscriberConnector() SubscriberConnector {
	return ds.connector
}

// DecodeString decodes an STTP string according to the defined operational modes.
func (ds *DataSubscriber) DecodeString(data []byte, length uint32) string {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if ds.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return string(data[:length])
}

func (ds *DataSubscriber) Connect(hostName string, port uint16, autoReconnecting bool) error {
	if ds.connected {
		panic("Subscriber is already connected; disconnect first")
	}

	var err error

	ds.connection, err = net.Dial("tcp", hostName+":"+strconv.Itoa(int(port)))

	return err
}
