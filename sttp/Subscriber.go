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
	"sync"
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"

	//lint:ignore ST1001 -- static access to transport namespace desirable here
	. "github.com/sttp/goapi/sttp/transport"
)

// Subscriber defines the primary functionality of an STTP data subscription.
// Struct implementations of this interface that embed the SubscriberBase as a
// composite field will inherit a default implementation of all required interface
// methods. This allows implementations to focus only on needed functionality.
type Subscriber interface {
	// StatusMessage handles informational message logging.
	StatusMessage(message string)
	// ErrorMessage handles error message logging.
	ErrorMessage(message string)
	// ReceivedMetadata handles reception of the metadata response.
	ReceivedMetadata(metadata []byte)
	// SubscriptionUpdated handles notifications that a new SignalIndexCache has been received.
	SubscriptionUpdated(signalIndexCache *SignalIndexCache)
	// DataStartTime handles notifications of first received measurement. This can be useful in
	// cases where SubscriptionInfo.IncludeTime has been set to false.
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

	// AutoRequestMetadata defines the flag that determines if metadata should be
	// automatically requested upon successful connection. When true, metadata will
	// be requested upon connection before subscription; otherwise, any metadata
	// operations must be handled manually.
	AutoRequestMetadata bool

	// AutoSubscribe defines the flag that determines if subscription should be
	// handled automatically upon successful connection. When AutoRequestMetadata
	// is true and AutoSubscribe is true, subscription will occur after reception
	// of metadata. When AutoRequestMetadata is false and AutoSubscribe is true,
	// subscription will occur at successful connection. When AutoSubscribe is
	// false, any subscribe operations must be handled manually.
	AutoSubscribe bool

	// CompressPayloadData determines whether payload data is compressed.
	CompressPayloadData bool

	// CompressMetadata determines whether the metadata transfer is compressed.
	CompressMetadata bool

	// CompressSignalIndexCache determines whether the signal index cache is compressed.
	CompressSignalIndexCache bool

	// MetadataFilters defines any filters to be applied to incoming metadata to reduce total
	// received metadata. Each filter expression should be separated by semi-colons.
	MetadataFilters string

	// Version defines the target STTP protocol version. This currently defaults to 2.
	Version byte

	sub         Subscriber      // Reference to consumer Subscriber implementation
	ds          *DataSubscriber // Reference to internal DataSubscriber instance
	consoleLock sync.Mutex      // Simple lock to synchronize console writes
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
		AutoRequestMetadata:      true,
		AutoSubscribe:            true,
		CompressPayloadData:      true,
		CompressMetadata:         true,
		CompressSignalIndexCache: true,
		Version:                  2,
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

// IsConnected determines if Subscriber is currently connected to a data publisher.
func (sb *SubscriberBase) IsConnected() bool {
	return sb.dataSubscriber().IsConnected()
}

// IsSubscribed determines if Subscriber is currently subscribed to a data stream.
func (sb *SubscriberBase) IsSubscribed() bool {
	return sb.dataSubscriber().IsSubscribed()
}

// Subscription gets subscription related settings for Subscriber.
func (sb *SubscriberBase) Subscription() *SubscriptionInfo {
	return sb.dataSubscriber().Subscription()
}

// GetSignalIndexCache gets the active signal index cache.
func (sb *SubscriberBase) ActiveSignalIndexCache() *SignalIndexCache {
	return sb.dataSubscriber().ActiveSignalIndexCache()
}

// SubscriberID gets the subscriber ID as assigned by the data publisher upon receipt of the SignalIndexCache.
func (sb *SubscriberBase) SubscriberID() guid.Guid {
	return sb.dataSubscriber().SubscriberID()
}

// TotalCommandChannelBytesReceived gets the total number of bytes received via the command channel since last connection.
func (sb *SubscriberBase) TotalCommandChannelBytesReceived() uint64 {
	return sb.dataSubscriber().TotalCommandChannelBytesReceived()
}

// TotalDataChannelBytesReceived gets the total number of bytes received via the data channel since last connection.
func (sb *SubscriberBase) TotalDataChannelBytesReceived() uint64 {
	return sb.dataSubscriber().TotalDataChannelBytesReceived()
}

// TotalMeasurementsReceived gets the total number of measurements received since last subscription.
func (sb *SubscriberBase) TotalMeasurementsReceived() uint64 {
	return sb.dataSubscriber().TotalMeasurementsReceived()
}

// LookupMetadata gets the MeasurementMetadata for the specified signalID from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (sb *SubscriberBase) LookupMetadata(signalID guid.Guid) *MeasurementMetadata {
	return sb.dataSubscriber().LookupMetadata(signalID)
}

// Metadata gets the MeasurementMetadata associated with a measurement from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (sb *SubscriberBase) Metadata(measurement *Measurement) *MeasurementMetadata {
	return sb.dataSubscriber().Metadata(measurement)
}

