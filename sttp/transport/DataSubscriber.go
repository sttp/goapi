//******************************************************************************************************
//  DataSubscriber.go - Gbtc
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
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sttp/goapi/sttp/format"
	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/thread"
	"github.com/sttp/goapi/sttp/ticks"
	"github.com/sttp/goapi/sttp/transport/tssc"
	"github.com/sttp/goapi/sttp/version"
	"github.com/tevino/abool/v2"
)

// DataSubscriber represents a client subscription for an STTP connection.
type DataSubscriber struct {
	subscription SubscriptionInfo
	subscriberID guid.Guid
	encoding     OperationalEncodingEnum
	connector    *SubscriberConnector
	connected    abool.AtomicBool
	subscribed   abool.AtomicBool

	commandChannelSocket         net.Conn
	commandChannelResponseThread *thread.Thread
	readBuffer                   []byte
	reader                       *bufio.Reader
	writeBuffer                  []byte
	dataChannelSocket            net.Conn
	dataChannelResponseThread    *thread.Thread

	assigningHandlerMutex sync.RWMutex

	connectActionMutex          sync.Mutex
	connectionTerminationThread *thread.Thread

	disconnectThread      *thread.Thread
	disconnectThreadMutex sync.Mutex
	disconnecting         abool.AtomicBool
	disconnected          abool.AtomicBool
	disposing             abool.AtomicBool

	// Statistics counters
	totalCommandChannelBytesReceived uint64
	totalDataChannelBytesReceived    uint64
	totalMeasurementsReceived        uint64

	// StatusMessageCallback is called when a informational message should be logged.
	StatusMessageCallback func(string)

	// ErrorMessageCallback is called when an error message should be logged.
	ErrorMessageCallback func(string)

	// ConnectionTerminatedCallback is called when DataSubscriber terminates its connection.
	ConnectionTerminatedCallback func()

	// AutoReconnectCallback is called when DataSubscriber automatically reconnects.
	AutoReconnectCallback func()

	// MetadataReceivedCallback is called when DataSubscriber receives a metadata response.
	MetadataReceivedCallback func([]byte)

	// SubscriptionUpdatedCallback is called when DataSubscriber receives a new signal index cache.
	SubscriptionUpdatedCallback func(signalIndexCache *SignalIndexCache)

	// DataStartTimeCallback is called with timestamp of first received measurement in a subscription.
	DataStartTimeCallback func(ticks.Ticks)

	// ConfigurationChangedCallback is called when the DataPublisher sends a notification that configuration has changed.
	ConfigurationChangedCallback func()

	// NewMeasurementsCallback is called when DataSubscriber receives a set of new measurements from the DataPublisher.
	NewMeasurementsCallback func([]Measurement)

	// NewBufferBlocksCallback is called when DataSubscriber receives a set of new buffer block measurements from the DataPublisher.
	NewBufferBlocksCallback func([]BufferBlock)

	// ProcessingCompleteCallback is called when the DataPublished sends a notification that temporal processing has completed,
	// i.e., the end of a historical playback data stream has been reached.
	ProcessingCompleteCallback func(string)

	// NotificationReceivedCallback is called when the DataPublisher sends a notification that requires receipt.
	NotificationReceivedCallback func(string)

	// CompressPayloadData determines whether payload data is compressed, defaults to TSSC.
	CompressPayloadData bool

	// CompressMetadata determines whether the metadata transfer is compressed, defaults to GZip.
	CompressMetadata bool

	// CompressSignalIndexCache determines whether the signal index cache is compressed, defaults to GZip.
	CompressSignalIndexCache bool

	// Version defines the STTP protocol version used by this library.
	Version byte

	// SwapGuidEndianness determines if Guid wire serialization should swap endianness. This should only be enabled for
	// implementations using non-RFC Guid byte ordering, i.e., little-endian. Default to false.
	SwapGuidEndianness bool

	// STTPSourceInfo defines the STTP library API title as identification information of DataSubscriber to a DataPublisher.
	STTPSourceInfo string

	// STTPVersionInfo defines the STTP library API version as identification information of DataSubscriber to a DataPublisher.
	STTPVersionInfo string

	// STTPUpdatedOnInfo defines when the STTP library API was last updated as identification information of DataSubscriber to a DataPublisher.
	STTPUpdatedOnInfo string

	// Measurement parsing
	metadataRequested       time.Time
	measurementRegistry     map[guid.Guid]*MeasurementMetadata
	signalIndexCache        [2]*SignalIndexCache
	signalIndexCacheMutex   sync.Mutex
	cacheIndex              int32
	timeIndex               int32
	baseTimeOffsets         [2]int64
	keyIVs                  [][][]byte
	lastMissingCacheWarning ticks.Ticks
	tsscResetRequested      abool.AtomicBool
	tsscLastOOSReport       time.Time
	tsscLastOOSReportMutex  sync.Mutex

	bufferBlockExpectedSequenceNumber uint32
	bufferBlockCache                  []BufferBlock
}

// NewDataSubscriber creates a new DataSubscriber.
func NewDataSubscriber() *DataSubscriber {
	ds := &DataSubscriber{
		subscription:             SubscriptionInfo{IncludeTime: true},
		encoding:                 OperationalEncoding.UTF8,
		connector:                &SubscriberConnector{},
		readBuffer:               make([]byte, maxPacketSize),
		writeBuffer:              make([]byte, maxPacketSize),
		CompressPayloadData:      true, // Defaults to TSSC
		CompressMetadata:         true, // Defaults to Gzip
		CompressSignalIndexCache: true, // Defaults to Gzip
		Version:                  2,
		SwapGuidEndianness:       false,
		STTPSourceInfo:           version.STTPSource,
		STTPVersionInfo:          version.STTPVersion,
		STTPUpdatedOnInfo:        version.STTPUpdatedOn,
		measurementRegistry:      make(map[guid.Guid]*MeasurementMetadata),
		signalIndexCache:         [2]*SignalIndexCache{NewSignalIndexCache(), NewSignalIndexCache()},
	}

	ds.connectionTerminationThread = thread.NewThread(func() {
		ds.disconnect(false, true)
	})

	return ds
}

