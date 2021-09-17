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
	"math"

	"github.com/sttp/goapi/sttp/ticks"
)

type compactStateFlagsEnum byte

// These constants represent each flag in the 8-bit compact measurement state flags.
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

	fixedLength uint32 = 9
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
	Measurement
	signalIndexCache         *SignalIndexCache
	baseTimeOffsets          *[2]int64
	includeTime              bool
	useMillisecondResolution bool
	timeIndex                int32
	usingBaseTimeOffset      bool
}

// NewCompactMeasurement creates a new CompactMeasurement
func NewCompactMeasurement(signalIndexCache *SignalIndexCache, includeTime bool, useMillisecondResolution bool, baseTimeOffsets *[2]int64) CompactMeasurement {
	return CompactMeasurement{
		Measurement:              Measurement{},
		signalIndexCache:         signalIndexCache,
		baseTimeOffsets:          baseTimeOffsets,
		includeTime:              includeTime,
		useMillisecondResolution: useMillisecondResolution,
		timeIndex:                0,
		usingBaseTimeOffset:      false,
	}
}

// GetBinaryLength gets the binary byte length of a CompactMeasurement
func (cm *CompactMeasurement) GetBinaryLength() uint32 {
	var length uint32 = fixedLength

	if !cm.includeTime {
		return length
	}

	baseTimeOffset := cm.baseTimeOffsets[cm.timeIndex]

	if baseTimeOffset > 0 {
		if cm.includeTime {
			// See if timestamp will fit within space allowed for active base offset. We cache result so that post call
			// to GetBinaryLength, result will speed other subsequent parsing operations by not having to reevaluate.
			difference := cm.TicksValue() - baseTimeOffset

			if difference > 0 {
				if cm.useMillisecondResolution {
					cm.usingBaseTimeOffset = difference/int64(ticks.PerMillisecond) < math.MaxUint16
				} else {
					cm.usingBaseTimeOffset = difference < math.MaxUint32
				}
			} else {
				cm.usingBaseTimeOffset = false
			}

			if cm.usingBaseTimeOffset {
				if cm.useMillisecondResolution {
					length += 2 // Use two bytes for millisecond resolution timestamp with valid offset
				} else {
					length += 4 // Use four bytes for tick resolution timestamp with valid offset
				}
			} else {
				length += 8 // Use eight bytes for full fidelity time
			}
		}
	} else {
		// Use eight bytes for full fidelity time
		length += 8
	}

	return length
}

// GetTimestampC2 gets offset compressed millisecond-resolution 2-byte timestamp.
func (cm *CompactMeasurement) GetTimestampC2() uint16 {
	return uint16((cm.TicksValue() - cm.baseTimeOffsets[cm.timeIndex]) / int64(ticks.PerMillisecond))
}

// GetTimestampC4 gets offset compressed tick-resolution 4-byte timestamp.
func (cm *CompactMeasurement) GetTimestampC4() uint32 {
	return uint32(cm.TicksValue() - cm.baseTimeOffsets[cm.timeIndex])
}

// GetCompactStateFlags gets byte level compact state flags with encoded time index and base time offset bits.
func (cm *CompactMeasurement) GetCompactStateFlags() byte {
	// Encode compact state flags
	flags := cm.Flags.mapToCompactFlags()

	if cm.timeIndex != 0 {
		flags |= compactStateFlags.TimeIndex
	}

	if cm.usingBaseTimeOffset {
		flags |= compactStateFlags.BaseTimeOffset
	}

	return byte(flags)
}

// SetCompactStateFlags sets byte level compact state flags with encoded time index and base time offset bits.
func (cm *CompactMeasurement) SetCompactStateFlags(value byte) {
	// Decode compact state flags
	flags := compactStateFlagsEnum(value)

	cm.Flags = flags.mapToFullFlags()

	if (flags & compactStateFlags.TimeIndex) > 0 {
		cm.timeIndex = 1
	} else {
		cm.timeIndex = 0
	}

	cm.usingBaseTimeOffset = (flags & compactStateFlags.BaseTimeOffset) > 0
}

// GetRuntimeID gets the 4-byte run-time signal index for this measurement.
func (cm *CompactMeasurement) GetRuntimeID() int32 {
	return cm.signalIndexCache.SignalIndex(cm.SignalID)
}

// SetRuntimeID assigns CompactMeasurement SignalID (UUID) from the specified signalIndex.
func (cm *CompactMeasurement) SetRuntimeID(signalIndex int32) {
	cm.SignalID = cm.signalIndexCache.SignalID(signalIndex)
}

// Decode parses a CompactMeasurement from the specified byte buffer.
func (cm *CompactMeasurement) Decode(buffer []byte) int {
	if len(buffer) < 1 {
		panic("Not enough buffer available to deserialize compact measurement.")
	}

	// Basic Compact Measurement Format:
	// 		Field:     Bytes:
	// 		--------   -------
	//		 Flags        1
	//		  ID          4
	//		 Value        4
	//		 [Time]    0/2/4/8
	var index int

	// Decode state flags
	cm.SetCompactStateFlags(buffer[0])
	index++

	// Decode runtime ID
	cm.SetRuntimeID(int32(binary.BigEndian.Uint32(buffer[index:])))
	index += 4

	// Decode value
	cm.Value = float64(math.Float32frombits(binary.BigEndian.Uint32(buffer[index:])))
	index += 4

	if !cm.includeTime {
		return index
	}

	if cm.usingBaseTimeOffset {
		baseTimeOffset := cm.baseTimeOffsets[cm.timeIndex]

		if cm.useMillisecondResolution {
			// Decode 2-byte millisecond offset timestamp
			if baseTimeOffset > 0 {
				cm.Timestamp = ticks.Ticks(baseTimeOffset + int64(binary.BigEndian.Uint16(buffer[index:]))*int64(ticks.PerMillisecond))
			}
			index += 2
		} else {
			// Decode 4-byte tick offset timestamp
			if baseTimeOffset > 0 {
				cm.Timestamp = ticks.Ticks(baseTimeOffset + int64(binary.BigEndian.Uint32(buffer[index:])))
			}
			index += 4
		}
	} else {
		// Decode 8-byte full fidelity timestamp
		// Note that only a full fidelity timestamp can carry leap second flags
		cm.Timestamp = ticks.Ticks(binary.BigEndian.Uint64(buffer[index:]))
		index += 8
	}

	return index
}

//// Encode serializes a CompactMeasurement to a byte buffer for publication to a DataSubscriber.
//func (cm *CompactMeasurement) Encode(buffer []byte) {
//	// TODO: This will be needed by DataPublisher implementation
//}
