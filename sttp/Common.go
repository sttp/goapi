//******************************************************************************************************
//  Common.go - Gbtc
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

package sttp

import (
	"time"

	"github.com/google/uuid"
)

// Ticks is a 64-bit integer used to designate time in STTP. The value represents the number of 100-nanosecond intervals
// that have elapsed since 12:00:00 midnight, January 1, 0001 UTC, Gregorian calendar. A single tick represents one hundred
// nanoseconds, or one ten-millionth of a second. There are 10,000 ticks in a millisecond and 10 million ticks in a second.
// Only bits 01 to 62 (0x3FFFFFFFFFFFFFFF) are used to represent the timestamp value. Bit 64 is used to denote leap second,
// i.e., second 60 (0x8000000000000000) – where actual second value would remain at 59. Bit 63 is reserved and always set.
type Ticks uint64

const TicksMin Ticks = 0                   // 01/01/0001 00:00:00.000
const TicksMax Ticks = 3155378975999999999 // 12/31/1999 11:59:59.999
const TicksPerSecond Ticks = 10000000
const TicksPerMillisecond Ticks = TicksPerSecond / 1000
const TicksPerMicrosecond Ticks = TicksPerSecond / 1000000
const TicksPerMinute Ticks = 60 * TicksPerSecond
const TicksPerHour Ticks = 60 * TicksPerMinute
const TicksPerDay Ticks = 24 * TicksPerHour
const TicksLeapSecondFlag Ticks = 1 << 63
const TicksReservedUTCFlag Ticks = 1 << 62
const TicksValueMask Ticks = ^TicksLeapSecondFlag & ^TicksReservedUTCFlag
const TicksUnixBaseOffset Ticks = 621355968000000000

func ToTime(ticks Ticks) time.Time {
	return time.Unix(0, int64((ticks-TicksUnixBaseOffset)&TicksValueMask)*100).UTC()
}

func ToTicks(time time.Time) Ticks {
	return (Ticks(time.UnixNano()/100) + TicksUnixBaseOffset) & TicksValueMask
}

func IsLeapSecond(ticks Ticks) bool {
	return (ticks & TicksLeapSecondFlag) > 0
}

func SetLeapSecond(ticks Ticks) Ticks {
	return ticks | TicksLeapSecondFlag
}

type Guid uuid.UUID

var EmptyGuid Guid = Guid(uuid.Nil)

func NewGuid() Guid {
	return Guid(uuid.New())
}

func ParseGuidFromBytes(data []byte, swapEndianness bool) Guid {
	swappedBytes := make([]byte, 16)
	var encodedBytes []byte

	if swapEndianness {
		var copy [8]byte

		for i := 0; i < 16; i++ {
			swappedBytes[i] = data[i]

			if i < 8 {
				copy[i] = swappedBytes[i]
			}
		}

		// Convert Microsoft encoding to RFC
		swappedBytes[3] = copy[0]
		swappedBytes[2] = copy[1]
		swappedBytes[1] = copy[2]
		swappedBytes[0] = copy[3]

		swappedBytes[4] = copy[5]
		swappedBytes[5] = copy[4]

		swappedBytes[6] = copy[7]
		swappedBytes[7] = copy[6]

		encodedBytes = swappedBytes
	} else {
		encodedBytes = data
	}

	guid, err := uuid.FromBytes(encodedBytes)

	if err == nil {
		return Guid(guid)
	}

	panic("Failed to parse Guid from bytes: " + err.Error())
}
