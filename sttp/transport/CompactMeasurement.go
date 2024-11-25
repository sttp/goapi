//******************************************************************************************************
//  CompactMeasurement.go - Gbtc
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
//  09/13/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package transport

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/sttp/goapi/sttp/ticks"
)

type compactStateFlagsEnum byte

// compactStateFlags constants represent each flag in the 8-bit compact measurement state flags.
var compactStateFlags = struct {
	DataRange       compactStateFlagsEnum
	DataQuality     compactStateFlagsEnum
	TimeQuality     compactStateFlagsEnum
	SystemIssue     compactStateFlagsEnum
	CalculatedValue compactStateFlagsEnum
	DiscardedValue  compactStateFlagsEnum
	BaseTimeOffset  compactStateFlagsEnum
	TimeIndex       compactStateFlagsEnum
}{
	DataRange:       0x01,
	DataQuality:     0x02,
	TimeQuality:     0x04,
	SystemIssue:     0x08,
	CalculatedValue: 0x10,
	DiscardedValue:  0x20,
	BaseTimeOffset:  0x40,
	TimeIndex:       0x80,
}

const (
	// These constants are masks used to set flags within the full 32-bit measurement state flags.
	dataRangeMask       StateFlagsEnum = 0x000000FC
	dataQualityMask     StateFlagsEnum = 0x0000EF03
	timeQualityMask     StateFlagsEnum = 0x00BF0000
	systemIssueMask     StateFlagsEnum = 0xE0000000
	calculatedValueMask StateFlagsEnum = 0x00001000
	discardedValueMask  StateFlagsEnum = 0x00400000
)

func (compactFlags compactStateFlagsEnum) mapToFullFlags() StateFlagsEnum {
	var fullFlags StateFlagsEnum

	if (compactFlags & compactStateFlags.DataRange) > 0 {
		fullFlags |= dataRangeMask
	}

	if (compactFlags & compactStateFlags.DataQuality) > 0 {
		fullFlags |= dataQualityMask
	}

	if (compactFlags & compactStateFlags.TimeQuality) > 0 {
		fullFlags |= timeQualityMask
	}

	if (compactFlags & compactStateFlags.SystemIssue) > 0 {
		fullFlags |= systemIssueMask
	}

	if (compactFlags & compactStateFlags.CalculatedValue) > 0 {
		fullFlags |= calculatedValueMask
	}

	if (compactFlags & compactStateFlags.DiscardedValue) > 0 {
		fullFlags |= discardedValueMask
	}

	return fullFlags
}

func (fullFlags StateFlagsEnum) mapToCompactFlags() compactStateFlagsEnum {
	var compactFlags compactStateFlagsEnum

	if (fullFlags & dataRangeMask) > 0 {
		compactFlags |= compactStateFlags.DataRange
	}

	if (fullFlags & dataQualityMask) > 0 {
		compactFlags |= compactStateFlags.DataQuality
	}

	if (fullFlags & timeQualityMask) > 0 {
		compactFlags |= compactStateFlags.TimeQuality
	}

	if (fullFlags & systemIssueMask) > 0 {
		compactFlags |= compactStateFlags.SystemIssue
	}

	if (fullFlags & calculatedValueMask) > 0 {
		compactFlags |= compactStateFlags.CalculatedValue
	}

	if (fullFlags & discardedValueMask) > 0 {
		compactFlags |= compactStateFlags.DiscardedValue
	}

	return compactFlags
}

// CompactMeasurement defines a measured value, in simple compact format, for transmission or reception in STTP.
type CompactMeasurement struct {
	Value                    float32
	Timestamp                ticks.Ticks
	SignalIndex              uint32
	Flags                    compactStateFlagsEnum
}

// Constructs a CompactMeasurement from the specified byte buffer; returns the measurement and the number of bytes occupied by this measurement.
func NewCompactMeasurement(includeTime, useMillisecondResolution bool, baseTimeOffsets *[2]int64, buffer []byte) (CompactMeasurement, int, error) {
	var cm CompactMeasurement

	if len(buffer) < 9 {
		return cm, 0, errors.New("not enough buffer available to deserialize compact measurement")
	}

	// Basic Compact Measurement Format:
	// 		Field:     Bytes:
	// 		--------   -------
	//		 Flags        1
	//		  ID          4
	//		 Value        4
	//		 [Time]    0/2/4/8

	cm.Flags = compactStateFlagsEnum(buffer[0])
	cm.SignalIndex = binary.BigEndian.Uint32(buffer[1:5])
	cm.Value = math.Float32frombits(binary.BigEndian.Uint32(buffer[5:9]))

	if !includeTime {
		return cm, 9, nil
	}

	if (cm.Flags & compactStateFlags.BaseTimeOffset) != 0 {
		timeIndex := (cm.Flags & compactStateFlags.TimeIndex) >> 7
		baseTimeOffset := baseTimeOffsets[timeIndex]
		if useMillisecondResolution {
			// Decode 2-byte millisecond offset timestamp
			if baseTimeOffset > 0 {
				cm.Timestamp = ticks.Ticks(baseTimeOffset + int64(binary.BigEndian.Uint16(buffer[9:11]))*int64(ticks.PerMillisecond))
			}
			return cm, 11, nil
		} else {
			// Decode 4-byte tick offset timestamp
			if baseTimeOffset > 0 {
				cm.Timestamp = ticks.Ticks(baseTimeOffset + int64(binary.BigEndian.Uint32(buffer[9:13])))
			}
			return cm, 13, nil
		}
	} else {
		// Decode 8-byte full fidelity timestamp
		// Note that only a full fidelity timestamp can carry leap second flags
		cm.Timestamp = ticks.Ticks(binary.BigEndian.Uint64(buffer[9:17]))
		return cm, 17, nil
	}
}

// Compute the full measurement from the compact representation
func (cm *CompactMeasurement) Expand(signalIndexCache *SignalIndexCache) Measurement {
	return Measurement{
		SignalID: signalIndexCache.SignalID(int32(cm.SignalIndex)),
		Timestamp: cm.Timestamp,
		Value: float64(cm.Value),
		Flags: cm.Flags.mapToFullFlags(),
	}
}

//// Serializes a CompactMeasurement to a byte buffer for publication to a DataSubscriber.
func (cm *CompactMeasurement) Marshal(b []byte) {
	b[0] = byte(cm.Flags)
	binary.BigEndian.PutUint32(b[1:], cm.SignalIndex)
	binary.BigEndian.PutUint32(b[5:], math.Float32bits(float32(cm.Value)))
	binary.BigEndian.PutUint64(b[9:], uint64(cm.Timestamp))
}

