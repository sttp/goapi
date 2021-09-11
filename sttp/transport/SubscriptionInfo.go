//******************************************************************************************************
//  SubscriptionInfo.go - Gbtc
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

type SubscriptionInfo struct {
	FilterExpression string

	// Down-sampling properties
	Throttled       bool
	PublishInterval float64

	// UDP channel properties
	UdpDataChannel       bool
	DataChannelLocalPort uint16

	// Compact measurement properties
	IncludeTime              bool
	LagTime                  float64
	LeadTime                 float64
	UseLocalClockAsRealTime  bool
	UseMillisecondResolution bool
	RequestNaNValueFilter    bool

	// Temporal playback properties
	StartTime            string
	StopTime             string
	ConstraintParameters string
	ProcessingInterval   int32

	ExtraConnectionStringParameters string
}
