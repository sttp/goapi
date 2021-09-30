//******************************************************************************************************
//  MeasurementReader.go - Gbtc
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
//  09/30/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package sttp

import (
	"github.com/sttp/goapi/sttp/transport"
)

// MeasurementReader defines an STTP measurement reader.
type MeasurementReader struct {
	current    chan *transport.Measurement
	originalCT func()
}

func newMeasurementReader(parent *Subscriber) *MeasurementReader {
	reader := &MeasurementReader{
		current:    make(chan *transport.Measurement),
		originalCT: parent.connectionTerminatedReceiver,
	}

	parent.SetNewMeasurementsReceiver(reader.receivedNewMeasurements)
	parent.SetConnectionTerminatedReceiver(reader.connectionTerminated)

	return reader
}

// NextMeasurement blocks current thread until a new measurement arrives.
func (mr *MeasurementReader) NextMeasurement() *transport.Measurement {
	return <-mr.current
}

func (mr *MeasurementReader) receivedNewMeasurements(measurements []transport.Measurement) {
	for _, measurement := range measurements {
		mr.current <- &measurement
	}
}

func (mr *MeasurementReader) connectionTerminated() {
	close(mr.current)

	if mr.originalCT != nil {
		mr.originalCT()
	}
}
