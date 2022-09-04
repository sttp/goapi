//******************************************************************************************************
//  Decoder.go - Gbtc
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
//  12/02/2016 - Steven E. Chisholm
//       Generated original version of source code.
//  09/20/2021 - J. Ritchie Carroll
//       Migrated code to Go.
//
//******************************************************************************************************

package tssc

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Decoder is the decoder for the Time-Series Special Compression (TSSC) algorithm of STTP.
type Decoder struct {
	data         []byte
	position     int
	lastPosition int

	prevTimestamp1 int64
	prevTimestamp2 int64

	prevTimeDelta1 int64
	prevTimeDelta2 int64
	prevTimeDelta3 int64
	prevTimeDelta4 int64

	lastPoint *pointMetadata
	points    map[int32]*pointMetadata

	// The number of bits in m_bitStreamCache that are valid. 0 Means the bitstream is empty
	bitStreamCount int32

	// A cache of bits that need to be flushed to m_buffer when full. Bits filled starting from the right moving left
	bitStreamCache int32

	// SequenceNumber is the sequence used to synchronize encoding and decoding.
	SequenceNumber uint16
}

// NewDecoder creates a new TSSC decoder.
func NewDecoder(maxSignalIndex uint32) *Decoder {
	td := &Decoder{
		prevTimeDelta1: math.MaxInt64,
		prevTimeDelta2: math.MaxInt64,
		prevTimeDelta3: math.MaxInt64,
		prevTimeDelta4: math.MaxInt64,
		points:         make(map[int32]*pointMetadata, maxSignalIndex+1),
	}

	td.lastPoint = td.newPointMetadata()

	return td
}

func (td *Decoder) newPointMetadata() *pointMetadata {
	return newPointMetadata(nil, td.readBit, td.readBits5)
}

func (td *Decoder) bitStreamIsEmpty() bool {
	return td.bitStreamCount == 0
}

func (td *Decoder) clearBitStream() {
	td.bitStreamCount = 0
	td.bitStreamCache = 0
}

// SetBuffer assigns the working buffer to use for decoding measurements.
func (td *Decoder) SetBuffer(data []byte) {
	td.clearBitStream()
	td.data = data
	td.position = 0
	td.lastPosition = len(data)
}

