//******************************************************************************************************
//  Measurement.go - Gbtc
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

package measurement

import (
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/stateflags"
	"github.com/sttp/goapi/sttp/ticks"
)

type MeasurementMetadata struct {
	// Measurement's globally unique identifier.
	SignalID guid.Guid

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
	SignalID guid.Guid

	// Instantaneous value of the measurement.
	Value float64

	// The time, in ticks, that this measurement was taken.
	Timestamp ticks.Ticks

	// Flags indicating the state of the measurement as reported by the device that took it.
	Flags stateflags.StateFlags
}

var (
	measurementRegistry = make(map[guid.Guid]MeasurementMetadata)
)

func RegisterMetadata(metadata MeasurementMetadata) {
	measurementRegistry[metadata.SignalID] = metadata
}

func LookupMetadata(signalID guid.Guid) (MeasurementMetadata, bool) {
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
	return ticks.ToTime(m.Timestamp)
}