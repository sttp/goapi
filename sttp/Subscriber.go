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
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sttp/goapi/sttp/format"
	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/metadata"
	"github.com/sttp/goapi/sttp/ticks"
	"github.com/sttp/goapi/sttp/transport"
)

// Subscriber represents an STTP data subscriber.
type Subscriber struct {
	// Configuration reference
	config *Config

	// DataSubscriber reference
	ds *transport.DataSubscriber

	// Callback references
	statusMessageLogger            func(message string)
	errorMessageLogger             func(message string)
	metadataReceiver               func(dataSet *metadata.DataSet)
	dataStartTimeReceiver          func(startTime time.Time)
	configurationChangedReceiver   func()
	historicalReadCompleteReceiver func()
	connectionEstablishedReceiver  func()
	connectionTerminatedReceiver   func()

	// Lock used to synchronize console writes
	consoleLock sync.Mutex
}

// NewSubscriber creates a new Subscriber.
func NewSubscriber() *Subscriber {
	sb := Subscriber{
		config: NewConfig(),
		ds:     transport.NewDataSubscriber(),
	}
	sb.statusMessageLogger = sb.DefaultStatusMessageLogger
	sb.errorMessageLogger = sb.DefaultErrorMessageLogger
	sb.connectionEstablishedReceiver = sb.DefaultConnectionEstablishedReceiver
	sb.connectionTerminatedReceiver = sb.DefaultConnectionTerminatedReceiver
	return &sb
}

// Close cleanly shuts down a DataSubscriber that is no longer being used, e.g.,
// during a normal application exit.
func (sb *Subscriber) Close() {
	if sb.ds != nil {
		sb.ds.Dispose()
	}
}

// dataSubscriber gets a reference to the internal DataSubscriber instance.
func (sb *Subscriber) dataSubscriber() *transport.DataSubscriber {
	if sb.ds == nil {
		panic("Internal DataSubscriber instance has not been initialized. Make sure to use NewSubscriber.")
	}

	return sb.ds
}

// IsConnected determines if Subscriber is currently connected to a data publisher.
func (sb *Subscriber) IsConnected() bool {
	return sb.dataSubscriber().IsConnected()
}

// IsSubscribed determines if Subscriber is currently subscribed to a data stream.
func (sb *Subscriber) IsSubscribed() bool {
	return sb.dataSubscriber().IsSubscribed()
}

// ActiveSignalIndexCache gets the active signal index cache.
func (sb *Subscriber) ActiveSignalIndexCache() *transport.SignalIndexCache {
	return sb.dataSubscriber().ActiveSignalIndexCache()
}

// SubscriberID gets the subscriber ID as assigned by the data publisher upon receipt of the SignalIndexCache.
func (sb *Subscriber) SubscriberID() guid.Guid {
	return sb.dataSubscriber().SubscriberID()
}

// TotalCommandChannelBytesReceived gets the total number of bytes received via the command channel since last connection.
func (sb *Subscriber) TotalCommandChannelBytesReceived() uint64 {
	return sb.dataSubscriber().TotalCommandChannelBytesReceived()
}

// TotalDataChannelBytesReceived gets the total number of bytes received via the data channel since last connection.
func (sb *Subscriber) TotalDataChannelBytesReceived() uint64 {
	return sb.dataSubscriber().TotalDataChannelBytesReceived()
}

// TotalMeasurementsReceived gets the total number of measurements received since last subscription.
func (sb *Subscriber) TotalMeasurementsReceived() uint64 {
	return sb.dataSubscriber().TotalMeasurementsReceived()
}

// LookupMetadata gets the MeasurementMetadata for the specified signalID from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (sb *Subscriber) LookupMetadata(signalID guid.Guid) *transport.MeasurementMetadata {
	return sb.dataSubscriber().LookupMetadata(signalID)
}

// Metadata gets the measurement-level metadata associated with a measurement from the local
// registry. If the metadata does not exist, a new record is created and returned.
func (sb *Subscriber) Metadata(measurement *transport.Measurement) *transport.MeasurementMetadata {
	return sb.dataSubscriber().Metadata(measurement)
}

// AdjustedValue gets the Value of a Measurement with any linear adjustments applied from the
// measurement's Adder and Multiplier metadata, if found.
func (sb *Subscriber) AdjustedValue(measurement *transport.Measurement) float64 {
	return sb.dataSubscriber().AdjustedValue(measurement)
}