// Dispose cleanly shuts down a DataSubscriber that is no longer being used, e.g.,
// during a normal application exit.
func (ds *DataSubscriber) Dispose() {
	ds.disposing.Set()
	ds.connector.Cancel()
	ds.disconnect(true, false)

	// Allow a moment for connection terminated event to complete
	waitTimer := time.NewTimer(time.Duration(10) * time.Millisecond)
	<-waitTimer.C
}

// BeginCallbackAssignment informs DataSubscriber that a callback change has been initiated.
func (ds *DataSubscriber) BeginCallbackAssignment() {
	ds.assigningHandlerMutex.Lock()
}

// BeginCallbackSync begins a callback synchronization operation.
func (ds *DataSubscriber) BeginCallbackSync() {
	ds.assigningHandlerMutex.RLock()
}

// EndCallbackSync ends a callback synchronization operation.
func (ds *DataSubscriber) EndCallbackSync() {
	ds.assigningHandlerMutex.RUnlock()
}

// EndCallbackAssignment informs DataSubscriber that a callback change has been completed.
func (ds *DataSubscriber) EndCallbackAssignment() {
	ds.assigningHandlerMutex.Unlock()
}

// IsConnected determines if a DataSubscriber is currently connected to a DataPublisher.
func (ds *DataSubscriber) IsConnected() bool {
	return ds.connected.IsSet()
}

// IsSubscribed determines if a DataSubscriber is currently subscribed to a data stream.
func (ds *DataSubscriber) IsSubscribed() bool {
	return ds.subscribed.IsSet()
}

// EncodeString encodes an STTP string according to the defined operational modes.
func (ds *DataSubscriber) EncodeString(data string) []byte {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if ds.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return []byte(data)
}

// DecodeString decodes an STTP string according to the defined operational modes.
func (ds *DataSubscriber) DecodeString(data []byte) string {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if ds.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return string(data)
}

// LookupMetadata gets the MeasurementMetadata for the specified signalID from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (ds *DataSubscriber) LookupMetadata(signalID guid.Guid) *MeasurementMetadata {
	metadata, ok := ds.measurementRegistry[signalID]

	if !ok {
		metadata = &MeasurementMetadata{
			SignalID:   signalID,
			Multiplier: 1.0,
		}

		ds.measurementRegistry[metadata.SignalID] = metadata
	}

	return metadata
}

// Metadata gets the MeasurementMetadata associated with a measurement from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (ds *DataSubscriber) Metadata(measurement *Measurement) *MeasurementMetadata {
	return ds.LookupMetadata(measurement.SignalID)
}

// AdjustedValue gets the Value of a Measurement with any linear adjustments applied from the
// measurement's Adder and Multiplier metadata, if found.
func (ds *DataSubscriber) AdjustedValue(measurement *Measurement) float64 {
	metadata, ok := ds.measurementRegistry[measurement.SignalID]

	if ok {
		return measurement.Value*metadata.Multiplier + metadata.Adder
	}

	return measurement.Value
}

// Connect requests the the DataSubscriber initiate a connection to the DataPublisher.
func (ds *DataSubscriber) Connect(hostName string, port uint16) error {
	// User requests to connection are not an auto-reconnect attempt
	return ds.connect(hostName, port, false)
}

func (ds *DataSubscriber) connect(hostName string, port uint16, autoReconnecting bool) error {
	if ds.connected.IsSet() {
		return errors.New("subscriber is already connected; disconnect first")
	}

	// Make sure any pending disconnect has completed to make sure socket is closed
	ds.disconnectThreadMutex.Lock()
	disconnectThread := ds.disconnectThread
	ds.disconnectThreadMutex.Unlock()

	if disconnectThread != nil {
		disconnectThread.Join()
	}

	// Let any pending connect or disconnect operation complete before new connect,
	// this prevents destruction disconnect before connection is completed
	ds.connectActionMutex.Lock()
	defer ds.connectActionMutex.Unlock()

	var err error

	ds.disconnected.UnSet()
	ds.subscribed.UnSet()

	atomic.StoreUint64(&ds.totalCommandChannelBytesReceived, 0)
	atomic.StoreUint64(&ds.totalDataChannelBytesReceived, 0)
	atomic.StoreUint64(&ds.totalMeasurementsReceived, 0)

	ds.keyIVs = nil
	ds.bufferBlockExpectedSequenceNumber = 0
	ds.measurementRegistry = make(map[guid.Guid]*MeasurementMetadata)

	if !autoReconnecting {
		ds.connector.ResetConnection()
	}

	ds.connector.connectionRefused.UnSet()

	// TODO: Add TLS implementation options
	// TODO: Add reverse (server-based) connection options, see:
	// https://sttp.github.io/documentation/reverse-connections/

	ds.commandChannelSocket, err = net.Dial("tcp", hostName+":"+strconv.Itoa(int(port)))

	if err == nil {
		ds.commandChannelResponseThread = thread.NewThread(ds.runCommandChannelResponseThread)
		ds.commandChannelResponseThread.Start()

		ds.connected.Set()
		ds.lastMissingCacheWarning = 0
		ds.sendOperationalModes()
	}

	return err
}

