//******************************************************************************************************
//  SignalKind.go - Gbtc
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
	"strings"
	"unicode"
)

// SignalKindEnum defines the type for the SignalKind enumeration.
type SignalKindEnum uint16

// SignalKind is an enumeration of the possible kinds of signals a Measurement can represent.
var SignalKind = struct {
	// Angle defines a phase angle signal kind (could be a voltage or a current).
	Angle SignalKindEnum
	// Magnitude defines a phase magnitude signal kind (could be a voltage or a current).
	Magnitude SignalKindEnum
	// Frequency defines a line frequency signal kind.
	Frequency SignalKindEnum
	// DfDt defines a frequency delta over time (dF/dt) signal kind.
	DfDt SignalKindEnum
	// Status defines a status flags signal kind.
	Status SignalKindEnum
	// Digital defines a digital value signal kind.
	Digital SignalKindEnum
	// Analog defines an analog value signal kind.
	Analog SignalKindEnum
	// Calculation defines a calculated value signal kind.
	Calculation SignalKindEnum
	// Statistic defines a statistical value signal kind.
	Statistic SignalKindEnum
	// Alarm defines an alarm value signal kind.
	Alarm SignalKindEnum
	// Quality defines a quality flags signal kind.
	Quality SignalKindEnum
	// Unknown defines an undetermined signal kind.
	Unknown SignalKindEnum
}{
	Angle:       0,
	Magnitude:   1,
	Frequency:   2,
	DfDt:        3,
	Status:      4,
	Digital:     5,
	Analog:      6,
	Calculation: 7,
	Statistic:   8,
	Alarm:       9,
	Quality:     10,
	Unknown:     11,
}

// String gets the SignalKind enumeration value as a string.
func (ske SignalKindEnum) String() string {
	switch ske {
	case SignalKind.Angle:
		return "Angle"
	case SignalKind.Magnitude:
		return "Magnitude"
	case SignalKind.Frequency:
		return "Frequency"
	case SignalKind.DfDt:
		return "DfDt"
	case SignalKind.Status:
		return "Status"
	case SignalKind.Digital:
		return "Digital"
	case SignalKind.Analog:
		return "Analog"
	case SignalKind.Calculation:
		return "Calculation"
	case SignalKind.Statistic:
		return "Statistic"
	case SignalKind.Alarm:
		return "Alarm"
	case SignalKind.Quality:
		return "Quality"
	default:
		return "Unknown"
	}
}

// Acronym gets the SignalKind enumeration value as its two-character acronym string.
func (ske SignalKindEnum) Acronym() string {
	switch ske {
	case SignalKind.Angle:
		return "PA"
	case SignalKind.Magnitude:
		return "PM"
	case SignalKind.Frequency:
		return "FQ"
	case SignalKind.DfDt:
		return "DF"
	case SignalKind.Status:
		return "SF"
	case SignalKind.Digital:
		return "DV"
	case SignalKind.Analog:
		return "AV"
	case SignalKind.Calculation:
		return "CV"
	case SignalKind.Statistic:
		return "ST"
	case SignalKind.Alarm:
		return "AL"
	case SignalKind.Quality:
		return "QF"
	default:
		return "??"
	}
}

// SignalTypeAcronym gets the specific four-character signal type acronym for a SignalKind
// enumeration value and phasor type, i.e., "V" voltage or "I" current.
func (ske SignalKindEnum) SignalTypeAcronym(phasorType rune) string {
	// A SignalType represents a more specific measurement type than SignalKind, i.e.,
	// a phasor type (voltage or current) can also be determined by the type.
	switch ske {
	case SignalKind.Angle:
		if unicode.ToUpper(phasorType) == 'V' {
			return "VPHA"
		}
		return "IPHA"
	case SignalKind.Magnitude:
		if unicode.ToUpper(phasorType) == 'V' {
			return "VPHM"
		}
		return "IPHM"
	case SignalKind.Frequency:
		return "FREQ"
	case SignalKind.DfDt:
		return "DFDT"
	case SignalKind.Status:
		return "FLAG"
	case SignalKind.Digital:
		return "DIGI"
	case SignalKind.Analog:
		return "ALOG"
	case SignalKind.Calculation:
		return "CALC"
	case SignalKind.Statistic:
		return "STAT"
	case SignalKind.Alarm:
		return "ALRM"
	case SignalKind.Quality:
		return "QUAL"
	}

	return "NULL"
}

// ParseSignalKindAcronym gets the SignalKind enumeration value for the specified two-character acronym.
func ParseSignalKindAcronym(acronym string) SignalKindEnum {
	acronym = strings.TrimSpace(strings.ToUpper(acronym))

	if acronym == "PA" { // Phase Angle
		return SignalKind.Angle
	}

	if acronym == "PM" { // Phase Magnitude
		return SignalKind.Magnitude
	}

	if acronym == "FQ" { // Frequency
		return SignalKind.Frequency
	}

	if acronym == "DF" { // dF/dt
		return SignalKind.DfDt
	}

	if acronym == "SF" { // Status Flags
		return SignalKind.Status
	}

	if acronym == "DV" { // Digital Value
		return SignalKind.Digital
	}

	if acronym == "AV" { // Analog Value
		return SignalKind.Analog
	}

	if acronym == "CV" { // Calculated Value
		return SignalKind.Calculation
	}

	if acronym == "ST" { // Statistical Value
		return SignalKind.Statistic
	}

	if acronym == "AL" { // Alarm Value
		return SignalKind.Alarm
	}

	if acronym == "QF" { // Quality Flags
		return SignalKind.Quality
	}

	return SignalKind.Unknown
}
