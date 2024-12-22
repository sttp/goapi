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
// that have elapsed since 12:00:00 midnight, January 1, 0001 UTC, in the Gregorian calendar. A single tick represents 100ns.
// Only bits 01 to 62 (0x3FFFFFFFFFFFFFFF) are used to represent the timestamp value. Bit 64 (0x8000000000000000) is used
// to denote a leap second, and bit 63 (0x4000000000000000) is used to denote the leap second's direction, 0 for add, 1 for delete.
// Leap seconds are exposed, but are silently discarded upon conversion to Go or Unix timestamps.
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

// PerHour is the number of Ticks that occur in an hour.
const PerHour Ticks = 60 * PerMinute

// PerDay is the number of Ticks that occur in a day.
const PerDay Ticks = 24 * PerHour

// LeapSecondFlag is the flag (64th bit) that marks a Ticks value as a leap second, i.e., second 60 (one beyond normal second 59).
const LeapSecondFlag Ticks = 1 << 63

// LeapSecondDirection is the flag (63rd bit) that indicates if leap second is positive or negative; 0 for add, 1 for delete.
const LeapSecondDirection Ticks = 1 << 62

// ValueMask defines all bits (bits 1 to 62) that make up the value portion of a Ticks that represent time.
const ValueMask Ticks = ^LeapSecondFlag & ^LeapSecondDirection

// UnixBaseOffset is the Ticks representation of the Unix epoch timestamp starting at January 1, 1970.
const UnixBaseOffset Ticks = 621355968000000000

// TimeFormat is the standard time.Time format used for a Ticks value.
const TimeFormat string = "2006-01-02 15:04:05.999999999"

// ShortTimeFormat is the standard time.Time format used for showing just the timestamp portion of a Ticks value.
const ShortTimeFormat string = "15:04:05.999"

// TimestampValue gets the timestamp portion of the Ticks value, i.e.,
// the 62-bit time value excluding any leap second flags.
func (ticks Ticks) TimestampValue() int64 {
	return int64(ticks & ValueMask)
}

// ToTime converts a Ticks value to standard Go Time value.
func ToTime(ticks Ticks) time.Time {
	return time.Unix(0, int64((ticks-UnixBaseOffset)&ValueMask)*100).UTC()
}

// Converts a unix nanoseconds timestamp into a Ticks value.
func FromUnixNs(ns uint64) Ticks {
	return Ticks(ns / 100) + UnixBaseOffset
}

// FromTime converts a standard Go Time value to a Ticks value.
func FromTime(time time.Time) Ticks {
	return FromUnixNs(uint64(time.UnixNano()))
}

// IsLeapSecond determines if the deserialized Ticks value represents a leap second, i.e., second 60.
func IsLeapSecond(ticks Ticks) bool {
	return (ticks & LeapSecondFlag) > 0
}

// SetLeapSecond returns a copy of this Ticks value flagged to represent a leap second, i.e., second 60, before wire serialization.
func SetLeapSecond(ticks Ticks) Ticks {
	return ticks | LeapSecondFlag
}

// ApplyLeapSecond updates this Ticks value to represent a leap second, i.e., second 60, before wire serialization.
func (t *Ticks) ApplyLeapSecond() {
	*t |= LeapSecondFlag
}

// IsNegativeLeapSecond determines if the deserialized Ticks value represents a negative leap second, i.e., checks flag on second 58 to see if second 59 will be missing.
func IsNegativeLeapSecond(ticks Ticks) bool {
	return IsLeapSecond(ticks) && (ticks&LeapSecondDirection) > 0
}

// SetNegativeLeapSecond returns a copy of this Ticks value flagged to represent a negative leap second, i.e., sets flag on second 58 to mark that second 59 will be missing, before wire serialization.
func SetNegativeLeapSecond(ticks Ticks) Ticks {
	return ticks | LeapSecondFlag | LeapSecondDirection
}

// ApplyNegativeLeapSecond updates this Ticks value to represent a negative leap second, i.e., sets flag on second 58 to mark that second 59 will be missing, before wire serialization.
func (t *Ticks) ApplyNegativeLeapSecond() {
	*t |= LeapSecondFlag | LeapSecondDirection
}

// Now gets the current local time as a Ticks value.
func Now() Ticks {
	return FromTime(time.Now())
}

// UtcNow gets the current time in UTC as a Ticks value.
func UtcNow() Ticks {
	return FromTime(time.Now().UTC())
}

// ToTime converts a Ticks value to standard Go Time value.
func (t Ticks) ToTime() time.Time {
	return time.Unix(0, int64((t-UnixBaseOffset)&ValueMask)*100).UTC()
}

// IsLeapSecond determines if the deserialized Ticks value represents a leap second, i.e., second 60.
func (t Ticks) IsLeapSecond() bool {
	return IsLeapSecond(t)
}

// SetLeapSecond flags a Ticks value to represent a leap second, i.e., second 60, before wire serialization.
func (t Ticks) SetLeapSecond() Ticks {
	return SetLeapSecond(t)
}

// Converts the ticks value into a Unix nanoseconds timestamp
func (t Ticks) ToUnixNs() uint64 {
	return uint64(((t & ValueMask) - UnixBaseOffset) * 100)
}

// String returns the string form of a Ticks value, i.e., a standard date/time value. See TimeFormat.
func (t Ticks) String() string {
	return t.ToTime().Format(TimeFormat)
}

// ShortTime returns the short time string form of a Ticks value.
func (t Ticks) ShortTime() string {
	return t.ToTime().Format(ShortTimeFormat)
}
