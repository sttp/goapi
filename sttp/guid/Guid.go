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
	"bytes"
	"errors"
	"strconv"

	"github.com/google/uuid"
)

// Guid is a standard 128-bit UUID value (16-bytes) that can handle alternate wire serialization encodings.
type Guid [16]byte

// Empty is a Guid with a zero value.
var Empty Guid = Guid{}

// New creates a new random Guid value.
func New() Guid {
	return Guid(uuid.New())
}

// IsZero determines if the Guid value is its zero value, i.e., empty.
func (g Guid) IsZero() bool {
	return Equal(g, Empty)
}

// Equal returns true if this Guid and other Guid values are equal.
func (g Guid) Equal(other Guid) bool {
	return Equal(g, other)
}

// Equal returns true if the a and b Guid values are equal.
func Equal(a, b Guid) bool {
	return bytes.Equal(a[:], b[:])
}

// Compare returns an integer comparing this Guid (g) to other Guid. The result will be 0 if g==other, -1 if this g < other, and +1 if g > other.
func (g Guid) Compare(other Guid) int {
	return Compare(g, other)
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
	a = (uint32(g[0]) << 24) | (uint32(g[1]) << 16) | (uint32(g[2]) << 8) | uint32(g[3])
	b = uint16((uint32(g[4]) << 8) | uint32(g[5]))
	c = uint16((uint32(g[6]) << 8) | uint32(g[7]))
	d[0] = g[8]
	d[1] = g[9]
	d[2] = g[10]
	d[3] = g[11]
	d[4] = g[12]
	d[5] = g[13]
	d[6] = g[14]
	d[7] = g[15]

	return
}

// Parse decodes a Guid value from a string.
func Parse(s string) (Guid, error) {
	value, err := uuid.Parse(s)
	return Guid(value), err
}

// String returns the string form of a Guid, i.e., {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}.
func (g Guid) String() string {
	return "{" + uuid.UUID(g).String() + "}"
}

// STTP standard implementations, including C#, already use RFC encoding, the endiananness
// parameters exist for interop with implementations using non-RFC wire serializations:

// FromBytes creates a new Guid from a byte slice. Only first 16 bytes of slice are used.
// Returns an error if slice length is less than 16. Bytes are copied from the slice.
func FromBytes(data []byte, swapEndianness bool) (Guid, error) {
	if len(data) < 16 {
		return Empty, errors.New("Guid is 16 bytes in length, received " + strconv.Itoa(len(data)))
	}

	if swapEndianness {
		swapGuidEndianness(&data)
	}

	var g Guid
	copy(g[:], data[:16])
	return g, nil
}

// ToBytes creates a byte slice from a Guid.
func (g Guid) ToBytes(swapEndianness bool) []byte {
	data := g[:]

	if swapEndianness {
		swapGuidEndianness(&data)
	}

	return data
}

func swapGuidEndianness(data *[]byte) {
	swappedBytes := make([]byte, 16)
	var copy [8]byte

	for i := 0; i < 16; i++ {
		swappedBytes[i] = (*data)[i]

		if i < 8 {
			copy[i] = swappedBytes[i]
		}
	}

	// Swap endiananness, e.g., Microsoft and RFC encoding
	swappedBytes[3] = copy[0]
	swappedBytes[2] = copy[1]
	swappedBytes[1] = copy[2]
	swappedBytes[0] = copy[3]

	swappedBytes[4] = copy[5]
	swappedBytes[5] = copy[4]

	swappedBytes[6] = copy[7]
	swappedBytes[7] = copy[6]

	*data = swappedBytes
}
