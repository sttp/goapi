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
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sttp/goapi/sttp"
	"github.com/sttp/goapi/sttp/transport"
)

// SimpleSubscriber is a simple STTP data subscriber implementation.
type SimpleSubscriber struct {
	sttp.SubscriberBase // Provides default implementation
}

// NewSimpleSubscriber creates a new SimpleSubscriber.
func NewSimpleSubscriber() *SimpleSubscriber {
	subscriber := &SimpleSubscriber{}
	subscriber.SubscriberBase = sttp.NewSubscriberBase(subscriber)
	return subscriber
}

func main() {
	hostname, port := parseCmdLineArgs()
	subscriber := NewSimpleSubscriber()
	subscription := subscriber.Subscription()

	subscriber.Hostname = hostname
	subscriber.Port = port
	subscriber.CompressPayloadData = false
	subscriber.Version = 1

	subscription.FilterExpression = "FILTER TOP 5 ActiveMeasurements WHERE True"

	subscriber.Connect()
	defer subscriber.Dispose()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadRune()
}

// ReceivedMetadata handles reception of the metadata response.
func (ss *SimpleSubscriber) ReceivedMetadata(metadata []byte) {
	ss.StatusMessage(fmt.Sprintf("Received %d bytes of metadata", len(metadata)))
}

// SubscriptionUpdated handles notifications that a new SignalIndexCache has been received.
func (ss *SimpleSubscriber) SubscriptionUpdated(signalIndexCache *transport.SignalIndexCache) {
	ss.StatusMessage(fmt.Sprintf("Received signal index cache with %d mappings", signalIndexCache.Count()))
}

var lastMessageDisplay time.Time

// ReceivedNewMeasurements handles reception of new measurements.
func (ss *SimpleSubscriber) ReceivedNewMeasurements(measurements []transport.Measurement) {

	if time.Since(lastMessageDisplay).Seconds() < 5.0 {
		return
	}

	defer func() { lastMessageDisplay = time.Now() }()

	if lastMessageDisplay.IsZero() {
		ss.StatusMessage("Receiving measurements...")
		return
	}

	var message strings.Builder

	message.WriteString(strconv.FormatUint(ss.TotalMeasurementsReceived(), 10))
	message.WriteString(" measurements received so far...\n")
	message.WriteString("Timestamp: ")
	message.WriteString(measurements[0].DateTime().Format("2006-01-02 15:04:05.999999999"))
	message.WriteRune('\n')
	message.WriteString("\tID\tSignal ID\t\t\t\tValue\n")

	for i := 0; i < len(measurements); i++ {
		measurement := measurements[i]
		metadata := ss.Metadata(&measurement)

		message.WriteRune('\t')
		message.WriteString(strconv.FormatUint(metadata.ID, 10))
		message.WriteRune('\t')
		message.WriteString(measurement.SignalID.String())
		message.WriteRune('\t')
		message.WriteString(strconv.FormatFloat(measurement.Value, 'f', 6, 64))
		message.WriteRune('\n')
	}

	ss.StatusMessage(message.String())
}

// ConnectionTerminated handles notification that a connection has been terminated.
func (ss *SimpleSubscriber) ConnectionTerminated() {
	// Call base implementation which will display a connection terminated message to stderr
	ss.SubscriberBase.ConnectionTerminated()

	// Reset last message display time on disconnect
	lastMessageDisplay = time.Time{}
}

func parseCmdLineArgs() (string, uint16) {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("    SimpleSubscribe HOSTNAME PORT")
		os.Exit(1)
	}

	hostname := args[1]
	port, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("Invalid port number \"%s\": %s\n", args[1], err.Error())
		os.Exit(2)
	}

	if port < 1 || port > math.MaxUint16 {
		fmt.Printf("Port number \"%s\" is out of range: must be 1 to %d\n", args[1], math.MaxUint16)
		os.Exit(2)
	}

	return hostname, uint16(port)
}