// Subscribe notifies the DataPublisher that a DataSubscriber would like to start receiving streaming data.
func (ds *DataSubscriber) Subscribe() error {
	if ds.connected.IsNotSet() {
		return errors.New("subscriber is not connected; cannot subscribe")
	}

	// Make sure to unsubscribe before attempting another
	// subscription so we don't leave UDP sockets open
	if ds.subscribed.IsSet() {
		ds.Unsubscribe()
	}

	atomic.StoreUint64(&ds.totalMeasurementsReceived, 0)

	var connectionBuilder strings.Builder

	connectionBuilder.WriteString("throttled=")
	connectionBuilder.WriteString(strconv.FormatBool(ds.subscription.Throttled))
	connectionBuilder.WriteString(";publishInterval=")
	connectionBuilder.WriteString(strconv.FormatFloat(ds.subscription.PublishInterval, 'f', 6, 64))
	connectionBuilder.WriteString(";includeTime=")
	connectionBuilder.WriteString(strconv.FormatBool(ds.subscription.IncludeTime))
	connectionBuilder.WriteString(";processingInterval=")
	connectionBuilder.WriteString(strconv.FormatInt(int64(ds.subscription.ProcessingInterval), 10))
	connectionBuilder.WriteString(";useMillisecondResolution=")
	connectionBuilder.WriteString(strconv.FormatBool(ds.subscription.UseMillisecondResolution))
	connectionBuilder.WriteString(";requestNaNValueFilter")
	connectionBuilder.WriteString(strconv.FormatBool(ds.subscription.RequestNaNValueFilter))
	connectionBuilder.WriteString(";assemblyInfo={source=")
	connectionBuilder.WriteString(ds.STTPSourceInfo)
	connectionBuilder.WriteString(";version=")
	connectionBuilder.WriteString(ds.STTPVersionInfo)
	connectionBuilder.WriteString(";updatedOn=")
	connectionBuilder.WriteString(ds.STTPUpdatedOnInfo)
	connectionBuilder.WriteString("}")

	if len(ds.subscription.FilterExpression) > 0 {
		connectionBuilder.WriteString(";filterExpression={")
		connectionBuilder.WriteString(ds.subscription.FilterExpression)
		connectionBuilder.WriteString("}")
	}

	if ds.subscription.UdpDataChannel {
		udpPort := strconv.Itoa(int(ds.subscription.DataChannelLocalPort))
		udpAddr, err := net.ResolveUDPAddr("udp", ":"+udpPort)

		if err != nil {
			return errors.New("failed to resolve UDP address for port " + udpPort + ": " + err.Error())
		}

		ds.dataChannelSocket, err = net.ListenUDP("udp", udpAddr)

		if err != nil {
			return errors.New("failed to open UDP socket for port " + udpPort + ": " + err.Error())
		}

		ds.dataChannelResponseThread = thread.NewThread(ds.runDataChannelResponseThread)
		ds.dataChannelResponseThread.Start()

		connectionBuilder.WriteString(";dataChannel={localport=")
		connectionBuilder.WriteString(udpPort)
		connectionBuilder.WriteString("}")
	}

	if len(ds.subscription.StartTime) > 0 {
		connectionBuilder.WriteString(";startTimeConstraint=")
		connectionBuilder.WriteString(ds.subscription.StartTime)
	}

	if len(ds.subscription.StopTime) > 0 {
		connectionBuilder.WriteString(";stopTimeConstraint=")
		connectionBuilder.WriteString(ds.subscription.StopTime)
	}

	if len(ds.subscription.ConstraintParameters) > 0 {
		connectionBuilder.WriteString(";timeConstraintParameters=")
		connectionBuilder.WriteString(ds.subscription.ConstraintParameters)
	}

	if len(ds.subscription.ExtraConnectionStringParameters) > 0 {
		connectionBuilder.WriteRune(';')
		connectionBuilder.WriteString(ds.subscription.ExtraConnectionStringParameters)
	}

	connectionString := connectionBuilder.String()
	length := uint32(len(connectionString))
	buffer := make([]byte, 5+length)

	buffer[0] = byte(DataPacketFlags.Compact)
	binary.BigEndian.PutUint32(buffer[1:], length)
	copy(buffer[5:], connectionString)

	ds.SendServerCommandWithPayload(ServerCommand.Subscribe, buffer)

	// Reset TSSC decompressor on successful (re)subscription
	ds.tsscLastOOSReportMutex.Lock()
	ds.tsscLastOOSReport = time.Time{}
	ds.tsscLastOOSReportMutex.Unlock()
	ds.tsscResetRequested.Set()

	return nil
}

// Unsubscribe notifies the DataPublisher that a DataSubscriber would like to stop receiving streaming data.
func (ds *DataSubscriber) Unsubscribe() {
	if ds.connected.IsSet() {
		return
	}

	ds.SendServerCommand(ServerCommand.Unsubscribe)

	ds.disconnecting.Set()

	if ds.dataChannelSocket != nil {
		if err := ds.dataChannelSocket.Close(); err != nil {
			ds.dispatchErrorMessage("Exception while disconnecting data subscriber UDP data channel: " + err.Error())
		}
	}

	if ds.dataChannelResponseThread != nil {
		ds.dataChannelResponseThread.Join()
	}

	ds.disconnecting.UnSet()
}

// Disconnect initiates a DataSubscriber disconnect sequence.
func (ds *DataSubscriber) Disconnect() {
	if ds.disconnecting.IsSet() {
		return
	}

	// Disconnect method executes shutdown on a separate thread without stopping to prevent
	// issues where user may call disconnect method from a dispatched event thread. Also,
	// user requests to disconnect are not an auto-reconnect attempt
	ds.disconnect(false, false)
}

func (ds *DataSubscriber) disconnect(joinThread bool, autoReconnecting bool) {
	// Check if disconnect thread is running or subscriber has already disconnected
	if ds.disconnecting.IsSet() {
		if !autoReconnecting && ds.disconnecting.IsSet() && ds.disconnected.IsNotSet() {
			ds.connector.Cancel()
		}

		ds.disconnectThreadMutex.Lock()
		disconnectThread := ds.disconnectThread
		ds.disconnectThreadMutex.Unlock()

		if joinThread && ds.disconnected.IsNotSet() && disconnectThread != nil {
			disconnectThread.Join()
		}

		return
	}

	// Notify running threads that the subscriber is disconnecting, i.e., disconnect thread is active
	ds.disconnecting.Set()
	ds.connected.UnSet()
	ds.subscribed.UnSet()

	disconnectThread := thread.NewThread(func() {
		ds.runDisconnectThread(autoReconnecting)
	})

	ds.disconnectThreadMutex.Lock()
	ds.disconnectThread = disconnectThread
	ds.disconnectThreadMutex.Unlock()

	disconnectThread.Start()

	if joinThread {
		disconnectThread.Join()
	}
}

