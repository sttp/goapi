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
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"

	//lint:ignore ST1001 -- static access to transport namespace desirable here
	. "github.com/sttp/goapi/sttp/transport"
)

// Subscriber defines the primary functionality of an STTP data subscription.
type Subscriber interface {
	// StatusMessage handles informational message logging.
	StatusMessage(message string)
	// ErrorMessage handles error message logging.
	ErrorMessage(message string)
	// ReceivedMetadata handles reception of the metadata response.
	ReceivedMetadata(metadata []byte)
	// SubscriptionUpdated handles notifications that a new SignalIndexCache has been received.
	SubscriptionUpdated(signalIndexCache *SignalIndexCache)
	// DataStartTime handles notifications of first received measurement.
	DataStartTime(startTime time.Time)
	// ConfigurationChanged handles notifications that the publisher configuration has changed.
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

// SubscriberBase provides the default functionality for a Subscriber implementation.
type SubscriberBase struct {
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

	// CompressPayloadData determines whether payload data is compressed.
	CompressPayloadData bool

	// CompressMetadata determines whether the metadata transfer is compressed.
	CompressMetadata bool

	// CompressSignalIndexCache determines whether the signal index cache is compressed.
	CompressSignalIndexCache bool

	// MetadataFilters defines any filters to be applied to incoming metadata to reduce total
	// received metadata. Each filter expression should be separated by semi-colons.
	MetadataFilters string

	sub Subscriber      // Reference to consumer Subscriber implementation
	ds  *DataSubscriber // Reference to internal DataSubscriber instance
}

// NewSubscriberBase creates a new SubscriberBase with specified Subscriber.
func NewSubscriberBase(subscriber Subscriber) SubscriberBase {
	return SubscriberBase{
		sub:                      subscriber,
		ds:                       NewDataSubscriber(),
		MaxRetries:               -1,
		RetryInterval:            1000,
		MaxRetryInterval:         30000,
		AutoReconnect:            true,
		AutoParseMetadata:        true,
		CompressPayloadData:      true,
		CompressMetadata:         true,
		CompressSignalIndexCache: true,
	}
}

// Dispose cleanly shuts down a DataSubscriber that is no longer being used, e.g.,
// during a normal application exit.
func (sb *SubscriberBase) Dispose() {
	if sb.ds != nil {
		sb.ds.Dispose()
	}
}

// dataSubscriber gets a reference to the internal DataSubscriber instance.
func (sb *SubscriberBase) dataSubscriber() *DataSubscriber {
	if sb.ds == nil {
		panic("Internal DataSubscriber instance has not been initialized. Make sure to use NewSubscriberBase.")
	}

	return sb.ds
}

// Subscription gets subscription related settings.
func (sb *SubscriberBase) Subscription() *SubscriptionInfo {
	ds := sb.dataSubscriber()
	return ds.Subscription()
}

// GetSignalIndexCache gets the active signal index cache.
func (sb *SubscriberBase) ActiveSignalIndexCache() *SignalIndexCache {
	ds := sb.dataSubscriber()
	return ds.ActiveSignalIndexCache()
}

// SubscriberID gets the subscriber ID as assigned by the DataPublisher upon receipt of the SignalIndexCache.
func (sb *SubscriberBase) SubscriberID() guid.Guid {
	ds := sb.dataSubscriber()
	return ds.SubscriberID()
}

// TotalCommandChannelBytesReceived gets the total number of bytes received via the command channel since last connection.
func (sb *SubscriberBase) TotalCommandChannelBytesReceived() uint64 {
	ds := sb.dataSubscriber()
	return ds.TotalCommandChannelBytesReceived()
}

// TotalDataChannelBytesReceived gets the total number of bytes received via the data channel since last connection.
func (sb *SubscriberBase) TotalDataChannelBytesReceived() uint64 {
	ds := sb.dataSubscriber()
	return ds.TotalDataChannelBytesReceived()
}

// TotalMeasurementsReceived gets the total number of measurements received since last subscription.
func (sb *SubscriberBase) TotalMeasurementsReceived() uint64 {
	ds := sb.dataSubscriber()
	return ds.TotalMeasurementsReceived()
}

// Connect starts the connection cycle to an STTP publisher. Upon connection, meta-data will be requested,
// when received, a subscription will be established.
func (sb *SubscriberBase) Connect() {
	ds := sb.dataSubscriber()
	con := ds.Connector()
	sub := sb.sub

	// Set connection properties
	con.Hostname = sb.Hostname
	con.Port = sb.Port
	con.MaxRetries = sb.MaxRetries
	con.RetryInterval = sb.RetryInterval
	con.MaxRetryInterval = sb.MaxRetryInterval
	con.AutoReconnect = sb.AutoReconnect

	ds.CompressPayloadData = sb.CompressPayloadData
	ds.CompressMetadata = sb.CompressMetadata
	ds.CompressSignalIndexCache = sb.CompressSignalIndexCache

	// Register direct Subscriber interface callbacks
	con.ErrorMessageCallback = sub.ErrorMessage
	ds.StatusMessageCallback = sub.StatusMessage
	ds.ErrorMessageCallback = sub.ErrorMessage
	ds.ConnectionTerminatedCallback = sub.ConnectionTerminated
	ds.SubscriptionUpdatedCallback = sub.SubscriptionUpdated
	ds.NewMeasurementsCallback = sub.ReceivedNewMeasurements
	ds.NewBufferBlocksCallback = sub.ReceivedNewBufferBlocks
	ds.NotificationReceivedCallback = sub.ReceivedNotification

	// Register callbacks with intermediate handlers
	con.ReconnectCallback = sb.handleReconnect
	ds.MetadataReceivedCallback = sb.handleMetadataReceived
	ds.DataStartTimeCallback = sb.handleDataStartTime
	ds.ConfigurationChangedCallback = sb.handleConfigurationChanged
	ds.ProcessingCompleteCallback = sb.handleProcessingComplete

	var status ConnectStatusEnum

	// Connect and subscribe to publisher
	if status = con.Connect(ds); status == ConnectStatus.Success {
		sub.ConnectionEstablished()

		// If automatically parsing metadata, request metadata upon successful connection,
		// after metadata is handled the SubscriberInstance will then initiate subscribe;
		// otherwise, initiate subscribe immediately
		if sb.AutoParseMetadata {
			sb.sendMetadataRefreshCommand()
		} else {
			ds.Subscribe()
		}
	} else if status == ConnectStatus.Failed {
		sb.ErrorMessage("All connection attempts failed")
	}
}

// Disconnect disconnects from an STTP publisher.
func (sb *SubscriberBase) Disconnect() {
	sb.dataSubscriber().Disconnect()
}

// Intermediate callback handlers:

func (sb *SubscriberBase) handleReconnect(ds *DataSubscriber) {
	if ds.IsConnected() {
		sb.sub.ConnectionEstablished()

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

func (sb *SubscriberBase) handleMetadataReceived(metadata []byte) {
	sb.sub.ReceivedMetadata(metadata)

	if sb.AutoParseMetadata {
		sb.dataSubscriber().Subscribe()
	}
}

func (sb *SubscriberBase) handleDataStartTime(startTime ticks.Ticks) {
	sb.sub.DataStartTime(ticks.ToTime(startTime))
}

func (sb *SubscriberBase) handleConfigurationChanged() {
	sb.sub.ConfigurationChanged()

	if sb.AutoParseMetadata {
		sb.sendMetadataRefreshCommand()
	}
}

func (sb *SubscriberBase) handleProcessingComplete(message string) {
	sb.sub.StatusMessage(message)
	sb.sub.HistoricalReadComplete()
}

func (sb *SubscriberBase) sendMetadataRefreshCommand() {
	ds := sb.dataSubscriber()

	if len(sb.MetadataFilters) == 0 {
		ds.SendServerCommand(ServerCommand.MetadataRefresh)
		return
	}

	filters := ds.EncodeString(sb.MetadataFilters)
	buffer := make([]byte, 4+len(filters))

	binary.BigEndian.PutUint32(buffer, uint32(len(filters)))
	copy(buffer[4:], filters)

	ds.SendServerCommandWithPayload(ServerCommand.MetadataRefresh, buffer)
}

// SubscriberBase default implementation of Subscriber interface:

// StatusMessage implements the default handler for informational message logging.
// Default implementation simply writes to stdio. Logging is recommended.
func (sb *SubscriberBase) StatusMessage(message string) {
	fmt.Println(message)
}

// ErrorMessage implements the default handler for error message logging.
// Default implementation simply writes to stderr. Logging is recommended.
func (sb *SubscriberBase) ErrorMessage(message string) {
	fmt.Fprintln(os.Stderr, message)
}

// ReceivedMetadata implements the default handler for reception of the metadata response.
func (sb *SubscriberBase) ReceivedMetadata(metadata []byte) {
}

// SubscriptionUpdated implements the default handler for notifications that a new SignalIndexCache has been received.
func (sb *SubscriberBase) SubscriptionUpdated(signalIndexCache *SignalIndexCache) {
}

// DataStartTime implements the default handler for notifications of first received measurement.
func (sb *SubscriberBase) DataStartTime(startTime time.Time) {
}

// ConfigurationChanged implements the default handler for notifications that the publisher configuration has changed.
func (sb *SubscriberBase) ConfigurationChanged() {
}

// ReceivedNewMeasurements implements the default handler for reception of new measurements.
func (sb *SubscriberBase) ReceivedNewMeasurements(measurements []Measurement) {
}

// ReceivedNewBufferBlocks implements the default handler for reception of new buffer blocks.
func (sb *SubscriberBase) ReceivedNewBufferBlocks(bufferBlocks []BufferBlock) {
}

// ReceivedNotification implements the default handler for reception of a notification.
func (sb *SubscriberBase) ReceivedNotification(notification string) {
}

// HistoricalReadComplete implements the default handler for notification that temporal processing has completed,
// i.e., the end of a historical playback data stream has been reached.
func (sb *SubscriberBase) HistoricalReadComplete() {
}

// ConnectionEstablished implements the default handler for notification that a connection has been established.
func (sb *SubscriberBase) ConnectionEstablished() {
	con := sb.ds.Connector()
	sb.sub.StatusMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " established.")
}

// ConnectionTerminated implements the default handler for notification that a connection has been terminated.
func (sb *SubscriberBase) ConnectionTerminated() {
	con := sb.ds.Connector()
	sb.sub.ErrorMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " terminated.")
}
