//******************************************************************************************************
//  MeasurementMetadata.go - Gbtc
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
	"strconv"
	"strings"
	"time"

	"github.com/sttp/goapi/sttp/guid"
)

// MeasurementMetadata defines the ancillary information associated with a Measurement.
// Metadata gets cached in a registry associated with a DataSubscriber.
type MeasurementMetadata struct {
	// Measurement's globally unique identifier.
	SignalID guid.Guid

	// Additive value modifier. Allows for linear value adjustment.
	Adder float64

	// Multiplicative value modifier. Allows for linear value adjustment.
	Multiplier float64

	// Identification number used in human-readable measurement key.
	ID uint64

	// Source used in human-readable measurement key.
	Source string

	// SignalType defines a signal type acronym for the measurement, e.g., FREQ.
	SignalType string

	// SignalReference defines reference info about a signal based on measurement original source.
	SignalReference string

	// Description defines a general description for the measurement.
	Description string

	// UpdatedOn defines the timestamp of when the metadata was last updated.
	UpdatedOn time.Time

	// Human-readable tag name or reference value used to help describe or help identify the measurement.
	Tag string
}

// ParseSignalReference attempts to parse a normally formatted signal reference into a
// signal kind and position representing original source protocol details.
func (mm *MeasurementMetadata) ParseSignalReference() (source string, signalKind SignalKindEnum, position int) {
	signalReference := mm.SignalReference
	parts := strings.Split(signalReference, "-")

	if len(parts) > 1 {
		lastIndex := len(parts) - 1
		typeInfo := parts[lastIndex]
		signalKind = ParseSignalKindAcronym(typeInfo[:2])
		position, _ = strconv.Atoi(typeInfo[2:])
		source = strings.Join(parts[:lastIndex], "-")
	}

	return
}
