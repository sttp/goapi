//******************************************************************************************************
//  UInt.go - Gbtc
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

// UInt formats an unsigned-integer with a comma as the numeric thousands grouping symbol.
func UInt(i uint) string {
	return UIntWith(i, ',')
}

// UIntWith formats an unsigned-integer with specified numeric thousands groupSymbol, e.g., ','.
func UIntWith(i uint, groupSymbol byte) string {
	return UInt64With(uint64(i), groupSymbol)
}

// UInt64 formats a 64-bit unsigned-integer with a comma as the numeric thousands grouping symbol.
func UInt64(i uint64) string {
	return UInt64With(i, ',')
}

// UInt64With formats a 64-bit unsigned-integer with specified numeric thousands groupSymbol, e.g., ','.
func UInt64With(i uint64, groupSymbol byte) string {
	in := strconv.FormatUint(i, 10)
	digits := len(in)
	commas := (digits - 1) / 3
	out := make([]byte, len(in)+commas)

	return formatNumber(in, out, groupSymbol)
}