// TryGetMeasurement attempts to get the next decoded measurement from the working buffer.
//gocyclo:ignore
func (td *Decoder) TryGetMeasurement(id *int32, timestamp *int64, stateFlags *uint32, value *float32) (bool, error) {
	if td.position == td.lastPosition && td.bitStreamIsEmpty() {
		td.clearBitStream()
		*id = 0
		*timestamp = 0
		*stateFlags = 0
		*value = 0.0
		return false, nil
	}

	// Given that the incoming pointID is not known in advance, the current
	// measurement will contain the encoding details for the next.

	// General compression strategy is to use delta-encoding for each
	// measurement component value that is received with the same identity.
	// See https://en.wikipedia.org/wiki/Delta_encoding

	// Delta-encoding sizes are embedded in the stream as type-specific
	// codes using as few bits as possible

	// Read next code for measurement ID decoding
	code, err := td.lastPoint.ReadCode()

	if err != nil {
		return false, err
	}

	if code == int32(codeWords.EndOfStream) {
		td.clearBitStream()
		*id = 0
		*timestamp = 0
		*stateFlags = 0
		*value = 0.0
		return false, nil
	}

	// Decode measurement ID and read next code for timestamp decoding
	if code <= int32(codeWords.PointIDXor32) {
		err := td.decodePointID(byte(code))

		if err != nil {
			return false, err
		}

		code, err = td.lastPoint.ReadCode()

		if err != nil {
			return false, err
		}

		if code < int32(codeWords.TimeDelta1Forward) {
			var message strings.Builder

			message.WriteString("expecting code >= ")
			message.WriteString(strconv.Itoa(int(codeWords.TimeDelta1Forward)))
			message.WriteString(" at position ")
			message.WriteString(strconv.Itoa(td.position))
			message.WriteString(" with last position ")
			message.WriteString(strconv.Itoa(td.lastPosition))

			return false, errors.New(message.String())
		}
	}

	// Assign decoded measurement ID to out parameter
	*id = td.lastPoint.PrevNextPointID1
	var nextPoint *pointMetadata
	var ok bool

	// Setup tracking for metadata associated with measurement ID and next point to decode
	if nextPoint, ok = td.points[*id]; !ok || nextPoint == nil {
		nextPoint = td.newPointMetadata()
		td.points[*id] = nextPoint
		nextPoint.PrevNextPointID1 = *id + 1
	}

	// Decode measurement timestamp and read next code for quality flags decoding
	if code <= int32(codeWords.TimeXor7Bit) {
		*timestamp = td.decodeTimestamp(byte(code))
		code, err = td.lastPoint.ReadCode()

		if err != nil {
			return false, err
		}

		if code < int32(codeWords.StateFlags2) {
			var message strings.Builder

			message.WriteString("expecting code >= ")
			message.WriteString(strconv.Itoa(int(codeWords.StateFlags2)))
			message.WriteString(" at position ")
			message.WriteString(strconv.Itoa(td.position))
			message.WriteString(" with last position ")
			message.WriteString(strconv.Itoa(td.lastPosition))

			return false, errors.New(message.String())
		}
	} else {
		*timestamp = td.prevTimestamp1
	}

	// Decode measurement state flags and read next code for measurement value decoding
	if code <= int32(codeWords.StateFlags7Bit32) {
		*stateFlags = td.decodeStateFlags(byte(code), nextPoint)
		code, err = td.lastPoint.ReadCode()

		if err != nil {
			return false, err
		}

		if code < int32(codeWords.Value1) {
			var message strings.Builder

			message.WriteString("expecting code >= ")
			message.WriteString(strconv.Itoa(int(codeWords.Value1)))
			message.WriteString(" at position ")
			message.WriteString(strconv.Itoa(td.position))
			message.WriteString(" with last position ")
			message.WriteString(strconv.Itoa(td.lastPosition))

			return false, errors.New(message.String())
		}
	} else {
		*stateFlags = nextPoint.PrevStateFlags1
	}

	// Since measurement value will almost always change, this is not put inside a function call
	var valueRaw uint32

	// Decode measurement value
	if code == int32(codeWords.Value1) {
		valueRaw = nextPoint.PrevValue1
	} else if code == int32(codeWords.Value2) {
		valueRaw = nextPoint.PrevValue2
		nextPoint.PrevValue2 = nextPoint.PrevValue1
		nextPoint.PrevValue1 = valueRaw
	} else if code == int32(codeWords.Value3) {
		valueRaw = nextPoint.PrevValue3
		nextPoint.PrevValue3 = nextPoint.PrevValue2
		nextPoint.PrevValue2 = nextPoint.PrevValue1
		nextPoint.PrevValue1 = valueRaw
	} else if code == int32(codeWords.ValueZero) {
		valueRaw = 0
		nextPoint.PrevValue3 = nextPoint.PrevValue2
		nextPoint.PrevValue2 = nextPoint.PrevValue1
		nextPoint.PrevValue1 = valueRaw
	} else {
		switch byte(code) {
		case codeWords.ValueXor4:
			valueRaw = uint32(td.readBits4()) ^ nextPoint.PrevValue1
		case codeWords.ValueXor8:
			valueRaw = uint32(td.data[td.position]) ^ nextPoint.PrevValue1
			td.position++
		case codeWords.ValueXor12:
			valueRaw = uint32(td.readBits4()) ^ uint32(td.data[td.position])<<4 ^ nextPoint.PrevValue1
			td.position++
		case codeWords.ValueXor16:
			valueRaw = uint32(td.data[td.position]) ^ uint32(td.data[td.position+1])<<8 ^ nextPoint.PrevValue1
			td.position += 2
		case codeWords.ValueXor20:
			valueRaw = uint32(td.readBits4()) ^ uint32(td.data[td.position])<<4 ^ uint32(td.data[td.position+1])<<12 ^ nextPoint.PrevValue1
			td.position += 2
		case codeWords.ValueXor24:
			valueRaw = uint32(td.data[td.position]) ^ uint32(td.data[td.position+1])<<8 ^ uint32(td.data[td.position+2])<<16 ^ nextPoint.PrevValue1
			td.position += 3
		case codeWords.ValueXor28:
			valueRaw = uint32(td.readBits4()) ^ uint32(td.data[td.position])<<4 ^ uint32(td.data[td.position+1])<<12 ^ uint32(td.data[td.position+2])<<20 ^ nextPoint.PrevValue1
			td.position += 3
		case codeWords.ValueXor32:
			valueRaw = uint32(td.data[td.position]) ^ uint32(td.data[td.position+1])<<8 ^ uint32(td.data[td.position+2])<<16 ^ uint32(td.data[td.position+3])<<24 ^ nextPoint.PrevValue1
			td.position += 4
		default:
			var message strings.Builder

			message.WriteString("invalid code received ")
			message.WriteString(strconv.Itoa(int(code)))
			message.WriteString(" at position ")
			message.WriteString(strconv.Itoa(td.position))
			message.WriteString(" with last position ")
			message.WriteString(strconv.Itoa(td.lastPosition))

			return false, errors.New(message.String())
		}

		nextPoint.PrevValue3 = nextPoint.PrevValue2
		nextPoint.PrevValue2 = nextPoint.PrevValue1
		nextPoint.PrevValue1 = valueRaw
	}

	// Assign decoded measurement value to out parameter
	*value = math.Float32frombits(valueRaw)
	td.lastPoint = nextPoint

	return true, nil
}

