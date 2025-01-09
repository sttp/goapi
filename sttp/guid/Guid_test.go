//******************************************************************************************************
//  Guid_test.go - Gbtc
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
//  10/07/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package guid

import (
	"bytes"
	"testing"
	"unsafe"
)

const (
	gs1 string = "{b4a26a66-a073-44a0-b03b-55d97badef74}"
	gs2 string = "{b4a26a66-a073-44a0-b03b-55d97badef75}"
	gs3 string = "{3db9da3a-6719-45ab-9bf6-87545f4025a8}"
	gs4 string = "{00000001-0002-0003-0405-060708090a0b}" // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
	gs5 string = "{3db9da3a-6719-9bf6-45ab-87545f4025a8}"
	gs6 string = "{3db9da3a-9bf6-45ab-6719-87545f4025a8}"
	gsz string = "{00000000-0000-0000-0000-000000000000}"
)

// gocyclo: ignore
func TestGuidParsing(t *testing.T) {
	var g1, g2, g3, g4 Guid
	var err error

	if g1, err = Parse(gs1); err != nil {
		t.Fatalf("TestGuidParsing: failed to parse guid " + gs1)
	}

	if g2, err = Parse(gs2); err != nil {
		t.Fatalf("TestGuidParsing: failed to parse guid " + gs2)
	}

	if g3, err = Parse(gs3); err != nil {
		t.Fatalf("TestGuidParsing: failed to parse guid " + gs3)
	}

	if g4, err = Parse(gs4); err != nil {
		t.Fatalf("TestGuidParsing: failed to parse guid " + gs4)
	}

	a1, b1, c1, d1 := g1.Components()
	a2, b2, c2, d2 := g2.Components()
	a3, b3, c3, d3 := g3.Components()
	a4, b4, c4, d4 := g4.Components()

	if a1 != 3030542950 && b1 != 41075 && c1 != 17568 &&
		d1[0] != 176 && d1[1] != 59 && d1[2] != 85 && d1[3] != 217 &&
		d1[4] != 123 && d1[5] != 173 && d1[6] != 239 && d1[7] != 116 {
		t.Fatalf("TestGuidParsing: failed to get proper components from guid " + gs1)
	}

	if a2 != 3030542950 && b2 != 41075 && c2 != 17568 &&
		d2[0] != 176 && d2[1] != 59 && d2[2] != 85 && d2[3] != 217 &&
		d2[4] != 123 && d2[5] != 173 && d2[6] != 239 && d2[7] != 117 {
		t.Fatalf("TestGuidParsing: failed to get proper components from guid " + gs2)
	}

	if a3 != 1035590202 && b3 != 26393 && c3 != 17835 &&
		d3[0] != 155 && d3[1] != 246 && d3[2] != 135 && d3[3] != 84 &&
		d3[4] != 95 && d3[5] != 64 && d3[6] != 37 && d3[7] != 168 {
		t.Fatalf("TestGuidParsing: failed to get proper components from guid " + gs3)
	}

	if a4 != 1 && b4 != 2 && c4 != 3 &&
		d4[0] != 4 && d4[1] != 5 && d4[2] != 6 && d4[3] != 7 &&
		d4[4] != 8 && d4[5] != 9 && d4[6] != 10 && d4[7] != 11 {
		t.Fatalf("TestGuidParsing: failed to get proper components from guid " + gs4)
	}

	if g1.String() != gs1 {
		t.Fatalf("TestGuidParsing: string generation does not match for " + gs1)
	}

	if g2.String() != gs2 {
		t.Fatalf("TestGuidParsing: string generation does not match for " + gs2)
	}

	if g3.String() != gs3 {
		t.Fatalf("TestGuidParsing: string generation does not match for " + gs3)
	}

	if g4.String() != gs4 {
		t.Fatalf("TestGuidParsing: string generation does not match for " + gs4)
	}

	if Empty.String() != gsz {
		t.Fatalf("TestGuidParsing: string generation does not match for " + gsz)
	}

}
func TestNewGuidRandomness(t *testing.T) {
	for i := 0; i < 10000; i++ {
		if New().Equal(New()) || Equal(New(), New()) {
			t.Fatalf("TestNewGuidRandomness: encountered non-unique Guid after %d generations", i*4)
		}
	}
}

// gocyclo: ignore
func TestZeroGuid(t *testing.T) {
	var gz, zero Guid
	var err error

	if gz, err = Parse(gsz); err != nil {
		t.Fatalf("TestZeroGuid: failed to parse guid " + gsz)
	}

	if !gz.Equal(zero) {
		t.Fatalf("TestZeroGuid: parsed zero-value guid not equal to zero guid")
	}

	if !gz.IsZero() {
		t.Fatalf("TestZeroGuid: parsed zero-value guid not equal to Empty (per IsZero receiver)")
	}

	if !gz.Equal(Empty) {
		t.Fatalf("TestZeroGuid: parsed zero-value guid not equal to Empty")
	}

	if !Empty.Equal(zero) {
		t.Fatalf("TestZeroGuid: Empty guid not equal to zero guid")
	}

	a1, b1, c1, d1 := gz.Components()
	a2, b2, c2, d2 := zero.Components()
	a3, b3, c3, d3 := Empty.Components()

	if a1 != 0 && b1 != 0 && c1 != 0 &&
		d1[0] != 0 && d1[1] != 0 && d1[2] != 0 && d1[3] != 0 &&
		d1[4] != 0 && d1[5] != 0 && d1[6] != 0 && d1[7] != 0 {
		t.Fatalf("TestZeroGuid: components of zero-value guid not all zero")
	}

	if a2 != 0 && b2 != 0 && c2 != 0 &&
		d2[0] != 0 && d2[1] != 0 && d2[2] != 0 && d2[3] != 0 &&
		d2[4] != 0 && d2[5] != 0 && d2[6] != 0 && d2[7] != 0 {
		t.Fatalf("TestZeroGuid: components of zero guid not all zero")
	}

	if a3 != 0 && b3 != 0 && c3 != 0 &&
		d3[0] != 0 && d3[1] != 0 && d3[2] != 0 && d3[3] != 0 &&
		d3[4] != 0 && d3[5] != 0 && d3[6] != 0 && d3[7] != 0 {
		t.Fatalf("TestZeroGuid: components of Empty guid not all zero")
	}
}

