//******************************************************************************************************
//  SubscriberConnection.go - Gbtc
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

// SubscriberConnection represents a subscriber connection to a data publisher
type SubscriberConnection struct {
	encoding OperationalEncodingEnum
}

func (sc *SubscriberConnection) DecodeString(data []byte, length uint32) string {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if sc.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return string(data[:length])
}

func (sc *SubscriberConnection) EncodeString(value string) []byte {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if sc.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return []byte(value)
}