// Dial starts the client-based connection cycle to an STTP publisher. Config parameter controls
// connection related settings, set value to nil for default values. When the config defines
// AutoReconnect as true, the connection will automatically be retried when the connection drops.
// If the config defines AutoRequestMetadata as true, then upon successful connection, meta-data
// will be requested. When the config defines both AutoRequestMetadata and AutoSubscribe as true,
// subscription will occur after reception of metadata. When the config defines AutoRequestMetadata
// as false and AutoSubscribe as true, subscription will occur at successful connection.
func (sb *Subscriber) Dial(address string, config *Config) error {
	hostname, portname, err := net.SplitHostPort(address)

	if err != nil {
		return err
	}

	port, err := strconv.Atoi(portname)

	if err != nil {
		return fmt.Errorf("invalid port number \"%s\": %s", portname, err.Error())
	}

	if port < 1 || port > math.MaxUint16 {
		return fmt.Errorf("port number \"%s\" is out of range: must be 1 to %d", portname, math.MaxUint16)
	}

	if config != nil {
		sb.config = config
	}

	sb.connect(hostname, uint16(port))
	return nil
}

func (sb *Subscriber) connect(hostname string, port uint16) {
	if sb.config == nil {
		panic("Internal Config instance has not been initialized. Make sure to use NewSubscriber.")
	}

	ds := sb.dataSubscriber()
	con := ds.Connector()

	// Set connection properties
	con.Hostname = hostname
	con.Port = port

	con.MaxRetries = sb.config.MaxRetries
	con.RetryInterval = sb.config.RetryInterval
	con.MaxRetryInterval = sb.config.MaxRetryInterval
	con.AutoReconnect = sb.config.AutoReconnect

	ds.CompressPayloadData = sb.config.CompressPayloadData
	ds.CompressMetadata = sb.config.CompressMetadata
	ds.CompressSignalIndexCache = sb.config.CompressSignalIndexCache
	ds.Version = sb.config.Version

	// Register direct Subscriber callbacks
	con.ErrorMessageCallback = sb.errorMessageLogger
	ds.StatusMessageCallback = sb.statusMessageLogger
	ds.ErrorMessageCallback = sb.errorMessageLogger
	ds.ConnectionTerminatedCallback = sb.connectionTerminatedReceiver

	// Register callbacks with intermediate handlers
	con.ReconnectCallback = sb.handleReconnect
	ds.MetadataReceivedCallback = sb.handleMetadataReceived
	ds.DataStartTimeCallback = sb.handleDataStartTime
	ds.ConfigurationChangedCallback = sb.handleConfigurationChanged
	ds.ProcessingCompleteCallback = sb.handleProcessingComplete

	var status transport.ConnectStatusEnum

	// Connect and subscribe to publisher
	if status = con.Connect(ds); status == transport.ConnectStatus.Success {
		if sb.connectionEstablishedReceiver != nil {
			sb.connectionEstablishedReceiver()
		}

		// If automatically parsing metadata, request metadata upon successful connection,
		// after metadata is received the SubscriberInstance will then initiate subscribe;
		// otherwise, subscribe is initiated immediately (when auto subscribe requested)
		if sb.config.AutoRequestMetadata {
			sb.RequestMetadata()
		} else if sb.config.AutoSubscribe {
			ds.Subscribe()
		}
	} else if status == transport.ConnectStatus.Failed {
		sb.ErrorMessage("All connection attempts failed")
	}
}

// Disconnect disconnects from an STTP publisher.
func (sb *Subscriber) Disconnect() {
	sb.dataSubscriber().Disconnect()
}

// RequestMetadata sends a request to the data publisher indicating that the Subscriber would
// like new metadata. Any defined MetadataFilters will be included in request.
func (sb *Subscriber) RequestMetadata() {
	ds := sb.dataSubscriber()

	if len(sb.config.MetadataFilters) == 0 {
		ds.SendServerCommand(transport.ServerCommand.MetadataRefresh)
		return
	}

	filters := ds.EncodeString(sb.config.MetadataFilters)
	buffer := make([]byte, 4+len(filters))

	binary.BigEndian.PutUint32(buffer, uint32(len(filters)))
	copy(buffer[4:], filters)

	ds.SendServerCommandWithPayload(transport.ServerCommand.MetadataRefresh, buffer)
}

