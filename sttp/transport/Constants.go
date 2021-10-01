//******************************************************************************************************
//  Constants.go - Gbtc
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
//  09/12/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package transport

import (
	"strconv"
	"strings"
)

const (
	maxPacketSize               = 32768
	payloadHeaderSize           = 4
	responseHeaderSize          = 6
	evenKey                     = 0
	oddKey                      = 1
	keyIndex                    = 0
	ivIndex                     = 1
	missingCacheWarningInterval = 20000000
	defaultLagTime              = 5.0
	defaultLeadTime             = 5.0
	defaultPublishInterval      = 1.0
)

// StateFlagsEnum defines the type for the StateFlags enumeration.
type StateFlagsEnum uint32

// StateFlags is an enumeration of the possible quality states of a Measurement value.
var StateFlags = struct {
	// Normal defines a Measurement flag for a normal state.
	Normal StateFlagsEnum
	// BadData defines a Measurement flag for a bad data state.
	BadData StateFlagsEnum
	// SuspectData defines a Measurement flag for a suspect data state.
	SuspectData StateFlagsEnum
	// OverRangeError defines a Measurement flag for a over range error, i.e., unreasonable high value.
	OverRangeError StateFlagsEnum
	// UnderRangeError defines a Measurement flag for a under range error, i.e., unreasonable low value.
	UnderRangeError StateFlagsEnum
	// AlarmHigh defines a Measurement flag for a alarm for high value.
	AlarmHigh StateFlagsEnum
	// AlarmLow defines a Measurement flag for a alarm for low value.
	AlarmLow StateFlagsEnum
	// WarningHigh defines a Measurement flag for a warning for high value.
	WarningHigh StateFlagsEnum
	// WarningLow defines a Measurement flag for a warning for low value.
	WarningLow StateFlagsEnum
	// FlatlineAlarm defines a Measurement flag for a alarm for flat-lined value, i.e., latched value test alarm.
	FlatlineAlarm StateFlagsEnum
	// ComparisonAlarm defines a Measurement flag for a comparison alarm, i.e., outside threshold of comparison with a real-time value.
	ComparisonAlarm StateFlagsEnum
	// ROCAlarm defines a Measurement flag for a rate-of-change alarm.
	ROCAlarm StateFlagsEnum
	// ReceivedAsBad defines a Measurement flag for a bad value received.
	ReceivedAsBad StateFlagsEnum
	// CalculatedValue defines a Measurement flag for a calculated value state.
	CalculatedValue StateFlagsEnum
	// CalculationError defines a Measurement flag for a calculation error with the value.
	CalculationError StateFlagsEnum
	// CalculationWarning defines a Measurement flag for a calculation warning with the value.
	CalculationWarning StateFlagsEnum
	// ReservedQualityFlag defines a Measurement flag for a reserved quality.
	ReservedQualityFlag StateFlagsEnum
	// BadTime defines a Measurement flag for a bad time state.
	BadTime StateFlagsEnum
	// SuspectTime defines a Measurement flag for a suspect time state.
	SuspectTime StateFlagsEnum
	// LateTimeAlarm defines a Measurement flag for a late time alarm.
	LateTimeAlarm StateFlagsEnum
	// FutureTimeAlarm defines a Measurement flag for a future time alarm.
	FutureTimeAlarm StateFlagsEnum
	// UpSampled defines a Measurement flag for a up-sampled state.
	UpSampled StateFlagsEnum
	// DownSampled defines a Measurement flag for a down-sampled state.
	DownSampled StateFlagsEnum
	// DiscardedValue defines a Measurement flag for a discarded value state.
	DiscardedValue StateFlagsEnum
	// ReservedTimeFlag defines a Measurement flag for a reserved time
	ReservedTimeFlag StateFlagsEnum
	// UserDefinedFlag1 defines a Measurement flag for user defined state 1.
	UserDefinedFlag1 StateFlagsEnum
	// UserDefinedFlag2 defines a Measurement flag for user defined state 2.
	UserDefinedFlag2 StateFlagsEnum
	// UserDefinedFlag3 defines a Measurement flag for user defined state 3.
	UserDefinedFlag3 StateFlagsEnum
	// UserDefinedFlag4 defines a Measurement flag for user defined state 4.
	UserDefinedFlag4 StateFlagsEnum
	// UserDefinedFlag5 defines a Measurement flag for user defined state 5.
	UserDefinedFlag5 StateFlagsEnum
	// SystemError defines a Measurement flag for a system error state.
	SystemError StateFlagsEnum
	// SystemWarning defines a Measurement flag for a system warning state.
	SystemWarning StateFlagsEnum
	// MeasurementError defines a Measurement flag for an error state.
	MeasurementError StateFlagsEnum
}{
	Normal:              0x0,
	BadData:             0x1,
	SuspectData:         0x2,
	OverRangeError:      0x4,
	UnderRangeError:     0x8,
	AlarmHigh:           0x10,
	AlarmLow:            0x20,
	WarningHigh:         0x40,
	WarningLow:          0x80,
	FlatlineAlarm:       0x100,
	ComparisonAlarm:     0x200,
	ROCAlarm:            0x400,
	ReceivedAsBad:       0x800,
	CalculatedValue:     0x1000,
	CalculationError:    0x2000,
	CalculationWarning:  0x4000,
	ReservedQualityFlag: 0x8000,
	BadTime:             0x10000,
	SuspectTime:         0x20000,
	LateTimeAlarm:       0x40000,
	FutureTimeAlarm:     0x80000,
	UpSampled:           0x100000,
	DownSampled:         0x200000,
	DiscardedValue:      0x400000,
	ReservedTimeFlag:    0x800000,
	UserDefinedFlag1:    0x1000000,
	UserDefinedFlag2:    0x2000000,
	UserDefinedFlag3:    0x4000000,
	UserDefinedFlag4:    0x8000000,
	UserDefinedFlag5:    0x10000000,
	SystemError:         0x20000000,
	SystemWarning:       0x40000000,
	MeasurementError:    0x80000000,
}

