//******************************************************************************************************
//  Measurement.go - Gbtc
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

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"
)

// Measurement defines a basic unit of data for use by the STTP API.
type Measurement struct {
	// Measurement's globally unique identifier.
	SignalID guid.Guid

	// Instantaneous value of the measurement.
	Value float64

	// The time, in ticks, that this measurement was taken.
	Timestamp ticks.Ticks

	// Flags indicating the state of the measurement as reported by the device that took it.
	Flags StateFlagsEnum
}

// TicksValue gets the integer-based time from a Measurement Ticks based timestamp, i.e.,
// the 62-bit time value excluding any reserved flags.
func (m *Measurement) TicksValue() int64 {
	return int64(m.Timestamp & ticks.ValueMask)
}

// DateTime gets a Measurement Ticks based timestamp as a standard Go Time value.
func (m *Measurement) DateTime() time.Time {
	return m.Timestamp.ToTime()
}

// String returns the string form of a Measurement value.
func (m *Measurement) String() string {
	return fmt.Sprintf("%s @ %s = %s (%s)",
		m.SignalID.String(),
		m.Timestamp.ShortTime(),
		strconv.FormatFloat(m.Value, 'f', 3, 64),
		m.Flags.String())
}
