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

import (
	"github.com/google/uuid"
)

// Guid is a standard UUID value that can handle alternate wire serialization options.
type Guid uuid.UUID

// Empty is a Guid with a zero value.
var Empty Guid = Guid(uuid.Nil)

// New creates a new random Guid value.
func New() Guid {
	return Guid(uuid.New())
}

// IsZero determines if the Guid value is its zero value, i.e., empty.
func (g Guid) IsZero() bool {
	return g == Empty
}

// Equal returns true if the a and b Guid values are equal.
func Equal(a, b Guid) bool {
	g1 := [16]byte(a)
	g2 := [16]byte(b)

	for i := 0; i < 16; i++ {
		if g1[i] != g2[2] {
			return false
		}
	}

	return true
}

// Compare returns an integer comparing two Guid values. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare(a, b Guid) int {
	a1, b1, c1, d1 := a.Components()
	a2, b2, c2, d2 := b.Components()

	if a1 != a2 {
		return result(a1, a2)
	}

	if b1 != b2 {
		return result(uint32(b1), uint32(b2))
	}

	if c1 != c2 {
		return result(uint32(c1), uint32(c2))
	}

	for i := 0; i < 8; i++ {
		if d1[i] != d2[i] {
			return result(uint32(d1[i]), uint32(d2[i]))
		}
	}

	return 0
}

func result(left, right uint32) int {
	if left < right {
		return -1
	}

	return 1
}

// Components gets the Guid value as its constituent components.
func (g Guid) Components() (a uint32, b, c uint16, d [8]byte) {

	bytes := [16]byte(g)

	a = (uint32(bytes[3]) << 24) | (uint32(bytes[2]) << 16) | (uint32(bytes[1]) << 8) | uint32(bytes[0])
	b = (uint16(bytes[5]) << 8) | uint16(bytes[4])
	c = (uint16(bytes[7]) << 8) | uint16(bytes[6])
	d[0] = bytes[8]
	d[1] = bytes[9]
	d[2] = bytes[10]
	d[3] = bytes[11]
	d[4] = bytes[12]
	d[5] = bytes[13]
	d[6] = bytes[14]
	d[7] = bytes[15]

	return
}

// Parse decodes a Guid value from a string.
func Parse(s string) (Guid, error) {
	value, err := uuid.Parse(s)
	return Guid(value), err
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
	var encodedBytes []byte

	if swapEndianness {
		swappedBytes := make([]byte, 16)
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
		encodedBytes = data[:16]
	}

	guid, err := uuid.FromBytes(encodedBytes)

	return Guid(guid), err
}