// String gets the StateFlags enumeration bit values as a string.
func (sfe StateFlagsEnum) String() string {
	if sfe == StateFlags.Normal {
		return "Normal"
	}

	var image strings.Builder

	addFlag := func(flag StateFlagsEnum, name string) {
		if flag&sfe > 0 {
			if image.Len() > 0 {
				image.WriteRune(',')
			}

			image.WriteString(name)
		}
	}

	addFlag(StateFlags.BadData, "BadData")
	addFlag(StateFlags.SuspectData, "SuspectData")
	addFlag(StateFlags.OverRangeError, "OverRangeError")
	addFlag(StateFlags.UnderRangeError, "UnderRangeError")
	addFlag(StateFlags.AlarmHigh, "AlarmHigh")
	addFlag(StateFlags.AlarmLow, "AlarmLow")
	addFlag(StateFlags.WarningHigh, "WarningHigh")
	addFlag(StateFlags.WarningLow, "WarningLow")
	addFlag(StateFlags.FlatlineAlarm, "FlatlineAlarm")
	addFlag(StateFlags.ComparisonAlarm, "ComparisonAlarm")
	addFlag(StateFlags.ROCAlarm, "ROCAlarm")
	addFlag(StateFlags.ReceivedAsBad, "ReceivedAsBad")
	addFlag(StateFlags.CalculatedValue, "CalculatedValue")
	addFlag(StateFlags.CalculationError, "CalculationError")
	addFlag(StateFlags.CalculationWarning, "CalculationWarning")
	addFlag(StateFlags.ReservedQualityFlag, "ReservedQualityFlag")
	addFlag(StateFlags.BadTime, "BadTime")
	addFlag(StateFlags.SuspectTime, "SuspectTime")
	addFlag(StateFlags.LateTimeAlarm, "LateTimeAlarm")
	addFlag(StateFlags.FutureTimeAlarm, "FutureTimeAlarm")
	addFlag(StateFlags.UpSampled, "UpSampled")
	addFlag(StateFlags.DownSampled, "DownSampled")
	addFlag(StateFlags.DiscardedValue, "DiscardedValue")
	addFlag(StateFlags.ReservedTimeFlag, "ReservedTimeFlag")
	addFlag(StateFlags.UserDefinedFlag1, "UserDefinedFlag1")
	addFlag(StateFlags.UserDefinedFlag2, "UserDefinedFlag2")
	addFlag(StateFlags.UserDefinedFlag3, "UserDefinedFlag3")
	addFlag(StateFlags.UserDefinedFlag4, "UserDefinedFlag4")
	addFlag(StateFlags.UserDefinedFlag5, "UserDefinedFlag5")
	addFlag(StateFlags.SystemError, "SystemError")
	addFlag(StateFlags.SystemWarning, "SystemWarning")
	addFlag(StateFlags.MeasurementError, "MeasurementError")

	return image.String()
}

