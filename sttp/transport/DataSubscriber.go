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
	"io"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/thread"
	"github.com/sttp/goapi/sttp/ticks"
	"github.com/sttp/goapi/sttp/version"
)

// DataSubscriber represents a client subscription for an STTP connection.
type DataSubscriber struct {
	subscription SubscriptionInfo
	subscriberID guid.Guid
	encoding     OperationalEncodingEnum
	connector    SubscriberConnector
	connected    bool
	subscribed   bool

	commandChannelSocket         net.Conn
	commandChannelResponseThread *thread.Thread
	readBuffer                   []byte
	reader                       *bufio.Reader
	writeBuffer                  []byte
	dataChannelSocket            net.Conn
	dataChannelResponseThread    *thread.Thread

	connectActionMutex          sync.Mutex
	connectionTerminationThread *thread.Thread

	disconnectThread *thread.Thread
	disconnecting    bool
	disconnected     bool
	disposing        bool

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
	AutoReconnectCallback func(*DataSubscriber)

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
	ProcessingCompleteCallback func()

	// NotificationReceivedCallback is called when the DataPublisher sends a notification that requires receipt.
	NotificationReceivedCallback func(string)

	// CompressPayloadData determines whether payload data is compressed using TSSC.
	CompressPayloadData bool

	// CompressMetadata determines whether the metadata transfer is compressed using GZip.
	CompressMetadata bool

	// CompressSignalIndexCache determines whether the signal index cache is compressed using GZip.
	CompressSignalIndexCache bool

	// Version defines the STTP protocol version used by this library
	Version byte

	// STTPSourceInfo defines the STTP library API title as identification information of DataSubscriber to a DataPublisher.
	STTPSourceInfo string

	// STTPVersionInfo defines the STTP library API version as identification information of DataSubscriber to a DataPublisher.
	STTPVersionInfo string

	// STTPUpdatedOnInfo defines when the STTP library API was last updated as identification information of DataSubscriber to a DataPublisher.
	STTPUpdatedOnInfo string

	// Measurement parsing
	signalIndexCache        [2]*SignalIndexCache
	signalIndexCacheLock    sync.Mutex
	cacheIndex              int32
	timeIndex               int32
	baseTimeOffsets         [2]int64
	keyIVs                  [][][]byte
	lastMissingCacheWarning ticks.Ticks
	tsscResetRequested      bool
	tsscSequenceNumber      uint16
	//tsscDecoder      tssc.TSSCDecoder

	bufferBlockExpectedSequenceNumber uint32
	bufferBlockCache                  []BufferBlock
}

// NewDataSubscriber creates a new DataSubscriber.
func NewDataSubscriber() *DataSubscriber {
	return &DataSubscriber{
		subscription:             SubscriptionInfo{IncludeTime: true},
		encoding:                 OperationalEncoding.UTF8,
		connector:                SubscriberConnector{},
		readBuffer:               make([]byte, maxPacketSize),
		writeBuffer:              make([]byte, maxPacketSize),
		CompressPayloadData:      true, // Defaults to TSSC
		CompressMetadata:         true, // Defaults to Gzip
		CompressSignalIndexCache: true, // Defaults to Gzip
		Version:                  2,
		STTPSourceInfo:           version.Source,
		STTPVersionInfo:          version.Version,
		STTPUpdatedOnInfo:        version.UpdatedOn,
		signalIndexCache:         [2]*SignalIndexCache{NewSignalIndexCache(), NewSignalIndexCache()},
	}
}

// Dispose cleanly shuts down a DataSubscriber that is no longer being used, e.g.,
// during a normal application exit.
func (ds *DataSubscriber) Dispose() {
	ds.disposing = true
	ds.connector.Cancel()
	ds.disconnect(true, false)
}

// IsConnected determines if a DataSubscriber is currently connected to a DataPublisher.
func (ds *DataSubscriber) IsConnected() bool {
	return ds.connected
}

// IsSubscribed determines if a DataSubscriber is currently subscribed to a data stream.
func (ds *DataSubscriber) IsSubscribed() bool {
	return ds.subscribed
}

// DecodeString decodes an STTP string according to the defined operational modes.
func (ds *DataSubscriber) DecodeString(data []byte) string {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if ds.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	return string(data)
}