func (td *Decoder) decodePointID(code byte) error {
	switch code {
	case codeWords.PointIDXor4:
		td.lastPoint.PrevNextPointID1 = td.readBits4() ^ td.lastPoint.PrevNextPointID1
	case codeWords.PointIDXor8:
		td.lastPoint.PrevNextPointID1 = int32(td.data[td.position]) ^ td.lastPoint.PrevNextPointID1
		td.position += 1
	case codeWords.PointIDXor12:
		td.lastPoint.PrevNextPointID1 = td.readBits4() ^ int32(td.data[td.position])<<4 ^ td.lastPoint.PrevNextPointID1
		td.position += 1
	case codeWords.PointIDXor16:
		td.lastPoint.PrevNextPointID1 = int32(td.data[td.position]) ^ int32(td.data[td.position+1])<<8 ^ td.lastPoint.PrevNextPointID1
		td.position += 2
	case codeWords.PointIDXor20:
		td.lastPoint.PrevNextPointID1 = td.readBits4() ^ int32(td.data[td.position])<<4 ^ int32(td.data[td.position+1])<<12 ^ td.lastPoint.PrevNextPointID1
		td.position += 2
	case codeWords.PointIDXor24:
		td.lastPoint.PrevNextPointID1 = int32(td.data[td.position]) ^ int32(td.data[td.position+1])<<8 ^ int32(td.data[td.position+2])<<16 ^ td.lastPoint.PrevNextPointID1
		td.position += 3
	case codeWords.PointIDXor32:
		td.lastPoint.PrevNextPointID1 = int32(td.data[td.position]) ^ int32(td.data[td.position+1])<<8 ^ int32(td.data[td.position+2])<<16 ^ int32(td.data[td.position+3])<<24 ^ td.lastPoint.PrevNextPointID1
		td.position += 4
	default:
		var message strings.Builder

		message.WriteString("invalid code received ")
		message.WriteString(strconv.Itoa(int(code)))
		message.WriteString(" at position ")
		message.WriteString(strconv.Itoa(td.position))
		message.WriteString(" with last position ")
		message.WriteString(strconv.Itoa(td.lastPosition))

		return errors.New(message.String())
	}

	return nil
}

//gocyclo:ignore
func (td *Decoder) decodeTimestamp(code byte) int64 {
	var timestamp int64

	switch code {
	case codeWords.TimeDelta1Forward:
		timestamp = td.prevTimestamp1 + td.prevTimeDelta1
	case codeWords.TimeDelta2Forward:
		timestamp = td.prevTimestamp1 + td.prevTimeDelta2
	case codeWords.TimeDelta3Forward:
		timestamp = td.prevTimestamp1 + td.prevTimeDelta3
	case codeWords.TimeDelta4Forward:
		timestamp = td.prevTimestamp1 + td.prevTimeDelta4
	case codeWords.TimeDelta1Reverse:
		timestamp = td.prevTimestamp1 - td.prevTimeDelta1
	case codeWords.TimeDelta2Reverse:
		timestamp = td.prevTimestamp1 - td.prevTimeDelta2
	case codeWords.TimeDelta3Reverse:
		timestamp = td.prevTimestamp1 - td.prevTimeDelta3
	case codeWords.TimeDelta4Reverse:
		timestamp = td.prevTimestamp1 - td.prevTimeDelta4
	case codeWords.Timestamp2:
		timestamp = td.prevTimestamp2
	default:
		timestamp = td.prevTimestamp1 ^ int64(decode7BitUInt64(td.data, &td.position))
	}

	// Save the smallest delta time
	minDelta := abs(td.prevTimestamp1 - timestamp)

	if minDelta < td.prevTimeDelta4 && minDelta != td.prevTimeDelta1 && minDelta != td.prevTimeDelta2 && minDelta != td.prevTimeDelta3 {
		if minDelta < td.prevTimeDelta1 {
			td.prevTimeDelta4 = td.prevTimeDelta3
			td.prevTimeDelta3 = td.prevTimeDelta2
			td.prevTimeDelta2 = td.prevTimeDelta1
			td.prevTimeDelta1 = minDelta
		} else if minDelta < td.prevTimeDelta2 {
			td.prevTimeDelta4 = td.prevTimeDelta3
			td.prevTimeDelta3 = td.prevTimeDelta2
			td.prevTimeDelta2 = minDelta
		} else if minDelta < td.prevTimeDelta3 {
			td.prevTimeDelta4 = td.prevTimeDelta3
			td.prevTimeDelta3 = minDelta
		} else {
			td.prevTimeDelta4 = minDelta
		}
	}

	td.prevTimestamp2 = td.prevTimestamp1
	td.prevTimestamp1 = timestamp

	return timestamp
}