// Subscribe sets up a request indicating that the Subscriber would like to start receiving
// streaming data from a data publisher. If the subscriber is already connected, the updated
// filter expression and subscription settings will be requested immediately; otherwise, the
// settings will be used when the connection to the data publisher is established.
//
// The filterExpression defines the desired measurements for a subscription. Examples include:
//
// * Directly specified signal IDs (UUID values in string format):
//     38A47B0-F10B-4143-9A0A-0DBC4FFEF1E8; E4BBFE6A-35BD-4E5B-92C9-11FF913E7877
//
// * Directly specified tag names:
//     DOM_GPLAINS-BUS1:VH; TVA_SHELBY-BUS1:VH
//
// * Directly specified identifiers in "measurement key" format:
//     PPA:15; STAT:20
//
// * A filter expression against a selection view:
//     FILTER ActiveMeasurements WHERE Company='GPA' AND SignalType='FREQ'
//
// Settings parameter controls subscription related settings, set value to nil for default values.
func (sb *Subscriber) Subscribe(filterExpression string, settings *Settings) {
	ds := sb.dataSubscriber()
	sub := ds.Subscription()

	if settings == nil {
		settings = &settingsDefaults
	}

	sub.FilterExpression = filterExpression
	sub.Throttled = settings.Throttled
	sub.PublishInterval = settings.PublishInterval

	if settings.UdpPort > 0 {
		sub.UdpDataChannel = true
		sub.DataChannelLocalPort = settings.UdpPort
	} else {
		sub.UdpDataChannel = false
		sub.DataChannelLocalPort = 0
	}

	sub.IncludeTime = settings.IncludeTime
	sub.UseMillisecondResolution = settings.UseMillisecondResolution
	sub.RequestNaNValueFilter = settings.RequestNaNValueFilter
	sub.StartTime = settings.StartTime
	sub.StopTime = settings.StopTime
	sub.ConstraintParameters = settings.ConstraintParameters
	sub.ProcessingInterval = settings.ProcessingInterval
	sub.ExtraConnectionStringParameters = settings.ExtraConnectionStringParameters

	if ds.IsConnected() {
		ds.Subscribe()
	}
}

// Unsubscribe sends a request to the data publisher indicating that the Subscriber would
// like to stop receiving streaming data.
func (sb *Subscriber) Unsubscribe() {
	sb.dataSubscriber().Unsubscribe()
}

// ReadMeasurements sets up a new MeasurementReader to start reading measurements.
func (sb *Subscriber) ReadMeasurements() *MeasurementReader {
	return newMeasurementReader(sb)
}

// Local callback handlers:

// StatusMessage executes the defined status message logger callback.
func (sb *Subscriber) StatusMessage(message string) {
	if sb.statusMessageLogger == nil {
		return
	}

	sb.statusMessageLogger(message)
}

// ErrorMessage executes the defined error message logger callback.
func (sb *Subscriber) ErrorMessage(message string) {
	if sb.errorMessageLogger == nil {
		return
	}

	sb.errorMessageLogger(message)
}

// Intermediate callback handlers:

func (sb *Subscriber) handleReconnect(ds *transport.DataSubscriber) {
	if ds.IsConnected() {
		if sb.connectionEstablishedReceiver != nil {
			sb.connectionEstablishedReceiver()
		}

		// If automatically parsing metadata, request metadata upon successful connection,
		// after metadata is received the SubscriberInstance will then initiate subscribe;
		// otherwise, subscribe is initiated immediately (when auto subscribe requested)
		if sb.config.AutoRequestMetadata {
			sb.RequestMetadata()
		} else if sb.config.AutoSubscribe {
			ds.Subscribe()
		}
	} else {
		ds.Disconnect()
		sb.StatusMessage("Connection retry attempts exceeded.")
	}
}

func (sb *Subscriber) handleMetadataReceived(data []byte) {
	parseStarted := time.Now()
	dataSet := metadata.NewDataSet()
	err := dataSet.ParseXml(data)

	if err == nil {
		sb.loadMeasurementMetadata(dataSet)
	} else {
		sb.ErrorMessage("Failed to parse received XML metadata: " + err.Error())
	}

	sb.showMetadataSummary(dataSet, parseStarted)

	if sb.metadataReceiver != nil {
		sb.metadataReceiver(dataSet)
	}

	if sb.config.AutoRequestMetadata && sb.config.AutoSubscribe {
		sb.dataSubscriber().Subscribe()
	}
}