// Connect requests the the DataSubscriber initiate a connection to the DataPublisher.
func (ds *DataSubscriber) Connect(hostName string, port uint16) error {
	// User requests to connection are not an auto-reconnect attempt
	return ds.connect(hostName, port, false)
}

func (ds *DataSubscriber) connect(hostName string, port uint16, autoReconnecting bool) error {
	if ds.connected {
		panic("Subscriber is already connected; disconnect first")
	}

	// Let any pending connect or disconnect operation complete before new connect,
	// this prevents destruction disconnect before connection is completed
	ds.connectActionMutex.Lock()
	defer ds.connectActionMutex.Unlock()

	var err error

	ds.disconnected = false
	ds.subscribed = false
	ds.totalCommandChannelBytesReceived = 0
	ds.totalDataChannelBytesReceived = 0
	ds.totalMeasurementsReceived = 0
	ds.keyIVs = nil
	ds.bufferBlockExpectedSequenceNumber = 0

	if !autoReconnecting {
		ds.connector.ResetConnection()
	}

	ds.connector.connectionRefused = false

	// TODO: Add TLS implementation options
	// TODO: Add reverse (server-based) connection options, see:
	// https://sttp.github.io/documentation/reverse-connections/

	ds.commandChannelSocket, err = net.Dial("tcp", hostName+":"+strconv.Itoa(int(port)))

	if err == nil {
		ds.commandChannelResponseThread = thread.NewThread(ds.runCommandChannelResponseThread)
		ds.connected = true
		ds.lastMissingCacheWarning = 0
		ds.sendOperationalModes()
	}

	return err
}

// Subscribe notifies the DataPublisher that a DataSubscriber would like to start receiving streaming data.
func (ds *DataSubscriber) Subscribe() error {
	if !ds.connected {
		return errors.New("subscriber is not connected; cannot subscribe")
	}

	// Make sure to unsubscribe before attempting another
	// subscription so we don't leave UDP sockets open
	if ds.subscribed {
		ds.Unsubscribe()
	}

	ds.totalMeasurementsReceived = 0

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
			return errors.New("Failed to resolve UDP address for port " + udpPort + ": " + err.Error())
		}

		ds.dataChannelSocket, err = net.ListenUDP("udp", udpAddr)

		if err != nil {
			return errors.New("Failed to open UDP socket for port " + udpPort + ": " + err.Error())
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
	ds.tsscResetRequested = true

	return nil
}

// Unsubscribe notifies the DataPublisher that a DataSubscriber would like to stop receiving streaming data.
func (ds *DataSubscriber) Unsubscribe() {
	if ds.connected {
		return
	}

	ds.disconnecting = true

	if ds.dataChannelSocket != nil {
		if err := ds.dataChannelSocket.Close(); err != nil {
			ds.dispatchErrorMessage("Exception while disconnecting data subscriber UDP data channel: " + err.Error())
		}
	}

	if ds.dataChannelResponseThread != nil {
		ds.dataChannelResponseThread.Join()
	}

	ds.disconnecting = false

	ds.SendServerCommand(ServerCommand.Unsubscribe)
}

// Disconnect initiates a DataSubscriber disconnect sequence.
func (ds *DataSubscriber) Disconnect() {
	if ds.disconnecting {
		return
	}

	// Disconnect method executes shutdown on a separate thread without stopping to prevent
	// issues where user may call disconnect method from a dispatched event thread. Also,
	// user requests to disconnect are not an auto-reconnect attempt
	ds.disconnect(false, false)
}