// AdjustedValue gets the Value of a Measurement with any linear adjustments applied from the
// measurement's Adder and Multiplier metadata, if found.
func (sb *SubscriberBase) AdjustedValue(measurement *Measurement) float64 {
	return sb.dataSubscriber().AdjustedValue(measurement)
}

// RequestMetadata sends a request to the data publisher indicating that the Subscriber would
// like new metadata. Any defined MetadataFilters will be included in request.
func (sb *SubscriberBase) RequestMetadata() {
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

// Subscribe sends a request to the data publisher indicating that the Subscriber would
// like to start receiving streaming data. Subscribe parameters are controlled by the
// SubscriptionInfo fields available through the Subscription receiver method.
func (sb *SubscriberBase) Subscribe() {
	sb.dataSubscriber().Subscribe()
}

// Unsubscribe sends a request to the data publisher indicating that the Subscriber would
// like to stop receiving streaming data.
func (sb *SubscriberBase) Unsubscribe() {
	sb.dataSubscriber().Unsubscribe()
}

// Connect starts the connection cycle to an STTP publisher. When AutoReconnect is true, the connection
// will automatically be retried when the connection drops. If AutoRequestMetadata is true, then upon
// successful connection, meta-data will be requested. When AutoRequestMetadata is true and AutoSubscribe
// is true, subscription will occur after reception of metadata. When AutoRequestMetadata is false and
// AutoSubscribe is true, subscription will occur at successful connection.
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
	ds.Version = sb.Version

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
		// after metadata is received the SubscriberInstance will then initiate subscribe;
		// otherwise, subscribe is initiated immediately (when auto subscribe requested)
		if sb.AutoRequestMetadata {
			sb.RequestMetadata()
		} else if sb.AutoSubscribe {
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
		// after metadata is received the SubscriberInstance will then initiate subscribe;
		// otherwise, subscribe is initiated immediately (when auto subscribe requested)
		if sb.AutoRequestMetadata {
			sb.RequestMetadata()
		} else if sb.AutoSubscribe {
			ds.Subscribe()
		}
	} else {
		ds.Disconnect()
		sb.StatusMessage("Connection retry attempts exceeded.")
	}
}

func (sb *SubscriberBase) handleMetadataReceived(metadata []byte) {
	sb.sub.ReceivedMetadata(metadata)

	if sb.AutoRequestMetadata && sb.AutoSubscribe {
		sb.dataSubscriber().Subscribe()
	}
}

func (sb *SubscriberBase) handleDataStartTime(startTime ticks.Ticks) {
	sb.sub.DataStartTime(ticks.ToTime(startTime))
}

func (sb *SubscriberBase) handleConfigurationChanged() {
	sb.sub.ConfigurationChanged()

	if sb.AutoRequestMetadata {
		sb.RequestMetadata()
	}
}

func (sb *SubscriberBase) handleProcessingComplete(message string) {
	sb.sub.StatusMessage(message)
	sb.sub.HistoricalReadComplete()
}

// SubscriberBase default implementation of Subscriber interface. Note that an OOP language
// would consider the following "overridable" methods - effect here is the same.

// StatusMessage implements the default handler for informational message logging.
// Default implementation synchronously writes output to stdio. Logging is recommended.
func (sb *SubscriberBase) StatusMessage(message string) {
	sb.consoleLock.Lock()
	defer sb.consoleLock.Unlock()
	fmt.Println(message)
}

// ErrorMessage implements the default handler for error message logging.
// Default implementation synchronously writes output to to stderr. Logging is recommended.
func (sb *SubscriberBase) ErrorMessage(message string) {
	sb.consoleLock.Lock()
	defer sb.consoleLock.Unlock()
	fmt.Fprintln(os.Stderr, message)
}

// ReceivedMetadata implements the default handler for reception of the metadata response.
func (sb *SubscriberBase) ReceivedMetadata(metadata []byte) {
}

// SubscriptionUpdated implements the default handler for notifications that a new
// SignalIndexCache has been received.
func (sb *SubscriberBase) SubscriptionUpdated(signalIndexCache *SignalIndexCache) {
}

// DataStartTime implements the default handler for notifications of first received measurement.
func (sb *SubscriberBase) DataStartTime(startTime time.Time) {
}

// ConfigurationChanged implements the default handler for notifications that the
// data publisher configuration has changed.
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

// HistoricalReadComplete implements the default handler for notification that temporal processing
// has completed, i.e., the end of a historical playback data stream has been reached.
func (sb *SubscriberBase) HistoricalReadComplete() {
}

// ConnectionEstablished implements the default handler for notification that a connection has been established.
// Default implementation simply writes connection feedback to StatusMessage handler.
func (sb *SubscriberBase) ConnectionEstablished() {
	con := sb.dataSubscriber().Connector()
	sb.sub.StatusMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " established.")
}

// ConnectionTerminated implements the default handler for notification that a connection has been terminated.
// Default implementation simply writes connection terminated feedback to ErrorMessage handler.
func (sb *SubscriberBase) ConnectionTerminated() {
	con := sb.dataSubscriber().Connector()
	sb.sub.ErrorMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " terminated.")
}
