//******************************************************************************************************
//  AdvancedSubscribeReverse.go - Gbtc
//
//  Copyright © 2022, Grid Protection Alliance.  All Rights Reserved.
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
//  11/27/2022 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sttp/goapi/sttp"
	"github.com/sttp/goapi/sttp/format"
	"github.com/sttp/goapi/sttp/transport"
)

// AdvancedSubscriber represents an STTP data subscriber implementation.
type AdvancedSubscriber struct {
	sttp.Subscriber
	Config   *sttp.Config
	Settings *sttp.Settings

	lastMessage time.Time
}

// NewAdvancedSubscriber creates a new AdvancedSubscriber.
func NewAdvancedSubscriber() *AdvancedSubscriber {
	subscriber := &AdvancedSubscriber{
		Subscriber: *sttp.NewSubscriber(),
		Config:     sttp.NewConfig(),
		Settings:   sttp.NewSettings(),
	}

	subscriber.SetSubscriptionUpdatedReceiver(subscriber.subscriptionUpdated)
	subscriber.SetNewMeasurementsReceiver(subscriber.receivedNewMeasurements)
	subscriber.SetConnectionTerminatedReceiver(subscriber.connectionTerminated)

	return subscriber
}

func main() {
	address := parseCmdLineArgs()
	subscriber := NewAdvancedSubscriber()

	subscriber.Config.CompressPayloadData = false
	subscriber.Settings.UdpPort = 9600
	subscriber.Settings.UseMillisecondResolution = true

	subscriber.Subscribe("FILTER TOP 20 ActiveMeasurements WHERE True", subscriber.Settings)
	subscriber.Listen(address, subscriber.Config)
	defer subscriber.Close()

	readKey()
}

func (sub *AdvancedSubscriber) subscriptionUpdated(signalIndexCache *transport.SignalIndexCache) {
	sub.StatusMessage(fmt.Sprintf("Received signal index cache with %d mappings", signalIndexCache.Count()))
}

func (sub *AdvancedSubscriber) receivedNewMeasurements(measurements []transport.Measurement) {
	if time.Since(sub.lastMessage).Seconds() < 5.0 {
		return
	}

	defer func() { sub.lastMessage = time.Now() }()

	if sub.lastMessage.IsZero() {
		sub.StatusMessage("Receiving measurements...")
		return
	}

	var message strings.Builder

	message.WriteString(format.UInt64(sub.TotalMeasurementsReceived()))
	message.WriteString(" measurements received so far...\n")
	message.WriteString("Timestamp: ")
	message.WriteString(measurements[0].Timestamp.String())
	message.WriteRune('\n')
	message.WriteString("\tID\tSignal ID\t\t\t\tValue\n")

	for i := 0; i < len(measurements); i++ {
		measurement := measurements[i]
		metadata := sub.Metadata(&measurement)

		message.WriteRune('\t')
		message.WriteString(strconv.FormatUint(metadata.ID, 10))
		message.WriteRune('\t')
		message.WriteString(measurement.SignalID.String())
		message.WriteRune('\t')
		message.WriteString(format.Float(measurement.Value, 6))
		message.WriteRune('\n')
	}

	sub.StatusMessage(message.String())
}

func (sub *AdvancedSubscriber) connectionTerminated() {
	// Call default implementation which will display a connection terminated message to stderr
	sub.DefaultConnectionEstablishedReceiver()

	// Reset last message display time on disconnect
	sub.lastMessage = time.Time{}
}
