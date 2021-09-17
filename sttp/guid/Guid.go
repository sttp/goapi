//******************************************************************************************************
//  Guid.go - Gbtc
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

package guid

import "github.com/google/uuid"

// Guid is a standard UUID value that can handle alternate wire serialization options.
type Guid uuid.UUID

// Empty is a Guid with a zero value.
var Empty Guid = Guid(uuid.Nil)

// New creates a new random Guid value.
func New() Guid {
	return Guid(uuid.New())
}

// Parse decodes a Guid value from a string.
func Parse(value string) Guid {
	guid, err := uuid.Parse(value)

	if err == nil {
		return Guid(guid)
	}

	panic("Failed to parse Guid from string \"" + value + "\": " + err.Error())
}

// String returns the string form of a Guid, i.e., {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx},
// or "" if Guid is invalid.
func (g Guid) String() string {
	image := uuid.UUID(g).String()

	if len(image) > 0 {
		return "{" + image + "}"
	}

	return ""
}

// FromBytes creates a new Guid from a byte slice.
func FromBytes(data []byte, swapEndianness bool) (Guid, error) {
	swappedBytes := make([]byte, 16)
	var encodedBytes []byte

	if swapEndianness {
		var copy [8]byte

		for i := 0; i < 16; i++ {
			swappedBytes[i] = data[i]

			if i < 8 {
				copy[i] = swappedBytes[i]
			}
		}

		// Convert Microsoft encoding to RFC
		swappedBytes[3] = copy[0]
		swappedBytes[2] = copy[1]
		swappedBytes[1] = copy[2]
		swappedBytes[0] = copy[3]

		swappedBytes[4] = copy[5]
		swappedBytes[5] = copy[4]

		swappedBytes[6] = copy[7]
		swappedBytes[7] = copy[6]

		encodedBytes = swappedBytes
	} else {
		encodedBytes = data
	}

	guid, err := uuid.FromBytes(encodedBytes)

	return Guid(guid), err
}