// DataPacketFlagsEnum defines the type for the DataPacketFlags enumeration.
type DataPacketFlagsEnum byte

// DataPacketFlags is an enumeration of the possible flags for a data packet.
var DataPacketFlags = struct {
	// Compact determines if serialized measurement is compact. Bit set = compact, bit clear = full fidelity.
	Compact DataPacketFlagsEnum
	// CipherIndex determines which cipher index to use when encrypting data packet. Bit set = use odd cipher index (i.e., 1), bit clear = use even cipher index (i.e., 0).
	CipherIndex DataPacketFlagsEnum
	// Compressed determines if data packet payload is compressed. Bit set = payload compressed, bit clear = payload normal.
	Compressed DataPacketFlagsEnum
	// CacheIndex determines with signal index cache to use when decoding a data packet. Used by STTP version 2 or greater.
	CacheIndex DataPacketFlagsEnum
	// NoFlags defines state where there are no flags set. This would represent unsynchronized, full fidelity measurement data packets.
	NoFlags DataPacketFlagsEnum
}{
	Compact:     0x02,
	CipherIndex: 0x04,
	Compressed:  0x08,
	CacheIndex:  0x10,
	NoFlags:     0x0,
}

// ServerCommandEnum defines the type for the ServerCommand enumeration.
type ServerCommandEnum byte

// ServerCommand is an enumeration of the possible server commands received
// by a DataPublisher and sent by a DataSubscriber during an STTP session.
var ServerCommand = struct {
	// Connect defines a service command code for handling connect operations. Only used as part of connection refused response.
	Connect ServerCommandEnum
	// MetaDataRefresh defines a service command code for requesting an updated set of metadata.
	MetadataRefresh ServerCommandEnum
	// Subscribe defines a service command code for requesting a subscription of streaming data from server based on connection string that follows.
	Subscribe ServerCommandEnum
	// Unsubscribe  defines a service command code for requesting that server stop sending streaming data to the client and cancel the current subscription.
	Unsubscribe ServerCommandEnum
	// RotateCipherKeys defines a service command code for manually requesting that server send a new set of cipher keys for data packet encryption (UDP only).
	RotateCipherKeys ServerCommandEnum
	// UpdateProcessingInterval defines a service command code for manually requesting that server to update the processing interval with the following specified value.
	UpdateProcessingInterval ServerCommandEnum
	// DefineOperationalModes defines a service command code for establishing operational modes. As soon as connection is established, requests that server set operational modes that affect how the subscriber and publisher will communicate.
	DefineOperationalModes ServerCommandEnum
	// ConfirmNotification defines a service command code for receipt of a notification. This message is sent in response to ServerResponse.Notify.
	ConfirmNotification ServerCommandEnum
	// ConfirmBufferBlock defines a service command code for receipt of a buffer block measurement. This message is sent in response to ServerResponse.BufferBlock.
	ConfirmBufferBlock ServerCommandEnum
	// ConfirmSignalIndexCache defines a service command for confirming the receipt of a signal index cache. This allows publisher to safely transition to next signal index cache.
	ConfirmSignalIndexCache ServerCommandEnum
	// UserCommand00 defines a service command code for handling user-defined commands.
	UserCommand00 ServerCommandEnum
	// UserCommand01 defines a service command code for handling user-defined commands.
	UserCommand01 ServerCommandEnum
	// UserCommand02 defines a service command code for handling user-defined commands.
	UserCommand02 ServerCommandEnum
	// UserCommand03 defines a service command code for handling user-defined commands.
	UserCommand03 ServerCommandEnum
	// UserCommand04 defines a service command code for handling user-defined commands.
	UserCommand04 ServerCommandEnum
	// UserCommand05 defines a service command code for handling user-defined commands.
	UserCommand05 ServerCommandEnum
	// UserCommand06 defines a service command code for handling user-defined commands.
	UserCommand06 ServerCommandEnum
	// UserCommand07 defines a service command code for handling user-defined commands.
	UserCommand07 ServerCommandEnum
	// UserCommand08 defines a service command code for handling user-defined commands.
	UserCommand08 ServerCommandEnum
	// UserCommand09 defines a service command code for handling user-defined commands.
	UserCommand09 ServerCommandEnum
	// UserCommand10 defines a service command code for handling user-defined commands.
	UserCommand10 ServerCommandEnum
	// UserCommand11 defines a service command code for handling user-defined commands.
	UserCommand11 ServerCommandEnum
	// UserCommand12 defines a service command code for handling user-defined commands.
	UserCommand12 ServerCommandEnum
	// UserCommand13 defines a service command code for handling user-defined commands.
	UserCommand13 ServerCommandEnum
	// UserCommand14 defines a service command code for handling user-defined commands.
	UserCommand14 ServerCommandEnum
	// UserCommand15 defines a service command code for handling user-defined commands.
	UserCommand15 ServerCommandEnum
}{
	/*
	   Solicited server commands will receive a ServerResponse.Succeeded or ServerResponse.Failed response
	   code along with an associated success or failure message. Message type for successful responses will
	   be based on server command - for example, server response for a successful MetaDataRefresh command
	   will return a serialized DataSet of the available server metadata. Message type for failed responses
	   will always be a string of text representing the error message.
	*/
	Connect:                  0x00,
	MetadataRefresh:          0x01,
	Subscribe:                0x02,
	Unsubscribe:              0x03,
	RotateCipherKeys:         0x04,
	UpdateProcessingInterval: 0x05,
	DefineOperationalModes:   0x06,
	ConfirmNotification:      0x07,
	ConfirmBufferBlock:       0x08,
	ConfirmSignalIndexCache:  0x0A,
	UserCommand00:            0xD0,
	UserCommand01:            0xD1,
	UserCommand02:            0xD2,
	UserCommand03:            0xD3,
	UserCommand04:            0xD4,
	UserCommand05:            0xD5,
	UserCommand06:            0xD6,
	UserCommand07:            0xD7,
	UserCommand08:            0xD8,
	UserCommand09:            0xD9,
	UserCommand10:            0xDA,
	UserCommand11:            0xDB,
	UserCommand12:            0xDC,
	UserCommand13:            0xDD,
	UserCommand14:            0xDE,
	UserCommand15:            0xDF,
}

