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
	"github.com/sttp/goapi/sttp/ticks"
)

// Function pointer types
type DispatcherFunction func(*DataSubscriber, []byte)
type MessageCallback func(*DataSubscriber, string)
type DataStartTimeCallback func(*DataSubscriber, ticks.Ticks)
type MetadataCallback func(*DataSubscriber, []byte)
type SubscriptionUpdatedCallback func(*DataSubscriber, *SignalIndexCache)
type NewMeasurementsCallback func(*DataSubscriber, []*Measurement)
type ConfigurationChangedCallback func(*DataSubscriber)
type ConnectionTerminatedCallback func(*DataSubscriber)

type DataSubscriber struct {
	subscriptionInfo SubscriptionInfo
}

func (ds *DataSubscriber) SetSubscriptionInfo(info SubscriptionInfo) {
	ds.subscriptionInfo = info
}
