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
	"math"

	"github.com/sttp/goapi/sttp/guid"
)

// Maps 32-bit runtime IDs to 128-bit globally unique IDs.
// Additionally provides reverse lookup and an extra mapping
// to human-readable measurement keys.
type SignalIndexCache struct {
	reference     map[int32]uint32
	signalIDList  []guid.Guid
	sourceList    []string
	idList        []uint64
	signalIDCache map[guid.Guid]int32
	binaryLength  uint32
}

func (sic *SignalIndexCache) AddMeasurementKey(signalIndex int32, signalID guid.Guid, source string, id uint64, charSizeEstimate uint32 /* = 1 */) {
	sic.reference[signalIndex] = uint32(len(sic.signalIDList))
	sic.signalIDList = append(sic.signalIDList, signalID)
	sic.sourceList = append(sic.sourceList, source)
	sic.idList = append(sic.idList, id)
	sic.signalIDCache[signalID] = signalIndex

	// Char size here helps provide a rough-estimate on binary length used to reserve
	// bytes for a vector, if exact size is needed call RecalculateBinaryLength first
	sic.binaryLength += 32 + uint32(len(source))*charSizeEstimate
}

func (sic *SignalIndexCache) Clear() {
	sic.reference = map[int32]uint32{}
	sic.signalIDList = nil
	sic.sourceList = nil
	sic.idList = nil
	sic.signalIDCache = map[guid.Guid]int32{}
}

func (sic *SignalIndexCache) Contains(signalIndex int32) bool {
	_, ok := sic.reference[signalIndex]
	return ok
}

func (sic *SignalIndexCache) GetSignalID(signalIndex int32) guid.Guid {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.signalIDList[index]
	}

	return guid.Empty
}

func (sic *SignalIndexCache) GetSignalIDs() guid.HashSet {
	return guid.NewHashSet(sic.signalIDList)
}

func (sic *SignalIndexCache) GetSource(signalIndex int32) string {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.sourceList[index]
	}

	return ""
}

func (sic *SignalIndexCache) GetID(signalIndex int32) uint64 {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.idList[index]
	}

	return math.MaxUint64
}

func (sic *SignalIndexCache) GetMeasurementKey(signalIndex int32) (guid.Guid, string, uint64, bool) {
	if index, ok := sic.reference[signalIndex]; ok {
		return sic.signalIDList[index], sic.sourceList[index], sic.idList[index], true
	}

	return guid.Empty, "", 0, false
}

func (sic *SignalIndexCache) GetSignalIndex(signalID guid.Guid) int32 {
	if index, ok := sic.signalIDCache[signalID]; ok {
		return index
	}

	return -1
}

func (sic *SignalIndexCache) Count() uint32 {
	return uint32(len(sic.signalIDCache))
}

func (sic *SignalIndexCache) GetBinaryLength() uint32 {
	return sic.binaryLength
}

func (sic *SignalIndexCache) RecalculateBinaryLength(connection *SubscriberConnection) {
	var binaryLength uint32 = 28

	for i := 0; i < len(sic.signalIDList); i++ {
		binaryLength += 32 + uint32(len(connection.EncodeString(sic.sourceList[i])))
	}

	sic.binaryLength = binaryLength
}

func (sic *SignalIndexCache) Parse(buffer []byte, subscriberID *guid.Guid) {
	*subscriberID = guid.FromBytes(buffer[4:], false)
}

func (sic *SignalIndexCache) Serialize(connection *SubscriberConnection, buffer []byte) {
	// TODO: This will be needed by DataPublisher implementation
}