// String gets the ServerCommand enumeration value as a string.
func (sce ServerCommandEnum) String() string {
	switch sce {
	case ServerCommand.Connect:
		return "Connect"
	case ServerCommand.MetadataRefresh:
		return "MetadataRefresh"
	case ServerCommand.Subscribe:
		return "Subscribe"
	case ServerCommand.Unsubscribe:
		return "Unsubscribe"
	case ServerCommand.RotateCipherKeys:
		return "RotateCipherKeys"
	case ServerCommand.UpdateProcessingInterval:
		return "UpdateProcessingInterval"
	case ServerCommand.DefineOperationalModes:
		return "DefineOperationalModes"
	case ServerCommand.ConfirmNotification:
		return "ConfirmNotification"
	case ServerCommand.ConfirmBufferBlock:
		return "ConfirmBufferBlock"
	case ServerCommand.ConfirmSignalIndexCache:
		return "ConfirmSignalIndexCache"
	case ServerCommand.UserCommand00:
		return "UserCommand00"
	case ServerCommand.UserCommand01:
		return "UserCommand01"
	case ServerCommand.UserCommand02:
		return "UserCommand02"
	case ServerCommand.UserCommand03:
		return "UserCommand03"
	case ServerCommand.UserCommand04:
		return "UserCommand04"
	case ServerCommand.UserCommand05:
		return "UserCommand05"
	case ServerCommand.UserCommand06:
		return "UserCommand06"
	case ServerCommand.UserCommand07:
		return "UserCommand07"
	case ServerCommand.UserCommand08:
		return "UserCommand08"
	case ServerCommand.UserCommand09:
		return "UserCommand09"
	case ServerCommand.UserCommand10:
		return "UserCommand10"
	case ServerCommand.UserCommand11:
		return "UserCommand11"
	case ServerCommand.UserCommand12:
		return "UserCommand12"
	case ServerCommand.UserCommand13:
		return "UserCommand13"
	case ServerCommand.UserCommand14:
		return "UserCommand14"
	case ServerCommand.UserCommand15:
		return "UserCommand15"
	default:
		return "0x" + strconv.FormatInt(int64(sce), 16)
	}
}

