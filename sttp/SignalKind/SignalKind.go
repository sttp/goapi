package SignalKind

import (
	"strings"
	"unicode"
)

type SignalKind uint16

const (
	Angle       SignalKind = iota // Phase angle
	Magnitude                     // Phase magnitude
	Frequency                     // Line frequency
	DfDt                          // Frequency delta over time (dF/dt)
	Status                        // Status flags
	Digital                       // Digital value
	Analog                        // Analog value
	Calculation                   // Calculated value
	Statistic                     // Statistical value
	Alarm                         // Alarm value
	Quality                       // Quality flags
	Unknown                       // Undetermined signal type
)

var (
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

func GetAcronym(kind SignalKind, phasorType rune) string {
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

// Gets the "SignalKind" enum for the specified "acronym".
//  params:
//       acronym: Acronym of the desired "SignalKind"
//  returns: The "SignalKind" for the specified "acronym".
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
