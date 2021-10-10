//******************************************************************************************************
//  Config.go - Gbtc
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
//  09/29/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package sttp

// Config defines the STTP connection parameters.
type Config struct {
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
	// received metadata. Each filter expression should be separated by semi-colon.
	MetadataFilters string

	// Version defines the target STTP protocol version. This currently defaults to 2.
	Version byte

	// RfcGuidEncoding determines if Guid wire serialization should use RFC encoding.
	// This defaults to true.
	RfcGuidEncoding bool
}

// configDefaults define the default values for an STTP connection Config.
var configDefaults = Config{
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
	RfcGuidEncoding:          true,
}

// NewConfig creates a new Config instance initialzed with default values.
func NewConfig() *Config {
	config := configDefaults
	return &config
}