func (sb *Subscriber) loadMeasurementMetadata(dataSet *metadata.DataSet) {
	measurements := dataSet.Table("MeasurementDetail")

	if measurements != nil {
		signalIDIndex := measurements.ColumnIndex("SignalID")

		if signalIDIndex > -1 {
			idIndex := measurements.ColumnIndex("ID")
			pointTagIndex := measurements.ColumnIndex("PointTag")
			signalRefIndex := measurements.ColumnIndex("SignalReference")
			signalTypeIndex := measurements.ColumnIndex("SignalAcronym")
			descriptionIndex := measurements.ColumnIndex("Description")
			updatedOnIndex := measurements.ColumnIndex("UpdatedOn")
			ds := sb.dataSubscriber()

			for i := 0; i < measurements.RowCount(); i++ {
				measurement := measurements.Row(i)

				if measurement == nil {
					continue
				}

				signalID, null, err := measurement.GuidValue(signalIDIndex)

				if null || err != nil {
					continue
				}

				metadata := ds.LookupMetadata(signalID)

				if idIndex > -1 {
					id, _, _ := measurement.StringValue(idIndex)
					parts := strings.Split(id, ":")

					if len(parts) == 2 {
						metadata.Source = parts[0]
						metadata.ID, _ = strconv.ParseUint(parts[1], 10, 64)
					}
				}

				if pointTagIndex > -1 {
					metadata.Tag, _, _ = measurement.StringValue(pointTagIndex)
				}

				if signalRefIndex > -1 {
					metadata.SignalReference, _, _ = measurement.StringValue(signalRefIndex)
				}

				if signalTypeIndex > -1 {
					metadata.SignalType, _, _ = measurement.StringValue(signalTypeIndex)
				}

				if descriptionIndex > -1 {
					metadata.Description, _, _ = measurement.StringValue(descriptionIndex)
				}

				if updatedOnIndex > -1 {
					metadata.UpdatedOn, _, _ = measurement.DateTimeValue(updatedOnIndex)
				}
			}
		} else {
			sb.ErrorMessage("Received metadata does not contain the required MeasurementDetail.SignalID field")
		}
	} else {
		sb.ErrorMessage("Received metadata does not contain the required MeasurementDetail table")
	}
}

func (sb *Subscriber) showMetadataSummary(dataSet *metadata.DataSet, parseStarted time.Time) {
	getRowCount := func(tableName string) int {
		table := dataSet.Table(tableName)

		if table == nil {
			return 0
		}

		return table.RowCount()
	}

	var tableDetails strings.Builder
	totalRows := 0

	tableDetails.WriteString("    Discovered:\n")

	for _, table := range dataSet.Tables() {
		tableName := table.Name()
		tableRows := getRowCount(tableName)
		totalRows += tableRows
		tableDetails.WriteString(fmt.Sprintf("        %s %s records\n", format.Int(tableRows), tableName))
	}

	var message strings.Builder

	message.WriteString("Parsed ")
	message.WriteString(format.Int(totalRows))
	message.WriteString(" metadata records in ")
	message.WriteString(format.Float(time.Since(parseStarted).Seconds(), 3))
	message.WriteString(" seconds.\n")
	message.WriteString(tableDetails.String())

	schemaVersion := dataSet.Table("SchemaVersion")

	if schemaVersion != nil {
		message.WriteString("Metadata schema version: " + schemaVersion.RowValueAsStringByName(0, "VersionNumber"))
	} else {
		message.WriteString("No SchemaVersion table found in metadata")
	}

	sb.StatusMessage(message.String())
}

func (sb *Subscriber) handleDataStartTime(startTime ticks.Ticks) {
	if sb.dataStartTimeReceiver == nil {
		return
	}

	sb.dataStartTimeReceiver(ticks.ToTime(startTime))
}

func (sb *Subscriber) handleConfigurationChanged() {
	if sb.configurationChangedReceiver != nil {
		sb.configurationChangedReceiver()
	}

	if sb.config.AutoRequestMetadata {
		sb.RequestMetadata()
	}
}

func (sb *Subscriber) handleProcessingComplete(message string) {
	sb.StatusMessage(message)

	if sb.historicalReadCompleteReceiver != nil {
		sb.historicalReadCompleteReceiver()
	}
}

// DefaultStatusMessageLogger implements the default handler for the statusMessage callback.
// Default implementation synchronously writes output to stdio. Logging is recommended.
func (sb *Subscriber) DefaultStatusMessageLogger(message string) {
	sb.consoleLock.Lock()
	defer sb.consoleLock.Unlock()
	fmt.Println(message)
}