func (ds *DataSubscriber) disconnect(joinThread bool, autoReconnecting bool) {
	// Check if disconnect thread is running or subscriber has already disconnected
	if ds.disconnecting {
		if !autoReconnecting && ds.disconnecting && !ds.disconnected {
			ds.connector.Cancel()
		}

		if joinThread && !ds.disconnected && ds.disconnectThread != nil {
			ds.disconnectThread.Join()
		}

		return
	}

	// Notify running threads that the subscriber is disconnecting, i.e., disconnect thread is active
	ds.disconnecting = true
	ds.connected = false
	ds.subscribed = false

	ds.disconnectThread = thread.NewThread(func() {
		ds.runDisconnectThread(autoReconnecting)
	})

	ds.disconnectThread.Start()

	if joinThread {
		ds.disconnectThread.Join()
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
	if ds.ConnectionTerminatedCallback != nil {
		ds.ConnectionTerminatedCallback()
	}

	// Disconnect complete
	ds.disconnected = true
	ds.disconnecting = false

	if autoReconnecting {
		// Handling auto-connect callback separately from connection terminated callback
		// since they serve two different use cases and current implementation does not
		// support multiple callback registrations
		if ds.AutoReconnectCallback != nil && !ds.disposing {
			ds.AutoReconnectCallback(ds)
		}
	} else {
		ds.connectActionMutex.Unlock()
	}
}

// Dispatcher for connection terminated. This is called from its own separate thread
// in order to cleanly shut down the subscriber in case the connection was terminated
// by the peer. Additionally, this allows the user to automatically reconnect in their
// callback function without having to spawn their own separate thread.
func (ds *DataSubscriber) dispatchConnectionTerminated() {
	ds.connectionTerminationThread = thread.NewThread(func() {
		ds.disconnect(false, true)
	})

	ds.connectionTerminationThread.Start()
}

func (ds *DataSubscriber) dispatchStatusMessage(message string) {
	if ds.StatusMessageCallback != nil {
		go ds.StatusMessageCallback(message)
	}
}

func (ds *DataSubscriber) dispatchErrorMessage(message string) {
	if ds.ErrorMessageCallback != nil {
		go ds.ErrorMessageCallback(message)
	}
}

func (ds *DataSubscriber) runCommandChannelResponseThread() {
	ds.reader = bufio.NewReader(ds.commandChannelSocket)

	for ds.connected {
		ds.readPayloadHeader(io.ReadFull(ds.reader, ds.readBuffer[:payloadHeaderSize]))
	}
}

func (ds *DataSubscriber) readPayloadHeader(bytesTransferred int, err error) {
	if ds.disconnecting {
		return
	}

	if err != nil {
		// Read error, connection may have been closed by peer; terminate connection
		ds.dispatchConnectionTerminated()
		return
	}

	// Gather statistics
	ds.totalCommandChannelBytesReceived += uint64(bytesTransferred)

	packetSize := binary.BigEndian.Uint32(ds.readBuffer)

	if int(packetSize) > cap(ds.readBuffer) {
		ds.readBuffer = ds.readBuffer[:packetSize]
	}

	// Read packet (payload body)
	// This read method is guaranteed not to return until the
	// requested size has been read or an error has occurred.
	ds.readPacket(io.ReadFull(ds.reader, ds.readBuffer[:packetSize]))
}

func (ds *DataSubscriber) readPacket(bytesTransferred int, err error) {
	if ds.disconnecting {
		return
	}

	if err != nil {
		// Read error, connection may have been closed by peer; terminate connection
		ds.dispatchConnectionTerminated()
		return
	}

	// Gather statistics
	ds.totalCommandChannelBytesReceived += uint64(bytesTransferred)

	// Process response
	ds.processServerResponse(ds.readBuffer[:bytesTransferred])
}

// If the user defines a separate UDP channel for their
// subscription, data packets get handled from this thread.
func (ds *DataSubscriber) runDataChannelResponseThread() {
	reader := bufio.NewReader(ds.dataChannelSocket)
	buffer := make([]byte, maxPacketSize)

	for ds.connected {
		length, err := reader.Read(buffer)

		if err != nil {
			ds.dispatchErrorMessage("Error reading data from command channel: " + err.Error())
			break
		}

		// Gather statistics
		ds.totalDataChannelBytesReceived += uint64(length)

		// Process response
		ds.processServerResponse(buffer[:length])
	}
}

func (ds *DataSubscriber) processServerResponse(buffer []byte) {
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
		var message strings.Builder
		message.WriteString("Encountered unexpected server response code: 0x")
		message.WriteString(strconv.FormatInt(int64(commandCode), 16))
		ds.dispatchErrorMessage(message.String())
	}
}

