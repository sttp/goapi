//******************************************************************************************************
//  StateFlags.go - Gbtc
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

package stateflags

// StateFlags define the possible quality states of Measurement value.
type StateFlags uint32

const (
	// Normal defines a Measurement flag for a normal state.
	Normal StateFlags = 0x0
	// BadData defines a Measurement flag for a bad data state.
	BadData StateFlags = 0x1
	// SuspectData defines a Measurement flag for a suspect data state.
	SuspectData StateFlags = 0x2
	// OverRangeError defines a Measurement flag for a over range error, i.e., unreasonable high value.
	OverRangeError StateFlags = 0x4
	// UnderRangeError defines a Measurement flag for a under range error, i.e., unreasonable low value.
	UnderRangeError StateFlags = 0x8
	// AlarmHigh defines a Measurement flag for a alarm for high value.
	AlarmHigh StateFlags = 0x10
	// AlarmLow defines a Measurement flag for a alarm for low value.
	AlarmLow StateFlags = 0x20
	// WarningHigh defines a Measurement flag for a warning for high value.
	WarningHigh StateFlags = 0x40
	// WarningLow defines a Measurement flag for a warning for low value.
	WarningLow StateFlags = 0x80
	// FlatlineAlarm defines a Measurement flag for a alarm for flat-lined value, i.e., latched value test alarm.
	FlatlineAlarm StateFlags = 0x100
	// ComparisonAlarm defines a Measurement flag for a comparison alarm, i.e., outside threshold of comparison with a real-time value.
	ComparisonAlarm StateFlags = 0x200
	// ROCAlarm defines a Measurement flag for a rate-of-change alarm.
	ROCAlarm StateFlags = 0x400
	// ReceivedAsBad defines a Measurement flag for a bad value received.
	ReceivedAsBad StateFlags = 0x800
	// CalculatedValue defines a Measurement flag for a calculated value state.
	CalculatedValue StateFlags = 0x1000
	// CalculationError defines a Measurement flag for a calculation error with the value.
	CalculationError StateFlags = 0x2000
	// CalculationWarning defines a Measurement flag for a calculation warning with the value.
	CalculationWarning StateFlags = 0x4000
	// ReservedQualityFlag defines a Measurement flag for a reserved quality.
	ReservedQualityFlag StateFlags = 0x8000
	// BadTime defines a Measurement flag for a bad time state.
	BadTime StateFlags = 0x10000
	// SuspectTime defines a Measurement flag for a suspect time state.
	SuspectTime StateFlags = 0x20000
	// LateTimeAlarm defines a Measurement flag for a late time alarm.
	LateTimeAlarm StateFlags = 0x40000
	// FutureTimeAlarm defines a Measurement flag for a future time alarm.
	FutureTimeAlarm StateFlags = 0x80000
	// UpSampled defines a Measurement flag for a up-sampled state.
	UpSampled StateFlags = 0x100000
	// DownSampled defines a Measurement flag for a down-sampled state.
	DownSampled StateFlags = 0x200000
	// DiscardedValue defines a Measurement flag for a discarded value state.
	DiscardedValue StateFlags = 0x400000
	// ReservedTimeFlag defines a Measurement flag for a reserved time
	ReservedTimeFlag StateFlags = 0x800000
	// UserDefinedFlag1 defines a Measurement flag for user defined state 1.
	UserDefinedFlag1 StateFlags = 0x1000000
	// UserDefinedFlag2 defines a Measurement flag for user defined state 2.
	UserDefinedFlag2 StateFlags = 0x2000000
	// UserDefinedFlag3 defines a Measurement flag for user defined state 3.
	UserDefinedFlag3 StateFlags = 0x4000000
	// UserDefinedFlag4 defines a Measurement flag for user defined state 4.
	UserDefinedFlag4 StateFlags = 0x8000000
	// UserDefinedFlag5 defines a Measurement flag for user defined state 5.
	UserDefinedFlag5 StateFlags = 0x10000000
	// SystemError defines a Measurement flag for a system error state.
	SystemError StateFlags = 0x20000000
	// SystemWarning defines a Measurement flag for a system warning state.
	SystemWarning StateFlags = 0x40000000
	// MeasurementError defines a Measurement flag for an error state.
	MeasurementError StateFlags = 0x80000000
)
