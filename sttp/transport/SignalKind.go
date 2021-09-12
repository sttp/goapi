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
	// Angle defines a phase angle signal type.
	Angle SignalKindEnum
	// Magnitude defines a phase magnitude signal type.
	Magnitude SignalKindEnum
	// Frequency defines a line frequency signal type.
	Frequency SignalKindEnum
	// DfDt defines a frequency delta over time (dF/dt) signal type.
	DfDt SignalKindEnum
	// Status defines a status flags signal type.
	Status SignalKindEnum
	// Digital defines a digital value signal type.
	Digital SignalKindEnum
	// Analog defines an analog value signal type.
	Analog SignalKindEnum
	// Calculation defines a calculated value signal type.
	Calculation SignalKindEnum
	// Statistic defines a statistical value signal type.
	Statistic SignalKindEnum
	// Alarm defines an alarm value signal type.
	Alarm SignalKindEnum
	// Quality defines a quality flags signal type.
	Quality SignalKindEnum
	// Unknown defines an undetermined signal type.
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

var (
	// Description defines the string representations of the SignalKind enumeration values.
	Description = [...]string{
		"Angle",
		"Magnitude",
		"Frequency",
		"DfDt",
		"Status",
		"Digital",
		"Analog",
		"Calculation",
		"Statistic",
		"Alarm",
		"Quality",
		"Unknown"}

	// Acronym defines the abbreviated string representations of the SignalKind enumeration values.
	Acronym = [...]string{
		"PA",
		"PM",
		"FQ",
		"DF",
		"SF",
		"DV",
		"AV",
		"CV",
		"ST",
		"AL",
		"QF",
		"??"}
)

// GetSignalTypeAcronym gets the specific four-character signal type acronym for a SignalKind
// enumeration value and phasor type, i.e., "V" voltage or "I" current.
func GetSignalTypeAcronym(kind SignalKindEnum, phasorType rune) string {
	switch kind {
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

// Parse gets the SignalKind enumeration value for the specified acronym.
func Parse(acronym string) SignalKindEnum {
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
