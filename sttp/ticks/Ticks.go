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
// Only bits 01 to 62 (0x3FFFFFFFFFFFFFFF) are used to represent the timestamp value. Bit 64 (0x8000000000000000) is used
// to denote leap second, i.e., second 60, where actual second value would remain at 59. Bit 63 is reserved and unset.
type Ticks uint64

// Min is the minimum value for Ticks. It represents UTC time 01/01/0001 00:00:00.000.
const Min Ticks = 0

// Max is the maximum value for Ticks. It represents UTC time 12/31/1999 11:59:59.999.
const Max Ticks = 3155378975999999999

// PerSecond is the number of Ticks that occur in a second.
const PerSecond Ticks = 10000000

// PerMillisecond is the number of Ticks that occur in a millisecond.
const PerMillisecond Ticks = PerSecond / 1000

// PerMicrosecond is the number of Ticks that occur in a microsecond.
const PerMicrosecond Ticks = PerSecond / 1000000

// PerMinute is the number of Ticks that occur in a minute.
const PerMinute Ticks = 60 * PerSecond

// PerHours is the number of Ticks that occur in an hour.
const PerHour Ticks = 60 * PerMinute

// PerDay is the number of Ticks that occur in a day.
const PerDay Ticks = 24 * PerHour

// LeapSecondFlag is the flag (64th bit) that marks a Ticks value as a leap second, i.e., second 60 (one beyond normal second 59).
const LeapSecondFlag Ticks = 1 << 63

// ReservedUTCFlag is the reserved flag (63rd bit) that should be unset when serializing and deserailing Ticks.
const ReservedUTCFlag Ticks = 1 << 62

// ValueMask defines all bits (bits 1 to 62) that make up the value porition of a Ticks that represent time.
const ValueMask Ticks = ^LeapSecondFlag & ^ReservedUTCFlag

// UnixBaseOffset is the Ticks representation of the Unix epcoh timestamp starting at January 1, 1970.
const UnixBaseOffset Ticks = 621355968000000000

// ToTime converts a Ticks value to standard Go Time value.
func ToTime(ticks Ticks) time.Time {
	return time.Unix(0, int64((ticks-UnixBaseOffset)&ValueMask)*100).UTC()
}

// FromTime converts a standard Go Time value to a Ticks value.
func FromTime(time time.Time) Ticks {
	return (Ticks(time.UnixNano()/100) + UnixBaseOffset) & ValueMask
}

// IsLeapSecond determines if the deserialized Ticks value represents a leap second, i.e., second 60.
func IsLeapSecond(ticks Ticks) bool {
	return (ticks & LeapSecondFlag) > 0
}

// SetLeapSecond flags a Ticks value to represent a leap second, i.e., second 60, before wire serialization.
func SetLeapSecond(ticks Ticks) Ticks {
	return ticks | LeapSecondFlag
}