// gocyclo: ignore
func TestGuidCompare(t *testing.T) {
	var g1, g2, g3, g4, g5, g6 Guid
	var err error

	if g1, err = Parse(gs1); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs1)
	}

	if g2, err = Parse(gs2); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs2)
	}

	if g3, err = Parse(gs3); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs3)
	}

	if g4, err = Parse(gs4); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs4)
	}

	if g5, err = Parse(gs5); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs5)
	}

	if g6, err = Parse(gs6); err != nil {
		t.Fatalf("TestGuidCompare: failed to parse guid " + gs6)
	}

	if Compare(g1, g1) != g1.Compare(g1) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g1, g2) < 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g1, g3) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g1, g4) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g2, g1) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g3, g5) < 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g3, g6) < 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g1, Empty) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g2, Empty) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g3, Empty) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(g4, Empty) > 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}

	if !(Compare(Empty, g4) < 0) {
		t.Fatalf("TestGuidCompare: results of guid Compare invalid")
	}
}

func TestGuidToFromBytes(t *testing.T) {
	var g1, g2, g3, g4, gz Guid
	var err error

	if g1, err = Parse(gs1); err != nil {
		t.Fatalf("TestGuidToFromBytes: failed to parse guid " + gs1)
	}

	if g2, err = Parse(gs2); err != nil {
		t.Fatalf("TestGuidToFromBytes: failed to parse guid " + gs2)
	}

	if g3, err = Parse(gs3); err != nil {
		t.Fatalf("TestGuidToFromBytes: failed to parse guid " + gs3)
	}

	if g4, err = Parse(gs4); err != nil {
		t.Fatalf("TestGuidToFromBytes: failed to parse guid " + gs4)
	}

	if gz, err = Parse(gsz); err != nil {
		t.Fatalf("TestGuidToFromBytes: failed to parse guid " + gsz)
	}

	testGuidToFromBytes(g1, gs1, false, t)
	testGuidToFromBytes(g2, gs2, false, t)
	testGuidToFromBytes(g3, gs3, false, t)
	testGuidToFromBytes(g4, gs4, false, t)
	testGuidToFromBytes(gz, gsz, false, t)

	testGuidToFromBytes(g1, gs1, true, t)
	testGuidToFromBytes(g2, gs2, true, t)
	testGuidToFromBytes(g3, gs3, true, t)
	testGuidToFromBytes(g4, gs4, true, t)
	testGuidToFromBytes(gz, gsz, true, t)

	// Test negative case
	if _, err := FromBytes([]byte{0, 0}, false); err == nil {
		t.Fatalf("TestGuidToFromBytes: unexpected success, short slice expected to fail guid parse")
	}
}

func testGuidToFromBytes(g Guid, gs string, swapEndianness bool, t *testing.T) {
	gbf := g.ToBytes(swapEndianness)
	gfs := [16]byte(g)
	gbs := gfs[:16]

	if swapEndianness {
		swapGuidEndianness(&gbs)
	}

	if !bytes.Equal(gbf, gbs) {
		t.Fatal("TestGuidToFromBytes: ToBytes test compare failed for guid " + gs)
	}

	g1fb, err := FromBytes(gbf, swapEndianness)

	if err != nil {
		t.Fatal("TestGuidToFromBytes: FromBytes failed for guid " + gs)
	}

	if !g1fb.Equal(g) {
		t.Fatal("TestGuidToFromBytes: FromBytes test compare failed for guid " + gs)
	}
}

func BenchmarkEqualityBaseline(b *testing.B) {
	list := []string{gs1, gs2, gs3, gs4, gs5, gs6, gsz}
	glist := [7]Guid{}
	for i := range list {
		glist[i], _ = Parse(list[i])
	}
	b.ResetTimer()
	for range b.N {
		equal := true
		a, b := glist[0], glist[1]
		for k := range 16 {
			if a[k] != b[k] {
				equal = false
				break
			}
		}
		_ = equal
	}
}

func BenchmarkEqualityCurrent(b *testing.B) {
	list := []string{gs1, gs2, gs3, gs4, gs5, gs6, gsz}
	glist := [7]Guid{}
	for i := range list {
		glist[i], _ = Parse(list[i])
	}
	b.ResetTimer()
	for range b.N {
		a1 := (*uint64)(unsafe.Pointer(&glist[0][0]))
		a2 := (*uint64)(unsafe.Pointer(&glist[0][8]))
		b1 := (*uint64)(unsafe.Pointer(&glist[1][0]))
		b2 := (*uint64)(unsafe.Pointer(&glist[1][8]))
		equal := *a1 == *b1 && *a2 == *b2
		_ = equal
	}
}

func BenchmarkEqualityDirect(b *testing.B) {
	list := []string{gs1, gs2, gs3, gs4, gs5, gs6, gsz}
	glist := [7]Guid{}
	for i := range list {
		glist[i], _ = Parse(list[i])
	}
	b.ResetTimer()
	for range b.N {
		equal := glist[0] == glist[1]
		_ = equal
	}
}