func (td *Decoder) decodeStateFlags(code byte, nextPoint *pointMetadata) uint32 {
	var stateFlags uint32

	if code == codeWords.StateFlags2 {
		stateFlags = nextPoint.PrevStateFlags2
	} else {
		stateFlags = decode7BitUInt32(td.data, &td.position)
	}

	nextPoint.PrevStateFlags2 = nextPoint.PrevStateFlags1
	nextPoint.PrevStateFlags1 = stateFlags

	return stateFlags
}

func (td *Decoder) readBit() int32 {
	if td.bitStreamCount == 0 {
		td.bitStreamCount = 8
		td.bitStreamCache = int32(td.data[td.position])
		td.position++
	}

	td.bitStreamCount--

	return td.bitStreamCache >> td.bitStreamCount & 1
}

func (td *Decoder) readBits4() int32 {
	return td.readBit()<<3 | td.readBit()<<2 | td.readBit()<<1 | td.readBit()
}

func (td *Decoder) readBits5() int32 {
	return td.readBit()<<4 | td.readBit()<<3 | td.readBit()<<2 | td.readBit()<<1 | td.readBit()
}

func decode7BitUInt32(stream []byte, position *int) uint32 {
	stream = stream[*position:]
	value := uint32(stream[0])

	if value < 128 {
		*position++
		return value
	}

	value ^= uint32(stream[1]) << 7

	if value < 16384 {
		*position += 2
		return value ^ 0x80
	}

	value ^= uint32(stream[2]) << 14

	if value < 2097152 {
		*position += 3
		return value ^ 0x4080
	}

	value ^= uint32(stream[3]) << 21

	if value < 268435456 {
		*position += 4
		return value ^ 0x204080
	}

	value ^= uint32(stream[4]) << 28
	*position += 5

	return value ^ 0x10204080
}

func decode7BitUInt64(stream []byte, position *int) uint64 {
	stream = stream[*position:]
	value := uint64(stream[0])

	if value < 128 {
		*position++
		return value
	}

	value ^= uint64(stream[1]) << 7

	if value < 16384 {
		*position += 2
		return value ^ 0x80
	}

	value ^= uint64(stream[2]) << 14

	if value < 2097152 {
		*position += 3
		return value ^ 0x4080
	}

	value ^= uint64(stream[3]) << 21

	if value < 268435456 {
		*position += 4
		return value ^ 0x204080
	}

	value ^= uint64(stream[4]) << 28

	if value < 34359738368 {
		*position += 5
		return value ^ 0x10204080
	}

	value ^= uint64(stream[5]) << 35

	if value < 4398046511104 {
		*position += 6
		return value ^ 0x810204080
	}

	value ^= uint64(stream[6]) << 42

	if value < 562949953421312 {
		*position += 7
		return value ^ 0x40810204080
	}

	value ^= uint64(stream[7]) << 49

	if value < 72057594037927936 {
		*position += 8
		return value ^ 0x2040810204080
	}

	value ^= uint64(stream[8]) << 56
	*position += 9

	return value ^ 0x102040810204080

}

func abs(value int64) int64 {
	if value < 0 {
		return value * -1
	}

	return value
}