// ServerResponseEnum defines the type for the ServerResponse enumeration.
type ServerResponseEnum byte

// ServerResponse is an enumeration of the possible server responses received sent
// by a DataPublisher and received by a DataSubscriber during an STTP session.
var ServerResponse = struct {
	// Succeeded defines a service response code for indicating a succeeded response. Informs client that its solicited server command succeeded, original command and success message follow.
	Succeeded ServerResponseEnum
	// Failed defines a service response code for indicating a failed response. Informs client that its solicited server command failed, original command and failure message follow.
	Failed ServerResponseEnum
	// DataPacket defines a service response code for indicating a data packet. Unsolicited response informs client that a data packet follows.
	DataPacket ServerResponseEnum
	// UpdateSignalIndexCache defines a service response code for indicating a signal index cache update. Unsolicited response requests that client update its runtime signal index cache with the one that follows.
	UpdateSignalIndexCache ServerResponseEnum
	// UpdateBaseTimes defines a service response code for indicating a runtime base-timestamp offsets have been updated. Unsolicited response requests that client update its runtime base-timestamp offsets with those that follow.
	UpdateBaseTimes ServerResponseEnum
	// UpdateCipherKeys defines a service response code for indicating a runtime cipher keys have been updated. Response, solicited or unsolicited, requests that client update its runtime data cipher keys with those that follow.
	UpdateCipherKeys ServerResponseEnum
	// DataStartTime defines a service response code for indicating the start time of data being published. Unsolicited response provides the start time of data being processed from the first measurement.
	DataStartTime ServerResponseEnum
	// ProcessingComplete defines a service response code for indicating that processing has completed. Unsolicited response provides notification that input processing has completed, typically via temporal constraint.
	ProcessingComplete ServerResponseEnum
	// BufferBlock defines a service response code for indicating a buffer block. Unsolicited response informs client that a raw buffer block follows.
	BufferBlock ServerResponseEnum
	// Notification defines a service response code for indicating a notification. Unsolicited response provides a notification message to the client.
	Notification ServerResponseEnum
	// ConfigurationChanged defines a service response code for indicating a that the publisher configuration metadata has changed. Unsolicited response provides a notification that the publisher's source configuration has changed and that client may want to request a meta-data refresh.
	ConfigurationChanged ServerResponseEnum
	// UserResponse00 defines a service response code for handling user-defined responses.
	UserResponse00 ServerResponseEnum
	// UserResponse01 defines a service response code for handling user-defined responses.
	UserResponse01 ServerResponseEnum
	// UserResponse02 defines a service response code for handling user-defined responses.
	UserResponse02 ServerResponseEnum
	// UserResponse03 defines a service response code for handling user-defined responses.
	UserResponse03 ServerResponseEnum
	// UserResponse04 defines a service response code for handling user-defined responses.
	UserResponse04 ServerResponseEnum
	// UserResponse05 defines a service response code for handling user-defined responses.
	UserResponse05 ServerResponseEnum
	// UserResponse06 defines a service response code for handling user-defined responses.
	UserResponse06 ServerResponseEnum
	// UserResponse07 defines a service response code for handling user-defined responses.
	UserResponse07 ServerResponseEnum
	// UserResponse08 defines a service response code for handling user-defined responses.
	UserResponse08 ServerResponseEnum
	// UserResponse09 defines a service response code for handling user-defined responses.
	UserResponse09 ServerResponseEnum
	// UserResponse10 defines a service response code for handling user-defined responses.
	UserResponse10 ServerResponseEnum
	// UserResponse11 defines a service response code for handling user-defined responses.
	UserResponse11 ServerResponseEnum
	// UserResponse12 defines a service response code for handling user-defined responses.
	UserResponse12 ServerResponseEnum
	// UserResponse13 defines a service response code for handling user-defined responses.
	UserResponse13 ServerResponseEnum
	// UserResponse14 defines a service response code for handling user-defined responses.
	UserResponse14 ServerResponseEnum
	// UserResponse15 defines a service response code for handling user-defined responses.
	UserResponse15 ServerResponseEnum
	// NoOP defines a service response code for indicating a nil-operation keep-alive ping. The command channel can remain quiet for some time, this command allows a period test of client connectivity.
	NoOP ServerResponseEnum
}{
	/*
		Although the server commands and responses will be on two different paths, the response enumeration values
		are defined as distinct from the command values to make it easier to identify codes from a wire analysis.
	*/
	Succeeded:              0x80,
	Failed:                 0x81,
	DataPacket:             0x82,
	UpdateSignalIndexCache: 0x83,
	UpdateBaseTimes:        0x84,
	UpdateCipherKeys:       0x85,
	DataStartTime:          0x86,
	ProcessingComplete:     0x87,
	BufferBlock:            0x88,
	Notification:           0x89,
	ConfigurationChanged:   0x8A,
	UserResponse00:         0xE0,
	UserResponse01:         0xE1,
	UserResponse02:         0xE2,
	UserResponse03:         0xE3,
	UserResponse04:         0xE4,
	UserResponse05:         0xE5,
	UserResponse06:         0xE6,
	UserResponse07:         0xE7,
	UserResponse08:         0xE8,
	UserResponse09:         0xE9,
	UserResponse10:         0xEA,
	UserResponse11:         0xEB,
	UserResponse12:         0xEC,
	UserResponse13:         0xED,
	UserResponse14:         0xEE,
	UserResponse15:         0xEF,
	NoOP:                   0xFF,
}