func (ds *DataSubscriber) handleSucceeded(commandCode ServerCommandEnum, data []byte) {
	switch commandCode {
	case ServerCommand.MetadataRefresh:
		ds.handleMetadataRefresh(data)
	case ServerCommand.Subscribe, ServerCommand.Unsubscribe:
		ds.subscribed = commandCode == ServerCommand.Subscribe
		// Fallthrough on these messages because there is
		// still an associated message to be processed.
		fallthrough
	case ServerCommand.RotateCipherKeys, ServerCommand.UpdateProcessingInterval:
		// Each of these responses come with a message that will
		// be delivered to the user via the status message callback.
		var message strings.Builder
		message.WriteString("Received success code in response to server command 0x")
		message.WriteString(strconv.FormatInt(int64(commandCode), 16))

		if data != nil {
			message.Write(data)
		}

		ds.dispatchStatusMessage(message.String())
	default:
		// If we don't know what the message is, we can't interpret
		// the data sent with the packet. Deliver an error message
		// to the user via the error message callback.
		var message strings.Builder
		message.WriteString("Received success code in response to unknown server command 0x")
		message.WriteString(strconv.FormatInt(int64(commandCode), 16))
		ds.dispatchErrorMessage(message.String())
	}
}

func (ds *DataSubscriber) handleFailed(commandCode ServerCommandEnum, data []byte) {
	if data == nil {
		return
	}

	var message strings.Builder

	if commandCode == ServerCommand.Connect {
		ds.connector.connectionRefused = true
	} else {
		message.WriteString("Received failure code in response to server command 0x")
		message.WriteString(strconv.FormatInt(int64(commandCode), 16))
	}

	if data != nil {
		message.Write(data)
	}

	ds.dispatchErrorMessage(message.String())
}

func (ds *DataSubscriber) handleMetadataRefresh(data []byte) {
	if ds.MetadataReceivedCallback != nil {
		if ds.CompressMetadata {
			var err error

			if data, err = decompressGZip(data); err != nil {
				ds.dispatchErrorMessage("Failed to decompress received metadata: " + err.Error())
				return
			}
		}

		go ds.MetadataReceivedCallback(data)
	}
}

func (ds *DataSubscriber) handleDataStartTime(data []byte) {
	if ds.DataStartTimeCallback != nil {
		// Do not use Go routine here, processing sequence may be important.
		// Execute callback directly from socket processing thread:
		ds.DataStartTimeCallback(ticks.Ticks(binary.BigEndian.Uint64(data)))
	}
}

func (ds *DataSubscriber) handleProcessingComplete(data []byte) {
	if ds.ProcessingCompleteCallback != nil {
		go ds.ProcessingCompleteCallback()
	}
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
	signalIndexCache.decode(ds, data, &ds.subscriberID)

	ds.signalIndexCacheLock.Lock()
	ds.signalIndexCache[cacheIndex] = signalIndexCache
	ds.cacheIndex = cacheIndex
	ds.signalIndexCacheLock.Unlock()

	if version > 1 {
		ds.SendServerCommand(ServerCommand.ConfirmSignalIndexCache)
	}

	if ds.SubscriptionUpdatedCallback != nil {
		go ds.SubscriptionUpdatedCallback(signalIndexCache)
	}
}

func (ds *DataSubscriber) handleUpdateBaseTimes(data []byte) {
	if data == nil {
		return
	}

	ds.timeIndex = int32(binary.BigEndian.Uint32(data))

	var baseTimeOffsets [2]int64

	baseTimeOffsets[0] = int64(binary.BigEndian.Uint64(data[4:]))
	baseTimeOffsets[1] = int64(binary.BigEndian.Uint64(data[8:]))

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

	if ds.ConfigurationChangedCallback != nil {
		go ds.ConfigurationChangedCallback()
	}
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

	ds.signalIndexCacheLock.Lock()
	signalIndexCache := ds.signalIndexCache[cacheIndex]
	ds.signalIndexCacheLock.Unlock()

	if compressed {
		ds.parseTSSCMeasurements(signalIndexCache, data[4:], measurements)
	} else {
		ds.parseCompactMeasurements(signalIndexCache, dataPacketFlags, data[4:], measurements)
	}

	if ds.NewMeasurementsCallback != nil {
		// Do not use Go routine here, processing sequence may be important.
		// Execute callback directly from socket processing thread:
		ds.NewMeasurementsCallback(measurements)
	}

	ds.totalMeasurementsReceived += uint64(count)
}

func (ds *DataSubscriber) parseTSSCMeasurements(signalIndexCache *SignalIndexCache, data []byte, measurements []Measurement) {
}

