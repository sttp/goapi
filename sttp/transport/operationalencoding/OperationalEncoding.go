//******************************************************************************************************
//  OperationalEncoding.go - Gbtc
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

package operationalencoding

// OperationalEncoding define the possible string encoding options of an STTP session.
type OperationalEncoding uint32

const (
	// UTF16LE targets little-endian 16-bit Unicode character encoding for strings (deprecated).
	UTF16LE OperationalEncoding = 0x00000000
	// UTF16BE targets big-endian 16-bit Unicode character encoding for strings (deprecated).
	UTF16BE OperationalEncoding = 0x00000100
	// UTF8 targets 8-bit variable-width Unicode character encoding for strings.
	UTF8 OperationalEncoding = 0x00000200
)
