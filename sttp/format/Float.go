//******************************************************************************************************
//  Float.go - Gbtc
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
//  09/23/2021 - J. Ritchie Carroll
//       Generated original version of source code, format functions inspired by:
//	     https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
//
//******************************************************************************************************

package format

import (
	"strconv"
	"strings"
)

// Float formats a floating-point number with a period as the decimal symbol and a comma as
// the numeric thousands grouping symbol.
func Float(f float64, prec int) string {
	return FloatWith(f, prec, '.', ',')
}

// FloatWith formats a floating-point number with the specified decimalSymbol, e.g., '.',
// and the specified numeric thousands groupSymbol, e.g., ','.
func FloatWith(f float64, prec int, decimalSymbol byte, groupSymbol byte) string {
	in := strconv.FormatFloat(f, 'f', prec, 64)
	decSymbolAsStr := string([]byte{decimalSymbol})

	if decimalSymbol != '.' {
		in = strings.Replace(in, ".", decSymbolAsStr, 1)
	}

	parts := strings.Split(in, decSymbolAsStr)
	var fraction string

	if len(parts) > 1 {
		in = parts[0]
		fraction = "." + parts[1]
	}

	digits := len(in)

	if f < 0 {
		digits-- // First character is the - sign (not a digit)
	}

	commas := (digits - 1) / 3
	out := make([]byte, len(in)+commas)

	if f < 0 {
		in, out[0] = in[1:], '-'
	}

	return formatNumber(in, out, groupSymbol) + fraction
}
