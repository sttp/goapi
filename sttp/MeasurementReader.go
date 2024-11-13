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
	"context"

	"github.com/sttp/goapi/sttp/transport"
)

// MeasurementReader defines an STTP measurement reader.
type MeasurementReader struct {
	current chan *transport.Measurement
}

func newMeasurementReader(parent *Subscriber) *MeasurementReader {
	reader := &MeasurementReader{
		current: make(chan *transport.Measurement),
	}

	parent.SetNewMeasurementsReceiver(func(measurements *[]transport.Measurement) {
		for _, m := range *measurements {
			reader.current <- &m
		}
	})

	return reader
}

// NextMeasurement blocks current thread until a new measurement arrives or provided context is completed.
// Returns tuple of measurement and completed state. Completed state flag will be false if a measurement
// was received; otherwise, state flag will be true along with a nil measurement when context is done.
func (mr *MeasurementReader) NextMeasurement(ctx context.Context) (*transport.Measurement, bool) {
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case <-ctx.Done():
		return nil, true
	case measurement := <-mr.current:
		return measurement, false
	}
}

// Close closes the measurement reader channel.
func (mr *MeasurementReader) Close() {
	close(mr.current)
}
