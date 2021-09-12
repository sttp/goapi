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

package signalkind

import (
	"strings"
	"unicode"
)

// SignalKind defines the kinds of signals a Measurement can represent
type SignalKind uint16

const (
	// Angle defines a phase angle signal type.
	Angle SignalKind = iota
	// Magnitude defines a phase magnitude signal type.
	Magnitude
	// Frequency defines a line frequency signal type.
	Frequency
	// DfDt defines a frequency delta over time (dF/dt) signal type.
	DfDt
	// Status defines a status flags signal type.
	Status
	// Digital defines a digital value signal type.
	Digital
	// Analog defines an analog value signal type.
	Analog
	// Calculation defines a calculated value signal type.
	Calculation
	// Statistic defines a statistical value signal type.
	Statistic
	// Alarm defines an alarm value signal type.
	Alarm
	// Quality defines a quality flags signal type.
	Quality
	// Unknown defines an undetermined signal type.
	Unknown
)

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
func GetSignalTypeAcronym(kind SignalKind, phasorType rune) string {
	switch kind {
	case Angle:
		if unicode.ToUpper(phasorType) == 'V' {
			return "VPHA"
		}
		return "IPHA"
	case Magnitude:
		if unicode.ToUpper(phasorType) == 'V' {
			return "VPHM"
		}
		return "IPHM"
	case Frequency:
		return "FREQ"
	case DfDt:
		return "DFDT"
	case Status:
		return "FLAG"
	case Digital:
		return "DIGI"
	case Analog:
		return "ALOG"
	case Calculation:
		return "CALC"
	case Statistic:
		return "STAT"
	case Alarm:
		return "ALRM"
	case Quality:
		return "QUAL"
	}

	return "NULL"
}

// Parse gets the SignalKind enumeration value for the specified acronym.
func Parse(acronym string) SignalKind {
	acronym = strings.TrimSpace(strings.ToUpper(acronym))

	if acronym == "PA" { // Phase Angle
		return Angle
	}

	if acronym == "PM" { // Phase Magnitude
		return Magnitude
	}

	if acronym == "FQ" { // Frequency
		return Frequency
	}

	if acronym == "DF" { // dF/dt
		return DfDt
	}

	if acronym == "SF" { // Status Flags
		return Status
	}

	if acronym == "DV" { // Digital Value
		return Digital
	}

	if acronym == "AV" { // Analog Value
		return Analog
	}

	if acronym == "CV" { // Calculated Value
		return Calculation
	}

	if acronym == "ST" { // Statistical Value
		return Statistic
	}

	if acronym == "AL" { // Alarm Value
		return Alarm
	}

	if acronym == "QF" { // Quality Flags
		return Quality
	}

	return Unknown
}
