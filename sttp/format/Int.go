//******************************************************************************************************
//  Int.go - Gbtc
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
)

// Int formats an integer with a comma as the numeric thousands grouping symbol.
func Int(i int) string {
	return IntWith(i, ',')
}

// IntWith formats an integer with specified numeric thousands groupSymbol, e.g., ','.
func IntWith(i int, groupSymbol byte) string {
	return Int64With(int64(i), groupSymbol)
}

// Int64 formats a 64-bit integer with a comma as the numeric thousands grouping symbol.
func Int64(i int64) string {
	return Int64With(i, ',')
}

// Int64With formats a 64-bit integer with specified numeric thousands groupSymbol, e.g., ','.
func Int64With(i int64, groupSymbol byte) string {
	in := strconv.FormatInt(i, 10)
	digits := len(in)

	if i < 0 {
		digits-- // First character is the - sign (not a digit)
	}

	commas := (digits - 1) / 3
	out := make([]byte, len(in)+commas)

	if i < 0 {
		in, out[0] = in[1:], '-'
	}

	return formatNumber(in, out, groupSymbol)
}