func (ds *DataSubscriber) parseCompactMeasurements(signalIndexCache *SignalIndexCache, dataPacketFlags DataPacketFlagsEnum, data []byte, measurements []Measurement) {
	useMillisecondResolution := ds.subscription.UseMillisecondResolution
	includeTime := ds.subscription.IncludeTime
	index := 0

	for i := 0; i < len(measurements); i++ {
		if signalIndexCache.Count() > 0 {
			// Deserialize compact measurement format
			compactMeasurement := NewCompactMeasurement(signalIndexCache, includeTime, useMillisecondResolution, &ds.baseTimeOffsets)
			index += compactMeasurement.Decode(data[index:])
			measurements[i] = compactMeasurement.Measurement
		} else if ds.lastMissingCacheWarning+missingCacheWarningInterval < ticks.UtcNow() {
			// Warning message for missing signal index cache
			if ds.lastMissingCacheWarning != 0 {
				ds.dispatchStatusMessage("Signal index cache has not arrived. No compact measurements can be parsed.")
			}

			ds.lastMissingCacheWarning = ticks.UtcNow()
		}
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

		ds.signalIndexCacheLock.Lock()
		signalIndexCache := ds.signalIndexCache[signalIndexCacheIndex]
		ds.signalIndexCacheLock.Unlock()

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

			// Remove published measurements from the buffer block queue
			if len(ds.bufferBlockCache) > 0 {
				ds.bufferBlockCache = ds.bufferBlockCache[i:]
			}

			// Publish measurements
			if ds.NewBufferBlocksCallback != nil {
				// Do not use Go routine here, processing sequence may be important.
				// Execute callback directly from socket processing thread:
				ds.NewBufferBlocksCallback(bufferBlockMeasurements)
			}
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

	if ds.NotificationReceivedCallback != nil {
		go ds.NotificationReceivedCallback(message)
	}

	// Send confirmation of receipt of the notification with 4-byte hash
	ds.SendServerCommandWithPayload(ServerCommand.ConfirmNotification, data[:4])
}

// SendServerCommand sends a server command code to the DataPublisher with no payload.
func (ds *DataSubscriber) SendServerCommand(commandCode ServerCommandEnum) {
	ds.SendServerCommandWithPayload(commandCode, nil)
}

// SendServerCommandWithMessage sends a server command code to the DataPublisher along with the specified string message as payload.
func (ds *DataSubscriber) SendServerCommandWithMessage(commandCode ServerCommandEnum, message string) {
	// Latest version of STTP only encodes to UTF8, the default for Go
	if ds.encoding != OperationalEncoding.UTF8 {
		panic("Go implementation of STTP only supports UTF8 string encoding")
	}

	ds.SendServerCommandWithPayload(commandCode, []byte(message))
}

// SendServerCommandWithPayload sends a server command code to the DataPublisher along with the specified data payload.
func (ds *DataSubscriber) SendServerCommandWithPayload(commandCode ServerCommandEnum, data []byte) {
	if !ds.connected {
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

	if _, err := ds.dataChannelSocket.Write(ds.writeBuffer[:commandBufferSize]); err != nil {
		// Write error, connection may have been closed by peer; terminate connection
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
		operationalModes |= OperationalModes.CommpressMetadata
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
	return &ds.connector
}

// GetSignalIndexCache gets the active signal index cache.
func (ds *DataSubscriber) ActiveSignalIndexCache() *SignalIndexCache {
	ds.signalIndexCacheLock.Lock()
	signalIndexCache := ds.signalIndexCache[ds.cacheIndex]
	ds.signalIndexCacheLock.Unlock()

	return signalIndexCache
}

// SubscriberID gets the subscriber ID as assigned by the DataPublisher upon receipt of the SignalIndexCache.
func (ds *DataSubscriber) SubscriberID() guid.Guid {
	return ds.subscriberID
}

// TotalCommandChannelBytesReceived gets the total number of bytes received via the command channel since last connection.
func (ds *DataSubscriber) TotalCommandChannelBytesReceived() uint64 {
	return ds.totalCommandChannelBytesReceived
}

// TotalDataChannelBytesReceived gets the total number of bytes received via the data channel since last connection.
func (ds *DataSubscriber) TotalDataChannelBytesReceived() uint64 {
	if ds.subscription.UdpDataChannel {
		return ds.totalDataChannelBytesReceived
	}

	return ds.totalCommandChannelBytesReceived
}

// TotalMeasurementsReceived gets the total number of measurements received since last subscription.
func (ds *DataSubscriber) TotalMeasurementsReceived() uint64 {
	return ds.totalMeasurementsReceived
}
