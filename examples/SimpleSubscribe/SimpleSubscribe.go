//******************************************************************************************************
//  SimpleSubscribe.go - Gbtc
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
//  09/17/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/sttp/goapi/sttp"
)

func main() {
	address := parseCmdLineArgs()
	subscriber := sttp.NewSubscriber()
	defer subscriber.Close()

	// Start data read at each connection
	subscriber.SetConnectionEstablishedReceiver(func() {
		go func() {
			reader := subscriber.ReadMeasurements()
			var lastMessage time.Time

			for subscriber.IsConnected() {
				measurement, _ := reader.NextMeasurement(nil)

				if time.Since(lastMessage).Seconds() < 5.0 {
					continue
				} else if lastMessage.IsZero() {
					subscriber.StatusMessage("Receiving measurements...")
					lastMessage = time.Now()
					continue
				}

				var message strings.Builder

				message.WriteString(strconv.FormatUint(subscriber.TotalMeasurementsReceived(), 10))
				message.WriteString(" measurements received so far. Current measurement:\n    ")
				message.WriteString(measurement.String())

				subscriber.StatusMessage(message.String())
				lastMessage = time.Now()
			}
		}()
	})

	subscriber.Subscribe("FILTER TOP 20 ActiveMeasurements WHERE True", nil)
	subscriber.Dial(address, nil)

	readKey()
}
