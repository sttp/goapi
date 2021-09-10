package transport

import (
	//lint:ignore ST1001 statically include native STTP types as root
	. "github.com/sttp/goapi/sttp"
)

type SignalIndexCache struct {
	reference     map[int32]uint32
	signalIDList  []Guid
	sourceList    []string
	idList        []uint64
	signalIDCache map[Guid]int32
	binaryLength  uint32
}

func (sic *SignalIndexCache) AddMeasurementKey(signalIndex int32, signalID Guid, source string, id uint64, charSizeEstimate uint32 /* = 1 */) {
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
	sic.signalIDCache = map[Guid]int32{}
}

func (sic *SignalIndexCache) Contains(signalIndex int32) bool {
	_, ok := sic.reference[signalIndex]
	return ok
}

func (sic *SignalIndexCache) GetSignalID(signalIndex int32) Guid {
	index, ok := sic.reference[signalIndex]

	if ok {
		return sic.signalIDList[index]
	}

	return EmptyGuid
}
