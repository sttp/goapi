package sttp

import (
	"time"

	"github.com/sttp/goapi/sttp/StateFlags"
)

type MeasurementMetadata struct {
	// Measurement's globally unique identifier.
	SignalID Guid

	// Additive value modifier.
	Adder float64

	// Multiplicative value modifier.
	Multipler float64

	// Identification number used in human-readable measurement key.
	ID uint64

	// Source used in human-readable measurement key.
	Source string

	// Human-readable tag name to help describe the measurement.
	Tag string
}

type Measurement struct {
	// Measurement's globally unique identifier.
	SignalID Guid

	// Instantaneous value of the measurement.
	Value float64

	// The time, in ticks, that this measurement was taken.
	Timestamp Ticks

	// Flags indicating the state of the measurement as reported by the device that took it.
	Flags StateFlags.StateFlags
}

var (
	measurementRegistry = make(map[Guid]MeasurementMetadata)
)

func RegisterMeasurementMetadata(metadata MeasurementMetadata) {
	measurementRegistry[metadata.SignalID] = metadata
}

func LookupMeasurementMetadata(signalID Guid) (MeasurementMetadata, bool) {
	metadata, ok := measurementRegistry[signalID]
	return metadata, ok
}

func (m *Measurement) AdjustedValue() float64 {
	metadata, ok := measurementRegistry[m.SignalID]

	if ok {
		return m.Value*metadata.Multipler + metadata.Adder
	}

	return m.Value
}

func (m *Measurement) GetDateTime() time.Time {
	return ToTime(m.Timestamp)
}