// String gets the ServerResponse enumeration value as a string.
func (sre ServerResponseEnum) String() string {
	switch sre {
	case ServerResponse.Succeeded:
		return "Succeeded"
	case ServerResponse.Failed:
		return "Failed"
	case ServerResponse.DataPacket:
		return "DataPacket"
	case ServerResponse.UpdateSignalIndexCache:
		return "UpdateSignalIndexCache"
	case ServerResponse.UpdateBaseTimes:
		return "UpdateBaseTimes"
	case ServerResponse.UpdateCipherKeys:
		return "UpdateCipherKeys"
	case ServerResponse.DataStartTime:
		return "DataStartTime"
	case ServerResponse.ProcessingComplete:
		return "ProcessingComplete"
	case ServerResponse.BufferBlock:
		return "BufferBlock"
	case ServerResponse.Notification:
		return "Notification"
	case ServerResponse.ConfigurationChanged:
		return "ConfigurationChanged"
	case ServerResponse.UserResponse00:
		return "UserResponse00"
	case ServerResponse.UserResponse01:
		return "UserResponse01"
	case ServerResponse.UserResponse02:
		return "UserResponse02"
	case ServerResponse.UserResponse03:
		return "UserResponse03"
	case ServerResponse.UserResponse04:
		return "UserResponse04"
	case ServerResponse.UserResponse05:
		return "UserResponse05"
	case ServerResponse.UserResponse06:
		return "UserResponse06"
	case ServerResponse.UserResponse07:
		return "UserResponse07"
	case ServerResponse.UserResponse08:
		return "UserResponse08"
	case ServerResponse.UserResponse09:
		return "UserResponse09"
	case ServerResponse.UserResponse10:
		return "UserResponse10"
	case ServerResponse.UserResponse11:
		return "UserResponse11"
	case ServerResponse.UserResponse12:
		return "UserResponse12"
	case ServerResponse.UserResponse13:
		return "UserResponse13"
	case ServerResponse.UserResponse14:
		return "UserResponse14"
	case ServerResponse.UserResponse15:
		return "UserResponse15"
	case ServerResponse.NoOP:
		return "NoOP"
	default:
		return "0x" + strconv.FormatInt(int64(sre), 16)
	}
}

/*
   Operational modes are sent from a subscriber to a publisher to request operational behaviors for the
   connection, as a result the operation modes must be sent before any other command. The publisher may
   silently refuse some requests (e.g., compression) based on its configuration. Operational modes only
   apply to fundamental protocol control.
*/

// OperationalModesEnum defines the type for the OperationalModes enumeration.
type OperationalModesEnum uint32

