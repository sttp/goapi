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

const (
	maxPacketSize          uint32  = 32768
	payloadHeaderSize      uint32  = 4
	responseHeaderSize     uint32  = 6
	defaultLagTime         float64 = 5.0
	defaultLeadTime        float64 = 5.0
	defaultPublishInterval float64 = 1.0
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

// DataPacketFlagsEnum defines the type for the DataPacketFlags enumeration.
type DataPacketFlagsEnum byte

// DataPacketFlags is an enumeration of the possible flags for a data packet.
var DataPacketFlags = struct {
	// Synchronized determines if data packet is synchronized. Bit set = synchronized, bit clear = unsynchronized.
	Synchronized DataPacketFlagsEnum
	// Compact determines if serialized measurement is compact. Bit set = compact, bit clear = full fidelity.
	Compact DataPacketFlagsEnum
	// CipherIndex determines which cipher index to use when encrypting data packet. Bit set = use odd cipher index (i.e., 1), bit clear = use even cipher index (i.e., 0).
	CipherIndex DataPacketFlagsEnum
	// Compressed determines if data packet payload is compressed. Bit set = payload compressed, bit clear = payload normal.
	Compressed DataPacketFlagsEnum
	// NoFlags defines state where there are no flags set. This would represent unsynchronized, full fidelity measurement data packets.
	NoFlags DataPacketFlagsEnum
}{
	Synchronized: 0x01,
	Compact:      0x02,
	CipherIndex:  0x04,
	Compressed:   0x08,
	NoFlags:      0x0,
}

// ServerCommandEnum defines the type for the ServerCommand enumeration.
type ServerCommandEnum byte

// ServerCommand is an enumeration of the possible server commands received
// by a DataPublisher and sent by a DataSubscriber during an STTP session.
var ServerCommand = struct {
	/*
	   Solicited server commands will receive a ServerResponse.Succeeded or ServerResponse.Failed response
	   code along with an associated success or failure message. Message type for successful responses will
	   be based on server command - for example, server response for a successful MetaDataRefresh command
	   will return a serialized DataSet of the available server metadata. Message type for failed responses
	   will always be a string of text representing the error message.
	*/

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
	Connect:                  0x00,
	MetadataRefresh:          0x01,
	Subscribe:                0x02,
	Unsubscribe:              0x03,
	RotateCipherKeys:         0x04,
	UpdateProcessingInterval: 0x05,
	DefineOperationalModes:   0x06,
	ConfirmNotification:      0x07,
	ConfirmBufferBlock:       0x08,
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