func (ds *DataSubscriber) runDisconnectThread(autoReconnecting bool) {
	// Let any pending connect operation complete before disconnect - prevents destruction disconnect before connection is completed
	if !autoReconnecting {
		ds.connector.Cancel()
		ds.connectionTerminationThread.Join()
		ds.connectActionMutex.Lock()
	}

	// Release queues and close sockets so that threads can shut down gracefully
	if ds.commandChannelSocket != nil {
		if err := ds.commandChannelSocket.Close(); err != nil {
			ds.dispatchErrorMessage("Exception while disconnecting data subscriber TCP command channel: " + err.Error())
		}
	}

	if ds.dataChannelSocket != nil {
		if err := ds.dataChannelSocket.Close(); err != nil {
			ds.dispatchErrorMessage("Exception while disconnecting data subscriber UDP data channel: " + err.Error())
		}
	}

	// Join with all threads to guarantee their completion before returning control to the caller
	if ds.commandChannelResponseThread != nil {
		ds.commandChannelResponseThread.Join()
	}

	if ds.dataChannelResponseThread != nil {
		ds.dataChannelResponseThread.Join()
	}

	// Notify consumers of disconnect
	ds.BeginCallbackSync()

	if ds.ConnectionTerminatedCallback != nil {
		ds.ConnectionTerminatedCallback()
	}

	ds.EndCallbackSync()

	// Disconnect complete
	ds.disconnected.Set()
	ds.disconnecting.UnSet()

	if autoReconnecting {
		// Handling auto-connect callback separately from connection terminated callback
		// since they serve two different use cases and current implementation does not
		// support multiple callback registrations
		ds.BeginCallbackSync()

		if ds.AutoReconnectCallback != nil && ds.disposing.IsNotSet() {
			ds.AutoReconnectCallback()
		}

		ds.EndCallbackSync()

	} else {
		ds.connectActionMutex.Unlock()
	}
}

// Dispatcher for connection terminated. This is called from its own separate thread
// in order to cleanly shut down the subscriber in case the connection was terminated
// by the peer. Additionally, this allows the user to automatically reconnect in their
// callback function without having to spawn their own separate thread.
func (ds *DataSubscriber) dispatchConnectionTerminated() {
	ds.connectionTerminationThread.TryStart()
}