// OperationalModes is an enumeration of the possible modes that affect how DataPublisher and DataSubscriber
// communicate during an STTP session.
var OperationalModes = struct {
	// VersionMask defines a bit mask used to get version number of protocol. Version number is currently set to 2.
	VersionMask OperationalModesEnum
	// CompressionModeMask defines a bit mask used to get mode of compression. GZip and TSSC compression are the only modes currently supported. Remaining bits are reserved for future compression modes.
	CompressionModeMask OperationalModesEnum
	// EncodingMask defines a bit mask used to get character encoding used when exchanging messages between publisher and subscriber.
	EncodingMask OperationalModesEnum
	// ReceiveExternalMetadata defines a bit flag used to determine whether external measurements are exchanged during metadata synchronization. Bit set = external measurements are exchanged, bit clear = no external measurements are exchanged.
	ReceiveExternalMetadata OperationalModesEnum
	// ReceiveInternalMetadata defines a bit flag used to determine whether internal measurements are exchanged during metadata synchronization. Bit set = internal measurements are exchanged, bit clear = no internal measurements are exchanged.
	ReceiveInternalMetadata OperationalModesEnum
	// CompressPayloadData defines a bit flag used to determine whether payload data is compressed when exchanging between publisher and subscriber. Bit set = compress, bit clear = no compression.
	CompressPayloadData OperationalModesEnum
	// CompressSignalIndexCache defines a bit flag used to determine whether the signal index cache is compressed when exchanging between publisher and subscriber. Bit set = compress, bit clear = no compression.
	CompressSignalIndexCache OperationalModesEnum
	// CommpressMetadata defines a bit flag used to determine whether metadata is compressed when exchanging between publisher and subscriber. Bit set = compress, bit clear = no compression.
	CommpressMetadata OperationalModesEnum
	// NoFlags defines state where there are no flags set.
	NoFlags OperationalModesEnum
}{
	VersionMask:              0x0000001F,
	CompressionModeMask:      0x000000E0,
	EncodingMask:             0x00000300,
	ReceiveExternalMetadata:  0x02000000,
	ReceiveInternalMetadata:  0x04000000,
	CompressPayloadData:      0x20000000,
	CompressSignalIndexCache: 0x40000000,
	CommpressMetadata:        0x80000000,
	NoFlags:                  0x00000000,
}

// OperationalEncodingEnum defines the type for the OperationalEncoding enumeration.
type OperationalEncodingEnum uint32

// OperationalEncoding is an enumeration of the possible string encoding options of an STTP session.
var OperationalEncoding = struct {
	// UTF16LE targets little-endian 16-bit Unicode character encoding for strings (deprecated).
	UTF16LE OperationalEncodingEnum
	// UTF16BE targets big-endian 16-bit Unicode character encoding for strings (deprecated).
	UTF16BE OperationalEncodingEnum
	// UTF8 targets 8-bit variable-width Unicode character encoding for strings.
	UTF8 OperationalEncodingEnum
}{
	UTF16LE: 0x00000000,
	UTF16BE: 0x00000100,
	UTF8:    0x00000200,
}

// CompressionModesEnum defines the type for the CompressionModes enumeration.
type CompressionModesEnum uint32

// CompressionModes is an enumeration of the possible compression modes supported by STTP.
var CompressionModes = struct {
	// GZip defines a bit flag used determine if GZip compression will be used to metadata exchange.
	GZip CompressionModesEnum
	// TSSC defines a bit flag used determine if the time-series special compression algorithm will be used for data exchange.
	TSSC CompressionModesEnum
	// None defines state where no compression will be used.
	None CompressionModesEnum
}{
	GZip: 0x00000020,
	TSSC: 0x00000040,
	None: 0x00000000,
}

// SecurityModeEnum defines the type for the SecurityMode enumeration.
type SecurityModeEnum int

// SecurityMode is an enumeration of the possible security modes used by the DataPublisher
// to secure data sent over the command channel in STTP.
var SecurityMode = struct {
	// None defines that data will be sent over the wire unencrypted.
	None SecurityModeEnum
	// TLS defines that data will be sent over wire using Transport Layer Security.
	TLS SecurityModeEnum
}{
	None: 0,
	TLS:  1,
}

// ConnectStatusEnum defines the type for the ConnectStatus enumeration.
type ConnectStatusEnum int

// ConnectStatus is an enumeration of the possible connection status results used
// by the SubscriberConnector.
var ConnectStatus = struct {
	// Success defines that a connection succeeded.
	Success ConnectStatusEnum
	// Failed defines that a connection failed.
	Failed ConnectStatusEnum
	// Canceled defines that a connection was cancelled.
	Canceled ConnectStatusEnum
}{
	Success:  1,
	Failed:   0,
	Canceled: -1,
}
