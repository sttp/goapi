//******************************************************************************************************
//  Ticks.go - Gbtc
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

package ticks

import (
	"time"
)

// Ticks is a 64-bit integer used to designate time in STTP. The value represents the number of 100-nanosecond intervals
// that have elapsed since 12:00:00 midnight, January 1, 0001 UTC, Gregorian calendar. A single tick represents one hundred
// nanoseconds, or one ten-millionth of a second. There are 10,000 ticks in a millisecond and 10 million ticks in a second.
// Only bits 01 to 62 (0x3FFFFFFFFFFFFFFF) are used to represent the timestamp value. Bit 64 is used to denote leap second,
// i.e., second 60 (0x8000000000000000) – where actual second value would remain at 59. Bit 63 is reserved and always set.
type Ticks uint64

const Min Ticks = 0                   // 01/01/0001 00:00:00.000
const Max Ticks = 3155378975999999999 // 12/31/1999 11:59:59.999
const PerSecond Ticks = 10000000
const PerMillisecond Ticks = PerSecond / 1000
const PerMicrosecond Ticks = PerSecond / 1000000
const PerMinute Ticks = 60 * PerSecond
const PerHour Ticks = 60 * PerMinute
const PerDay Ticks = 24 * PerHour
const LeapSecondFlag Ticks = 1 << 63
const ReservedUTCFlag Ticks = 1 << 62
const ValueMask Ticks = ^LeapSecondFlag & ^ReservedUTCFlag
const UnixBaseOffset Ticks = 621355968000000000

func ToTime(ticks Ticks) time.Time {
	return time.Unix(0, int64((ticks-UnixBaseOffset)&ValueMask)*100).UTC()
}

func FromTime(time time.Time) Ticks {
	return (Ticks(time.UnixNano()/100) + UnixBaseOffset) & ValueMask
}

func IsLeapSecond(ticks Ticks) bool {
	return (ticks & LeapSecondFlag) > 0
}

func SetLeapSecond(ticks Ticks) Ticks {
	return ticks | LeapSecondFlag
}
