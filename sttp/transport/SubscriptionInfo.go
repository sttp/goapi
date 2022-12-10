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

// SubscriptionInfo defines subscription related settings for a DataSubscriber instance.
type SubscriptionInfo struct {
	// FilterExpression defines the desired measurements for a subscription. Examples include:
	//
	// * Directly specified signal IDs (UUID values in string format):
	//     38A47B0-F10B-4143-9A0A-0DBC4FFEF1E8; E4BBFE6A-35BD-4E5B-92C9-11FF913E7877
	//
	// * Directly specified tag names:
	//     DOM_GPLAINS-BUS1:VH; TVA_SHELBY-BUS1:VH
	//
	// * Directly specified identifiers in "measurement key" format:
	//     PPA:15; STAT:20
	//
	// * A filter expression against a selection view:
	//     FILTER ActiveMeasurements WHERE Company='GPA' AND SignalType='FREQ'
	FilterExpression string

	// Throttled determines if data will be published using down-sampling.
	Throttled bool
	// PublishInterval defines the down-sampling publish interval to use when Throttled is true.
	PublishInterval float64

	// UdpDataChannel requests that a UDP channel be used for data publication.
	UdpDataChannel bool
	// DataChannelLocalPort defines the desired UDP port to use for publication.
	DataChannelLocalPort uint16

	// IncludeTime determines if time should be included in non-compressed, compact measurements.
	IncludeTime bool
	// EnableTimeReasonabilityCheck determines if publisher should perform time reasonability checks.
	// When enabled LagTime and LeadTime will be used to determine if a measurement timestamp is reasonable.
	EnableTimeReasonabilityCheck bool
	// LagTime defines the allowed past time deviation tolerance in seconds (can be sub-second).
	// Value is used to determine if a measurement timestamp is reasonable.
	// Only applicable when EnableTimeReasonabilityCheck is true.
	LagTime float64
	// LeadTime defines the allowed future time deviation tolerance in seconds (can be sub-second).
	// Value is used to determine if a measurement timestamp is reasonable.
	// Only applicable when EnableTimeReasonabilityCheck is true.
	LeadTime float64
	// UseLocalClockAsRealTime determines if publisher should use local clock as real time. If false,
	// the timestamp of the latest measurement will be used as real-time.
	// Only applicable when EnableTimeReasonabilityCheck is true.
	UseLocalClockAsRealTime bool
	// UseMillisecondResolution determines if time should be restricted to milliseconds in non-compressed, compact measurements.
	UseMillisecondResolution bool
	// RequestNaNValueFilter requests that the publisher filter, i.e., does not send, any NaN values.
	RequestNaNValueFilter bool

	// StartTime defines the start time for a requested temporal data playback, i.e., a historical subscription.
	// Simply by specifying a StartTime and StopTime, a subscription is considered a historical subscription.
	// Note that the publisher may not support historical subscriptions, in which case the subscribe will fail.
	StartTime string
	// StopTime defines the stop time for a requested temporal data playback, i.e., a historical subscription.
	// Simply by specifying a StartTime and StopTime, a subscription is considered a historical subscription.
	// Note that the publisher may not support historical subscriptions, in which case the subscribe will fail.
	StopTime string
	// ConstraintParameters defines any custom constraint parameters for a requested temporal data playback. This can
	// include parameters that may be needed to initiate, filter, or control historical data access.
	ConstraintParameters string
	// ProcessingInterval defines the initial playback speed, in milliseconds, for a requested temporal data playback.
	// With the exception of the values of -1 and 0, this value specifies the desired processing interval for data, i.e.,
	// basically a delay, or timer interval, over which to process data. A value of -1 means to use the default processing
	// interval while a value of 0 means to process data as fast as possible.
	ProcessingInterval int32

	// ExtraConnectionStringParameters defines any extra or custom connection string parameters that may be needed for a subscription.
	ExtraConnectionStringParameters string
}
