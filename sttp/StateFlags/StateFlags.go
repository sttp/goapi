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

package StateFlags

type StateFlags uint32

const (
	// Defines normal state.
	Normal StateFlags = 0x0
	// Defines bad data state.
	BadData StateFlags = 0x1
	// Defines suspect data state.
	SuspectData StateFlags = 0x2
	// Defines over range error, i.e., unreasonable high value.
	OverRangeError StateFlags = 0x4
	// Defines under range error, i.e., unreasonable low value.
	UnderRangeError StateFlags = 0x8
	// Defines alarm for high value.
	AlarmHigh StateFlags = 0x10
	// Defines alarm for low value.
	AlarmLow StateFlags = 0x20
	// Defines warning for high value.
	WarningHigh StateFlags = 0x40
	// Defines warning for low value.
	WarningLow StateFlags = 0x80
	// Defines alarm for flat-lined value, i.e., latched value test alarm.
	FlatlineAlarm StateFlags = 0x100
	// Defines comparison alarm, i.e., outside threshold of comparison with a real-time value.
	ComparisonAlarm StateFlags = 0x200
	// Defines rate-of-change alarm.
	ROCAlarm StateFlags = 0x400
	// Defines bad value received.
	ReceivedAsBad StateFlags = 0x800
	// Defines calculated value state.
	CalculatedValue StateFlags = 0x1000
	// Defines calculation error with the value.
	CalculationError StateFlags = 0x2000
	// Defines calculation warning with the value.
	CalculationWarning StateFlags = 0x4000
	// Defines reserved quality flag.
	ReservedQualityFlag StateFlags = 0x8000
	// Defines bad time state.
	BadTime StateFlags = 0x10000
	// Defines suspect time state.
	SuspectTime StateFlags = 0x20000
	// Defines late time alarm.
	LateTimeAlarm StateFlags = 0x40000
	// Defines future time alarm.
	FutureTimeAlarm StateFlags = 0x80000
	// Defines up-sampled state.
	UpSampled StateFlags = 0x100000
	// Defines down-sampled state.
	DownSampled StateFlags = 0x200000
	// Defines discarded value state.
	DiscardedValue StateFlags = 0x400000
	// Defines reserved time flag.
	ReservedTimeFlag StateFlags = 0x800000
	// Defines user defined flag 1.
	UserDefinedFlag1 StateFlags = 0x1000000
	// Defines user defined flag 2.
	UserDefinedFlag2 StateFlags = 0x2000000
	// Defines user defined flag 3.
	UserDefinedFlag3 StateFlags = 0x4000000
	// Defines user defined flag 4.
	UserDefinedFlag4 StateFlags = 0x8000000
	// Defines user defined flag 5.
	UserDefinedFlag5 StateFlags = 0x10000000
	// Defines system error state.
	SystemError StateFlags = 0x20000000
	// Defines system warning state.
	SystemWarning StateFlags = 0x40000000
	// Defines measurement error flag.
	MeasurementError StateFlags = 0x80000000
)
