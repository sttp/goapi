//******************************************************************************************************
//  SignalIndexCache.go - Gbtc
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

package transport

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/transport/tssc"
)

// SignalIndexCache maps 32-bit runtime IDs to 128-bit globally unique Measurement IDs. The structure
// additionally provides reverse lookup and an extra mapping  to human-readable measurement keys.
type SignalIndexCache struct {
	reference      map[int32]uint32
	signalIDList   []guid.Guid
	sourceList     []string
	idList         []uint64
	signalIDCache  map[guid.Guid]int32
	binaryLength   uint32
	maxSignalIndex uint32
	tsscDecoder    *tssc.Decoder
}

// NewSignalIndexCache makes a new SignalIndexCache
func NewSignalIndexCache() *SignalIndexCache {
	return &SignalIndexCache{
		reference:     make(map[int32]uint32),
		signalIDCache: make(map[guid.Guid]int32),
	}
}

// addRecord adds a new record to the SignalIndexCache for provided key Measurement details.
func (sic *SignalIndexCache) addRecord(ds *DataSubscriber, signalIndex int32, signalID guid.Guid, source string, id uint64, charSizeEstimate uint32 /* = 1 */) {
	index := uint32(len(sic.signalIDList))
	sic.reference[signalIndex] = index
	sic.signalIDList = append(sic.signalIDList, signalID)
	sic.sourceList = append(sic.sourceList, source)
	sic.idList = append(sic.idList, id)
	sic.signalIDCache[signalID] = signalIndex

	if index > sic.maxSignalIndex {
		sic.maxSignalIndex = index
	}

	metadata := ds.LookupMetadata(signalID)

	// Register measurement metadata if not defined already
	if len(metadata.Source) == 0 {
		metadata.Source = source
		metadata.ID = id
	}

	// Char size here helps provide a rough-estimate on binary length used to reserve
	// bytes for a vector, if exact size is needed call RecalculateBinaryLength first
	sic.binaryLength += 32 + uint32(len(source))*charSizeEstimate
}

// TODO: Function for use by DataPublisher
// clear removes all records from the SignalIndexCache.
// func (sic *SignalIndexCache) clear() {
// 	sic.reference = map[int32]uint32{}
// 	sic.signalIDList = nil
// 	sic.sourceList = nil
// 	sic.idList = nil
// 	sic.signalIDCache = map[guid.Guid]int32{}
// }

// Contains determines if the specified signalIndex exists with the SignalIndexCache.
func (sic *SignalIndexCache) Contains(signalIndex int32) bool {
	_, ok := sic.reference[signalIndex]
	return ok
}

// SignalID returns the signal ID Guid for the specified signalIndex in the SignalIndexCache.
func (sic *SignalIndexCache) SignalID(signalIndex int32) guid.Guid {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.signalIDList[index]
	}

	return guid.Empty
}

// SignalIDs returns a HashSet for all the Guid values found in the SignalIndexCache.
func (sic *SignalIndexCache) SignalIDs() guid.HashSet {
	return guid.NewHashSet(sic.signalIDList)
}

// Source returns the Measurement source string for the specified signalIndex in the SignalIndexCache.
func (sic *SignalIndexCache) Source(signalIndex int32) string {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.sourceList[index]
	}

	return ""
}

// ID returns the Measurement integer ID for the specified signalIndex in the SignalIndexCache.
func (sic *SignalIndexCache) ID(signalIndex int32) uint64 {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.idList[index]
	}

	return math.MaxUint64
}

// Record returns the key Measurement values, signalID Guid, source string, and integer ID and a
// final boolean value representing find success for the specified signalIndex in the SignalIndexCache.
func (sic *SignalIndexCache) Record(signalIndex int32) (guid.Guid, string, uint64, bool) {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.signalIDList[index], sic.sourceList[index], sic.idList[index], true
	}

	return guid.Empty, "", 0, false
}

// SignalIndex returns the signal index for the specified signalID Guid in the SignalIndexCache.
func (sic *SignalIndexCache) SignalIndex(signalID guid.Guid) int32 {
	if index, ok := sic.signalIDCache[signalID]; ok {
		return index
	}

	return -1
}

// MaxSignalIndex gets the largest signal index in the SignalIndexCache.
func (sic *SignalIndexCache) MaxSignalIndex() uint32 {
	return sic.maxSignalIndex
}

// Count returns the number of Measurement records that can be found in the SignalIndexCache.
func (sic *SignalIndexCache) Count() uint32 {
	return uint32(len(sic.signalIDCache))
}

// BinaryLength gets the binary length, in bytes, for the SignalIndexCache.
func (sic *SignalIndexCache) BinaryLength() uint32 {
	return sic.binaryLength
}

// TODO: Function for use by DataPublisher
// recalculateBinaryLength forces a new recalculation the cached binary length of the SignalIndexCache.
// func (sic *SignalIndexCache) recalculateBinaryLength(connection *SubscriberConnection) {
// 	var binaryLength uint32 = 28

// 	for i := 0; i < len(sic.signalIDList); i++ {
// 		binaryLength += 32 + uint32(len(connection.EncodeString(sic.sourceList[i])))
// 	}

// 	sic.binaryLength = binaryLength
// }

// decode parses a SignalIndexCache from the specified byte buffer received from a DataPublisher.
func (sic *SignalIndexCache) decode(ds *DataSubscriber, buffer []byte, subscriberID *guid.Guid) error {
	length := uint32(len(buffer))

	if length < 4 {
		return errors.New("not enough buffer provided to parse")
	}

	var offset uint32 = 0

	// Byte size of cache
	binaryLength := binary.BigEndian.Uint32(buffer)
	offset += 4

	if length < binaryLength {
		return errors.New("not enough buffer provided to parse")
	}

	var err error

	// Subscriber ID
	*subscriberID, err = guid.FromBytes(buffer[offset:], false)

	if err != nil {
		return errors.New("failed to parse SubscriberID: " + err.Error())
	}

	offset += 16

	// Number of references
	referenceCount := binary.BigEndian.Uint32(buffer[offset:])
	offset += 4

	var i uint32

	for i = 0; i < referenceCount; i++ {
		// Signal index
		signalIndex := int32(binary.BigEndian.Uint32(buffer[offset:]))
		offset += 4

		// Signal ID
		signalID, err := guid.FromBytes(buffer[offset:], false)

		if err != nil {
			return errors.New("failed to parse SignalID: " + err.Error())
		}

		offset += 16

		// Source
		sourceSize := binary.BigEndian.Uint32(buffer[offset:])
		offset += 4

		source := ds.DecodeString(buffer[offset : offset+sourceSize])
		offset += sourceSize

		// ID
		id := binary.BigEndian.Uint64(buffer[offset:])
		offset += 8

		sic.addRecord(ds, signalIndex, signalID, source, id, 1)
	}

	// There is additional data here about unauthorized signal IDs
	// that may need to be parsed in the future...

	return nil
}

//// Encode serializes a SignalIndexCache to a byte buffer for publication to a DataSubscriber.
//func (sic *SignalIndexCache) Encode(connection *SubscriberConnection, buffer []byte) {
//	// TODO: This will be needed by DataPublisher implementation
//}