// DefaultErrorMessageLogger implements the default handler for the errorMessage callback.
// Default implementation synchronously writes output to to stderr. Logging is recommended.
func (sb *Subscriber) DefaultErrorMessageLogger(message string) {
	sb.consoleLock.Lock()
	defer sb.consoleLock.Unlock()
	fmt.Fprintln(os.Stderr, message)
}

// DefaultConnectionEstablishedReceiver implements the default handler for the ConnectionEstablished callback.
// Default implementation simply writes connection feedback to statusMessage callback.
func (sb *Subscriber) DefaultConnectionEstablishedReceiver() {
	con := sb.dataSubscriber().Connector()
	sb.StatusMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " established.")
}

// DefaultConnectionTerminatedReceiver implements the default handler for the ConnectionTerminated callback.
// Default implementation simply writes connection terminated feedback to errorMessage callback.
func (sb *Subscriber) DefaultConnectionTerminatedReceiver() {
	con := sb.dataSubscriber().Connector()
	sb.ErrorMessage("Connection to " + con.Hostname + ":" + strconv.Itoa(int(con.Port)) + " terminated.")
}

// SetStatusMessageLogger defines the callback that handles informational message logging.
func (sb *Subscriber) SetStatusMessageLogger(callback func(message string)) {
	sb.statusMessageLogger = callback
}

// SetErrorMessageLogger defines the callback that handles error message logging.
func (sb *Subscriber) SetErrorMessageLogger(callback func(message string)) {
	sb.errorMessageLogger = callback
}

// SetMetadataReceiver defines the callback that handles reception of the metadata response.
func (sb *Subscriber) SetMetadataReceiver(callback func(dataSet *metadata.DataSet)) {
	sb.metadataReceiver = callback
}

// SetSubscriptionUpdatedReceiver defines the callback that handles notifications that a new
// SignalIndexCache has been received.
func (sb *Subscriber) SetSubscriptionUpdatedReceiver(callback func(signalIndexCache *transport.SignalIndexCache)) {
	sb.dataSubscriber().SubscriptionUpdatedCallback = callback
}

// SetDataStartTimeReceiver defines the callback that handles notification of first received measurement.
func (sb *Subscriber) SetDataStartTimeReceiver(callback func(startTime time.Time)) {
	sb.dataStartTimeReceiver = callback
}

// SetConfigurationChangedReceiver defines the callback that handles notifications that the data publisher
// configuration has changed.
func (sb *Subscriber) SetConfigurationChangedReceiver(callback func()) {
	sb.configurationChangedReceiver = callback
}

// SetNewMeasurementsReceiver defines the callback that handles reception of new measurements.
func (sb *Subscriber) SetNewMeasurementsReceiver(callback func(measurements []transport.Measurement)) {
	sb.dataSubscriber().NewMeasurementsCallback = callback
}

// SetNewBufferBlocksReceiver defines the callback that handles reception of new buffer blocks.
func (sb *Subscriber) SetNewBufferBlocksReceiver(callback func(bufferBlocks []transport.BufferBlock)) {
	sb.dataSubscriber().NewBufferBlocksCallback = callback
}

// SetNotificationReceiver defines the callback that handles reception of a notification.
func (sb *Subscriber) SetNotificationReceiver(callback func(notification string)) {
	sb.dataSubscriber().NotificationReceivedCallback = callback
}

// SetHistoricalReadCompleteReceiver defines the callback that handles notification that temporal processing
// has completed, i.e., the end of a historical playback data stream has been reached.
func (sb *Subscriber) SetHistoricalReadCompleteReceiver(callback func()) {
	sb.historicalReadCompleteReceiver = callback
}

// SetConnectionEstablishedReceiver defines the callback that handles notification that a connection has been established.
// Default implementation simply writes connection feedback to StatusMessage handler.
func (sb *Subscriber) SetConnectionEstablishedReceiver(callback func()) {
	sb.connectionEstablishedReceiver = callback
}

// SetConnectionTerminatedReceiver defines the callback that handles notification that a connection has been terminated.
// Default implementation simply writes connection terminated feedback to ErrorMessage handler.
func (sb *Subscriber) SetConnectionTerminatedReceiver(callback func()) {
	sb.connectionTerminatedReceiver = callback
}
