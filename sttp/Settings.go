//******************************************************************************************************
//  Settings.go - Gbtc
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
//  09/29/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package sttp

// Settings defines the STTP subscription related settings.
type Settings struct {
	// Throttled determines if data will be published using down-sampling.
	Throttled bool
	// PublishInterval defines the down-sampling publish interval to use when Throttled is true.
	PublishInterval float64

	// UdpPort defines the desired UDP port to use for publication. Zero value
	UdpPort uint16

	// IncludeTime determines if time should be included in non-compressed, compact measurements.
	IncludeTime bool
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

	// ExtraConnectionStringParameters defines any extra custom connection string parameters that may be needed for a subscription.
	ExtraConnectionStringParameters string
}

// settingsDefaults define the default values for STTP subscription Settings.
var settingsDefaults = Settings{
	PublishInterval:    1.0,
	IncludeTime:        true,
	ProcessingInterval: -1,
}

// NewSettings creates a new Settings instance initalized with default values.
func NewSettings() *Settings {
	settings := settingsDefaults
	return &settings
}