func (ds *DataSubscriber) dispatchStatusMessage(message string) {
	ds.BeginCallbackSync()

	if ds.StatusMessageCallback != nil {
		go ds.StatusMessageCallback(message)
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) dispatchErrorMessage(message string) {
	ds.BeginCallbackSync()

	if ds.ErrorMessageCallback != nil {
		go ds.ErrorMessageCallback(message)
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) runCommandChannelResponseThread() {
	ds.reader = bufio.NewReader(ds.commandChannelSocket)

	for ds.connected.IsSet() {
		ds.readPayloadHeader(io.ReadFull(ds.reader, ds.readBuffer[:payloadHeaderSize]))
	}
}

func (ds *DataSubscriber) readPayloadHeader(bytesTransferred int, err error) {
	if ds.disconnecting.IsSet() {
		return
	}

	if err != nil {
		// Read error, connection may have been closed by peer; terminate connection
		ds.dispatchConnectionTerminated()
		return
	}

	// Gather statistics
	atomic.AddUint64(&ds.totalCommandChannelBytesReceived, uint64(bytesTransferred))

	packetSize := binary.BigEndian.Uint32(ds.readBuffer)

	if int(packetSize) > cap(ds.readBuffer) {
		ds.readBuffer = make([]byte, packetSize)
	}

	// Read packet (payload body)
	// This read method is guaranteed not to return until the
	// requested size has been read or an error has occurred.
	ds.readPacket(io.ReadFull(ds.reader, ds.readBuffer[:packetSize]))
}

func (ds *DataSubscriber) readPacket(bytesTransferred int, err error) {
	if ds.disconnecting.IsSet() {
		return
	}

	if err != nil {
		// Read error, connection may have been closed by peer; terminate connection
		ds.dispatchConnectionTerminated()
		return
	}

	// Gather statistics
	atomic.AddUint64(&ds.totalCommandChannelBytesReceived, uint64(bytesTransferred))

	// Process response
	ds.processServerResponse(ds.readBuffer[:bytesTransferred])
}

// If the user defines a separate UDP channel for their
// subscription, data packets get handled from this thread.
func (ds *DataSubscriber) runDataChannelResponseThread() {
	reader := bufio.NewReader(ds.dataChannelSocket)
	buffer := make([]byte, maxPacketSize)

	for ds.connected.IsSet() {
		length, err := reader.Read(buffer)

		if err != nil {
			ds.dispatchErrorMessage("Error reading data from command channel: " + err.Error())
			break
		}

		// Gather statistics
		atomic.AddUint64(&ds.totalDataChannelBytesReceived, uint64(length))

		// Process response
		ds.processServerResponse(buffer[:length])
	}
}

func (ds *DataSubscriber) processServerResponse(buffer []byte) {
	// Note: internal payload size at buffer[2:6] ignored - future versions of STTP will likely exclude this
	data := buffer[responseHeaderSize:]
	responseCode := ServerResponseEnum(buffer[0])
	commandCode := ServerCommandEnum(buffer[1])

	switch responseCode {
	case ServerResponse.Succeeded:
		ds.handleSucceeded(commandCode, data)
	case ServerResponse.Failed:
		ds.handleFailed(commandCode, data)
	case ServerResponse.DataPacket:
		ds.handleDataPacket(data)
	case ServerResponse.DataStartTime:
		ds.handleDataStartTime(data)
	case ServerResponse.ProcessingComplete:
		ds.handleProcessingComplete(data)
	case ServerResponse.UpdateSignalIndexCache:
		ds.handleUpdateSignalIndexCache(data)
	case ServerResponse.UpdateBaseTimes:
		ds.handleUpdateBaseTimes(data)
	case ServerResponse.UpdateCipherKeys:
		ds.handleUpdateCipherKeys(data)
	case ServerResponse.ConfigurationChanged:
		ds.handleConfigurationChanged()
	case ServerResponse.BufferBlock:
		ds.handleBufferBlock(data)
	case ServerResponse.Notification:
		ds.handleNotification(data)
	case ServerResponse.NoOP:
		// NoOP handled
	default:
		ds.dispatchErrorMessage("Encountered unexpected server response code: " + responseCode.String())
	}
}

func (ds *DataSubscriber) handleSucceeded(commandCode ServerCommandEnum, data []byte) {
	switch commandCode {
	case ServerCommand.MetadataRefresh:
		ds.handleMetadataRefresh(data)
	case ServerCommand.Subscribe, ServerCommand.Unsubscribe:
		if commandCode == ServerCommand.Subscribe {
			ds.subscribed.Set()
		} else {
			ds.subscribed.UnSet()
		}

		// Fallthrough on these messages because there is
		// still an associated message to be processed.
		fallthrough
	case ServerCommand.RotateCipherKeys, ServerCommand.UpdateProcessingInterval:
		// Each of these responses come with a message that will
		// be delivered to the user via the status message callback.
		var message strings.Builder
		message.WriteString("Received success code in response to server command: ")
		message.WriteString(commandCode.String())

		if len(data) > 0 {
			message.WriteRune('\n')
			message.Write(data)
		}

		ds.dispatchStatusMessage(message.String())
	default:
		// If we don't know what the message is, we can't interpret
		// the data sent with the packet. Deliver an error message
		// to the user via the error message callback.
		ds.dispatchErrorMessage("Received success code in response to unknown server command: " + commandCode.String())
	}
}

func (ds *DataSubscriber) handleFailed(commandCode ServerCommandEnum, data []byte) {
	var message strings.Builder

	if commandCode == ServerCommand.Connect {
		ds.connector.connectionRefused.Set()
	} else {
		message.WriteString("Received failure code in response to server command: ")
		message.WriteString(commandCode.String())
	}

	if len(data) > 0 {
		if message.Len() > 0 {
			message.WriteRune('\n')
		}

		message.Write(data)		
	}

	if message.Len() > 0 {
		ds.dispatchErrorMessage(message.String())
	}
}

func (ds *DataSubscriber) handleMetadataRefresh(data []byte) {
	ds.BeginCallbackSync()

	if ds.MetadataReceivedCallback != nil {
		if ds.CompressMetadata {
			ds.dispatchStatusMessage(fmt.Sprintf("Received %s bytes of metadata in %s seconds. Decompressing...", format.Int(len(data)), format.Float(time.Since(ds.metadataRequested).Seconds(), 3)))

			decompressStarted := time.Now()
			var err error

			if data, err = decompressGZip(data); err != nil {
				ds.dispatchErrorMessage("Failed to decompress received metadata: " + err.Error())
				return
			}

			ds.dispatchStatusMessage(fmt.Sprintf("Decompressed %s bytes of metadata in %s seconds. Parsing...", format.Int(len(data)), format.Float(time.Since(decompressStarted).Seconds(), 3)))
		} else {
			ds.dispatchStatusMessage(fmt.Sprintf("Received %s bytes of metadata in %s seconds. Parsing...", format.Int(len(data)), format.Float(time.Since(ds.metadataRequested).Seconds(), 3)))
		}

		go ds.MetadataReceivedCallback(data)
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) handleDataStartTime(data []byte) {
	ds.BeginCallbackSync()

	if ds.DataStartTimeCallback != nil {
		// Do not use Go routine here, processing sequence may be important.
		// Execute callback directly from socket processing thread:
		ds.DataStartTimeCallback(ticks.Ticks(binary.BigEndian.Uint64(data)))
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) handleProcessingComplete(data []byte) {
	ds.BeginCallbackSync()

	if ds.ProcessingCompleteCallback != nil {
		go ds.ProcessingCompleteCallback(ds.DecodeString(data))
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) handleUpdateSignalIndexCache(data []byte) {
	if data == nil {
		return
	}

	version := ds.Version
	var cacheIndex int32

	// Get active cache index
	if version > 1 {
		if data[0] > 0 {
			cacheIndex = 1
		}

		data = data[1:]
	}

	if ds.CompressSignalIndexCache {
		var err error

		if data, err = decompressGZip(data); err != nil {
			ds.dispatchErrorMessage("Failed to decompress received signal index cache: " + err.Error())
			return
		}
	}

	signalIndexCache := NewSignalIndexCache()
	err := signalIndexCache.decode(ds, data, &ds.subscriberID)

	if err != nil {
		ds.dispatchErrorMessage("Failed to parse signal index cache: " + err.Error())
		return
	}

	ds.signalIndexCacheMutex.Lock()
	ds.signalIndexCache[cacheIndex] = signalIndexCache
	ds.cacheIndex = cacheIndex
	ds.signalIndexCacheMutex.Unlock()

	if version > 1 {
		ds.SendServerCommand(ServerCommand.ConfirmSignalIndexCache)
	}

	ds.BeginCallbackSync()

	if ds.SubscriptionUpdatedCallback != nil {
		go ds.SubscriptionUpdatedCallback(signalIndexCache)
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) handleUpdateBaseTimes(data []byte) {
	if data == nil {
		return
	}

	timeIndex := int32(binary.BigEndian.Uint32(data))

	if timeIndex != 0 {
		timeIndex = 1
	}

	ds.timeIndex = timeIndex

	var baseTimeOffsets [2]int64

	baseTimeOffsets[0] = int64(binary.BigEndian.Uint64(data[4:]))
	baseTimeOffsets[1] = int64(binary.BigEndian.Uint64(data[12:]))

	ds.baseTimeOffsets = baseTimeOffsets

	timestamp, _ := ticks.ToTime(ticks.Ticks(ds.baseTimeOffsets[ds.timeIndex^1])).MarshalText()
	ds.dispatchStatusMessage("Received new base time offset from publisher: " + string(timestamp))
}

func (ds *DataSubscriber) handleUpdateCipherKeys(data []byte) {
	// Deserialize new cipher keys
	keyIVs := make([][][]byte, 2)
	keyIVs[evenKey] = make([][]byte, 2)
	keyIVs[oddKey] = make([][]byte, 2)

	// Move past active cipher index (not currently used anywhere else)
	var index uint32 = 1

	// Read even key size
	bufferLen := binary.BigEndian.Uint32(data[index:])
	index += 4

	// Read even key
	keyIVs[evenKey][keyIndex] = make([]byte, bufferLen)
	copy(keyIVs[evenKey][keyIndex], data[index:])
	index += bufferLen

	// Read even initialization vector size
	bufferLen = binary.BigEndian.Uint32(data[index:])
	index += 4

	// Read even initialization vector
	keyIVs[evenKey][ivIndex] = make([]byte, bufferLen)
	copy(keyIVs[evenKey][ivIndex], data[index:])
	index += bufferLen

	// Read odd key size
	bufferLen = binary.BigEndian.Uint32(data[index:])
	index += 4

	// Read odd key
	keyIVs[oddKey][keyIndex] = make([]byte, bufferLen)
	copy(keyIVs[oddKey][keyIndex], data[index:])
	index += bufferLen

	// Read odd initialization vector size
	bufferLen = binary.BigEndian.Uint32(data[index:])
	index += 4

	// Read odd initialization vector
	keyIVs[oddKey][ivIndex] = make([]byte, bufferLen)
	copy(keyIVs[oddKey][ivIndex], data[index:])
	//index += bufferLen

	// Exchange keys
	ds.keyIVs = keyIVs

	ds.dispatchStatusMessage("Successfully established new cipher keys for UDP data packet transmissions.")
}

func (ds *DataSubscriber) handleConfigurationChanged() {
	ds.dispatchStatusMessage("Received notification from publisher that configuration has changed.")

	ds.BeginCallbackSync()

	if ds.ConfigurationChangedCallback != nil {
		go ds.ConfigurationChangedCallback()
	}

	ds.EndCallbackSync()
}

func (ds *DataSubscriber) handleDataPacket(data []byte) {
	dataPacketFlags := DataPacketFlagsEnum(data[0])
	compressed := dataPacketFlags&DataPacketFlags.Compressed > 0
	compact := dataPacketFlags&DataPacketFlags.Compact > 0

	if !compressed && !compact {
		ds.dispatchErrorMessage("Go implementation of STTP only supports compact or compressed data packet encoding - disconnecting.")
		ds.dispatchConnectionTerminated()
		return
	}

	data = data[1:]

	if ds.keyIVs != nil {
		// Get a local copy keyIVs - these can change at any time
		keyIVs := ds.keyIVs
		var cipherIndex int
		var err error

		if dataPacketFlags&DataPacketFlags.CipherIndex > 0 {
			cipherIndex = 1
		}

		data, err = decipherAES(keyIVs[cipherIndex][keyIndex], keyIVs[cipherIndex][ivIndex], data)

		if err != nil {
			ds.dispatchErrorMessage("Failed to decrypt data packet - disconnecting: " + err.Error())
			ds.dispatchConnectionTerminated()
			return
		}
	}

	count := binary.BigEndian.Uint32(data)
	measurements := make([]Measurement, count)
	var cacheIndex int

	if dataPacketFlags&DataPacketFlags.CacheIndex > 0 {
		cacheIndex = 1
	}

	ds.signalIndexCacheMutex.Lock()
	signalIndexCache := ds.signalIndexCache[cacheIndex]
	ds.signalIndexCacheMutex.Unlock()

	if compressed {
		ds.parseTSSCMeasurements(signalIndexCache, data[4:], measurements)
	} else {
		ds.parseCompactMeasurements(signalIndexCache, data[4:], measurements)
	}

	ds.BeginCallbackSync()

	if ds.NewMeasurementsCallback != nil {
		// Do not use Go routine here, processing sequence may be important.
		// Execute callback directly from socket processing thread:
		ds.NewMeasurementsCallback(measurements)
	}

	ds.EndCallbackSync()

	atomic.AddUint64(&ds.totalMeasurementsReceived, uint64(count))
}

func (ds *DataSubscriber) parseTSSCMeasurements(signalIndexCache *SignalIndexCache, data []byte, measurements []Measurement) {
	decoder := signalIndexCache.tsscDecoder
	var newDecoder bool

	// Use TSSC to decompress measurements
	if decoder == nil {
		signalIndexCache.tsscDecoder = tssc.NewDecoder(signalIndexCache.MaxSignalIndex())
		decoder = signalIndexCache.tsscDecoder
		decoder.SequenceNumber = 0
		newDecoder = true
	}

	if data[0] != 85 {
		ds.dispatchErrorMessage("TSSC version not recognized - disconnecting. Received version: " + strconv.Itoa(int(data[0])))
		ds.dispatchConnectionTerminated()
		return
	}

	sequenceNumber := binary.BigEndian.Uint16(data[1:])

	if sequenceNumber == 0 {
		if !newDecoder {
			if decoder.SequenceNumber > 0 {
				ds.dispatchStatusMessage("TSSC algorithm reset before sequence number: " + strconv.Itoa(int(decoder.SequenceNumber)))
			}

			signalIndexCache.tsscDecoder = tssc.NewDecoder(signalIndexCache.MaxSignalIndex())
			decoder = signalIndexCache.tsscDecoder
			decoder.SequenceNumber = 0
		}

		ds.tsscResetRequested.UnSet()
		ds.tsscLastOOSReportMutex.Lock()
		ds.tsscLastOOSReport = time.Time{}
		ds.tsscLastOOSReportMutex.Unlock()
	}

	if decoder.SequenceNumber != sequenceNumber {
		if ds.tsscResetRequested.IsNotSet() {
			ds.tsscLastOOSReportMutex.Lock()

			if time.Since(ds.tsscLastOOSReport).Seconds() > 2.0 {
				ds.dispatchErrorMessage("TSSC is out of sequence. Expecting: " + strconv.Itoa(int(decoder.SequenceNumber)) + ", received: " + strconv.Itoa(int(sequenceNumber)))
				ds.tsscLastOOSReport = time.Now()
			}

			ds.tsscLastOOSReportMutex.Unlock()
		}

		// Ignore packets until the reset has occurred
		return
	}

	decoder.SetBuffer(data[3:])

	var id int32
	var timestamp int64
	var stateFlags uint32
	var value float32
	var err error

	ok := true
	index := 0

	for ok {
		if ok, err = decoder.TryGetMeasurement(&id, &timestamp, &stateFlags, &value); ok {
			measurements[index] = Measurement{
				SignalID:  signalIndexCache.SignalID(id),
				Value:     float64(value),
				Timestamp: ticks.Ticks(timestamp),
				Flags:     StateFlagsEnum(stateFlags),
			}

			index++
		}
	}

	if err != nil {
		ds.dispatchErrorMessage("Failed to parse TSSC measurements - disconnecting: " + err.Error())
		ds.dispatchConnectionTerminated()
		return
	}

	decoder.SequenceNumber++

	// Do not increment to 0 on roll-over
	if decoder.SequenceNumber == 0 {
		decoder.SequenceNumber = 1
	}
}

func (ds *DataSubscriber) parseCompactMeasurements(signalIndexCache *SignalIndexCache, data []byte, measurements []Measurement) {
	if signalIndexCache.Count() == 0 {
		if ds.lastMissingCacheWarning+missingCacheWarningInterval < ticks.UtcNow() {
			// Warning message for missing signal index cache
			if ds.lastMissingCacheWarning != 0 {
				ds.dispatchStatusMessage("Signal index cache has not arrived. No compact measurements can be parsed.")
			}

			ds.lastMissingCacheWarning = ticks.UtcNow()
		}
		
		return
	}

	useMillisecondResolution := ds.subscription.UseMillisecondResolution
	includeTime := ds.subscription.IncludeTime
	index := 0

	for i := 0; i < len(measurements); i++ {
		// Deserialize compact measurement format
		compactMeasurement := NewCompactMeasurement(signalIndexCache, includeTime, useMillisecondResolution, &ds.baseTimeOffsets)
		bytesDecoded, err := compactMeasurement.Decode(data[index:])

		if err != nil {
			ds.dispatchErrorMessage("Failed to parse compact measurements - disconnecting: " + err.Error())
			ds.dispatchConnectionTerminated()
			return
		}

		index += bytesDecoded
		measurements[i] = compactMeasurement.Measurement
	}
}

func (ds *DataSubscriber) handleBufferBlock(data []byte) {
	// Buffer block received - wrap as a BufferBlockMeasurement and expose back to consumer
	sequenceNumber := binary.BigEndian.Uint32(data)
	bufferCacheIndex := int(sequenceNumber - ds.bufferBlockExpectedSequenceNumber)
	var signalIndexCacheIndex int32

	if ds.Version > 1 && data[4:][0] > 0 {
		signalIndexCacheIndex = 1
	}

	// Check if this buffer block has already been processed (e.g., mistaken retransmission due to timeout)
	if bufferCacheIndex >= 0 && (bufferCacheIndex >= len(ds.bufferBlockCache) || ds.bufferBlockCache[bufferCacheIndex].Buffer == nil) {
		// Send confirmation that buffer block is received
		ds.SendServerCommandWithPayload(ServerCommand.ConfirmBufferBlock, data[:4])

		if ds.Version > 1 {
			data = data[5:]
		} else {
			data = data[4:]
		}

		// Get measurement key from signal index cache
		signalIndex := int32(binary.BigEndian.Uint32(data))

		ds.signalIndexCacheMutex.Lock()
		signalIndexCache := ds.signalIndexCache[signalIndexCacheIndex]
		ds.signalIndexCacheMutex.Unlock()

		signalID := signalIndexCache.SignalID(signalIndex)
		bufferBlockMeasurement := BufferBlock{SignalID: signalID}

		// Determine if this is the next buffer block in the sequence
		if sequenceNumber == ds.bufferBlockExpectedSequenceNumber {
			bufferBlockMeasurements := make([]BufferBlock, 1+len(ds.bufferBlockCache))
			var i int

			// Add the buffer block measurement to the list of measurements to be published
			bufferBlockMeasurements[0] = bufferBlockMeasurement
			ds.bufferBlockExpectedSequenceNumber++

			// Add cached buffer block measurements to the list of measurements to be published
			for i = 1; i < len(ds.bufferBlockCache); i++ {
				if ds.bufferBlockCache[i].Buffer == nil {
					break
				}

				bufferBlockMeasurements[i] = ds.bufferBlockCache[i]
				ds.bufferBlockExpectedSequenceNumber++
			}

			// Remove published buffer block measurements from the buffer block queue
			if len(ds.bufferBlockCache) > 0 {
				ds.bufferBlockCache = ds.bufferBlockCache[i:]
			}

			// Publish buffer block measurements
			ds.BeginCallbackSync()

			if ds.NewBufferBlocksCallback != nil {
				// Do not use Go routine here, processing sequence may be important.
				// Execute callback directly from socket processing thread:
				ds.NewBufferBlocksCallback(bufferBlockMeasurements)
			}

			ds.EndCallbackSync()
		} else {
			// Ensure that the list has at least as many elements as it needs to cache this measurement.
			// This edge case handles possible dropouts and/or out of order packet deliver when data
			// transport is UDP - this use case is not expected when using a TCP only connection.
			for i := len(ds.bufferBlockCache); i <= bufferCacheIndex; i++ {
				ds.bufferBlockCache = append(ds.bufferBlockCache, BufferBlock{})
			}

			// Insert this buffer block into the proper location in the list
			ds.bufferBlockCache[bufferCacheIndex] = bufferBlockMeasurement
		}
	}
}

func (ds *DataSubscriber) handleNotification(data []byte) {
	// Skip the 4-byte hash and decode notification message
	message := ds.DecodeString(data[4:])

	ds.dispatchStatusMessage("NOTIFICATION: " + message)

	ds.BeginCallbackSync()

	if ds.NotificationReceivedCallback != nil {
		go ds.NotificationReceivedCallback(message)
	}

	ds.EndCallbackSync()

	// Send confirmation of receipt of the notification with 4-byte hash
	ds.SendServerCommandWithPayload(ServerCommand.ConfirmNotification, data[:4])
}

// SendServerCommand sends a server command code to the DataPublisher with no payload.
func (ds *DataSubscriber) SendServerCommand(commandCode ServerCommandEnum) {
	ds.SendServerCommandWithPayload(commandCode, nil)
}

// SendServerCommandWithMessage sends a server command code to the DataPublisher along with the specified string message as payload.
func (ds *DataSubscriber) SendServerCommandWithMessage(commandCode ServerCommandEnum, message string) {
	ds.SendServerCommandWithPayload(commandCode, ds.EncodeString(message))
}

// SendServerCommandWithPayload sends a server command code to the DataPublisher along with the specified data payload.
func (ds *DataSubscriber) SendServerCommandWithPayload(commandCode ServerCommandEnum, data []byte) {
	if ds.connected.IsNotSet() {
		return
	}

	var packetSize uint32 = uint32(len(data)) + 1
	var commandBufferSize uint32 = packetSize + payloadHeaderSize

	if int(commandBufferSize) > cap(ds.writeBuffer) {
		ds.writeBuffer = ds.writeBuffer[:commandBufferSize]
	}

	// Insert packet size
	binary.BigEndian.PutUint32(ds.writeBuffer, packetSize)

	// Insert command code
	ds.writeBuffer[4] = byte(commandCode)

	if data != nil {
		for i := 0; i < len(data); i++ {
			ds.writeBuffer[5+i] = data[i]
		}
	}

	if commandCode == ServerCommand.MetadataRefresh {
		// Track start time of metadata request to calculate round-trip receive time
		ds.metadataRequested = time.Now()
	}

	if _, err := ds.commandChannelSocket.Write(ds.writeBuffer[:commandBufferSize]); err != nil {
		// Write error, connection may have been closed by peer; terminate connection
		ds.dispatchErrorMessage("Failed to send server command - disconnecting: " + err.Error())
		ds.dispatchConnectionTerminated()
	}
}

func (ds *DataSubscriber) sendOperationalModes() {
	var operationalModes OperationalModesEnum = OperationalModesEnum(CompressionModes.GZip)

	operationalModes |= OperationalModes.VersionMask & OperationalModesEnum(ds.Version)
	operationalModes |= OperationalModesEnum(ds.encoding)

	// TSSC compression only works with stateful connections
	if ds.CompressPayloadData && !ds.subscription.UdpDataChannel {
		operationalModes |= OperationalModes.CompressPayloadData | OperationalModesEnum(CompressionModes.TSSC)
	}

	if ds.CompressMetadata {
		operationalModes |= OperationalModes.CompressMetadata
	}

	if ds.CompressSignalIndexCache {
		operationalModes |= OperationalModes.CompressSignalIndexCache
	}

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, uint32(operationalModes))

	ds.SendServerCommandWithPayload(ServerCommand.DefineOperationalModes, buffer)
}

// Subscription gets the SubscriptionInfo associated with this DataSubscriber.
func (ds *DataSubscriber) Subscription() *SubscriptionInfo {
	return &ds.subscription
}

// Connector gets the SubscriberConnector associated with this DataSubscriber.
func (ds *DataSubscriber) Connector() *SubscriberConnector {
	return ds.connector
}

// ActiveSignalIndexCache gets the active signal index cache.
func (ds *DataSubscriber) ActiveSignalIndexCache() *SignalIndexCache {
	ds.signalIndexCacheMutex.Lock()
	signalIndexCache := ds.signalIndexCache[ds.cacheIndex]
	ds.signalIndexCacheMutex.Unlock()

	return signalIndexCache
}

// SubscriberID gets the subscriber ID as assigned by the DataPublisher upon receipt of the SignalIndexCache.
func (ds *DataSubscriber) SubscriberID() guid.Guid {
	return ds.subscriberID
}

// TotalCommandChannelBytesReceived gets the total number of bytes received via the command channel since last connection.
func (ds *DataSubscriber) TotalCommandChannelBytesReceived() uint64 {
	return atomic.LoadUint64(&ds.totalCommandChannelBytesReceived)
}

// TotalDataChannelBytesReceived gets the total number of bytes received via the data channel since last connection.
func (ds *DataSubscriber) TotalDataChannelBytesReceived() uint64 {
	if ds.subscription.UdpDataChannel {
		return atomic.LoadUint64(&ds.totalDataChannelBytesReceived)
	}

	return atomic.LoadUint64(&ds.totalCommandChannelBytesReceived)
}

// TotalMeasurementsReceived gets the total number of measurements received since last subscription.
func (ds *DataSubscriber) TotalMeasurementsReceived() uint64 {
	return atomic.LoadUint64(&ds.totalMeasurementsReceived)
}
