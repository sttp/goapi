//******************************************************************************************************
//  Subscriber.go - Gbtc
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
//  09/16/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package sttp

import (
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"

	//lint:ignore ST1001 -- static access to transport namespace desirable here
	. "github.com/sttp/goapi/sttp/transport"
)

type Subscriber interface {
	// StatusMessage handles informational message logging.
	StatusMessage(message string)
	// ErrorMessage handles error message logging.
	ErrorMessage(message string)
	// ReceivedMetadata handles reception of the metadata response.
	ReceivedMetadata(metadata []byte)
	// SubscriptionUpdated handles
	SubscriptionUpdated(signalIndexCache *SignalIndexCache)
	// DataStartTime handles notifications of first received measurement.
	DataStartTime(startTime time.Time)
	// ConfigurationChanged handles
	ConfigurationChanged()
	// ReceivedNewMeasurements handles reception of new measurements.
	ReceivedNewMeasurements(measurements []Measurement)
	// ReceivedNewBufferBlocks handles reception of new buffer blocks.
	ReceivedNewBufferBlocks(bufferBlocks []BufferBlock)
	// ReceivedNotification handles reception of a notification.
	ReceivedNotification(notification string)
	// HistoricalReadComplete handles notification that temporal processing has completed,
	// i.e., the end of a historical playback data stream has been reached.
	HistoricalReadComplete()
	// ConnectionEstablished handles notification that a connection has been established.
	ConnectionEstablished()
	// ConnectionTerminated handles notification that a connection has been terminated.
	ConnectionTerminated()
}

type SubscriberBase struct {
	// Subscriber is the user implementation of the Subscriber interface.
	Subscriber

	// Hostname is the DataPublisher DNS name or IP.
	Hostname string

	// Port it the TCP/IP listening port of the DataPublisher.
	Port uint16

	// MaxRetries defines the maximum number of times to retry a connection.
	// Set value to -1 to retry infinitely.
	MaxRetries int32

	// RetryInterval defines the base retry interval, in milliseconds. Retries will
	// exponentially back-off starting from this interval.
	RetryInterval int32

	// MaxRetryInterval defines the maximum retry interval, in milliseconds.
	MaxRetryInterval int32

	// AutoReconnect defines flag that determines if connections should be
	// automatically reattempted.
	AutoReconnect bool

	// AutoParseMetadata defines the flag that determines if metadata should be
	// automatically parsed. When true, metadata will be requested upon connection
	// before subscription; otherwise, metadata will not be manually requested and
	// subscribe will happen upon connection.
	AutoParseMetadata bool

	// dataSubscriber is local DataSubscriber reference
	dataSubscriber *DataSubscriber
}

// NewSubscriberBase creates a new SubscriberBase for the specified Subscriber.
func NewSubscriberBase(subscriber Subscriber) SubscriberBase {
	return SubscriberBase{
		dataSubscriber: NewDataSubscriber(),
	}
}

// Dispose cleanly shuts down a DataSubscriber that is no longer being used, e.g.,
// during a normal application exit.
func (sb *SubscriberBase) Dispose() {
	if sb.dataSubscriber != nil {
		sb.dataSubscriber.Dispose()
	}
}

func (sb *SubscriberBase) getDataSubscriber() *DataSubscriber {
	if sb.dataSubscriber == nil {
		panic("DataSubscriber has not been initialized. Make sure to use NewSubscriberBase.")
	}

	return sb.dataSubscriber
}

// GetSubscription gets subscription related settings.
func (sb *SubscriberBase) GetSubscription() *SubscriptionInfo {
	ds := sb.getDataSubscriber()
	return ds.GetSubscription()
}

// GetSubscriberID gets the subscriber ID as assigned by the DataPublisher upon receipt of the SignalIndexCache.
func (sb *SubscriberBase) GetSubscriberID() guid.Guid {
	ds := sb.getDataSubscriber()
	return ds.GetSubscriberID()
}

// Connect starts the connection cycle to an STTP publisher. Upon connection, meta-data will be requested,
// when received, a subscription will be established.
func (sb *SubscriberBase) Connect() {
	dataSubscriber := sb.getDataSubscriber()
	connector := dataSubscriber.GetConnector()

	// Set connector properties
	connector.Hostname = sb.Hostname
	connector.Port = sb.Port
	connector.MaxRetries = sb.MaxRetries
	connector.RetryInterval = sb.RetryInterval
	connector.AutoReconnect = sb.AutoReconnect

	// Register callbacks
	connector.ErrorMessageCallback = sb.ErrorMessage
	connector.ReconnectCallback = sb.handleReconnect
	dataSubscriber.StatusMessageCallback = sb.StatusMessage
	dataSubscriber.ErrorMessageCallback = sb.ErrorMessage
	dataSubscriber.MetadataReceivedCallback = sb.ReceivedMetadata
	dataSubscriber.SubscriptionUpdatedCallback = sb.SubscriptionUpdated
	dataSubscriber.DataStartTimeCallback = sb.handleDataStartTime
	dataSubscriber.ConfigurationChangedCallback = sb.ConfigurationChanged
	dataSubscriber.NewMeasurementsCallback = sb.ReceivedNewMeasurements
	dataSubscriber.NewBufferBlocksCallback = sb.ReceivedNewBufferBlocks
	dataSubscriber.NotificationReceivedCallback = sb.ReceivedNotification
	dataSubscriber.ProcessingCompleteCallback = sb.HistoricalReadComplete

	// TODO: Complete operation
}

func (sb *SubscriberBase) handleReconnect(ds *DataSubscriber) {
	if ds.IsConnected() {
		sb.ConnectionEstablished()

		// If automatically parsing metadata, request metadata upon successful connection,
		// after metadata is handled the SubscriberInstance will then initiate subscribe;
		// otherwise, initiate subscribe immediately
		if sb.AutoParseMetadata {
			sb.sendMetadataRefreshCommand()
		} else {
			ds.Subscribe()
		}
	} else {
		ds.Disconnect()
		sb.StatusMessage("Connection retry attempts exceeded.")
	}
}

func (sb *SubscriberBase) handleDataStartTime(startTime ticks.Ticks) {
	sb.DataStartTime(ticks.ToTime(startTime))
}

func (sb *SubscriberBase) sendMetadataRefreshCommand() {

}

// DataStartTime defines the default handler for
func (sb *SubscriberBase) DataStartTime(startTime time.Time) {

}
