//******************************************************************************************************
//  TableIDFields.go - Gbtc
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
//  10/07/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************
package data

// TableIDFields represents the primary identification field names for a metadata
// table that is being used as the source for an STTP filter expression. See:
// https://sttp.github.io/documentation/filter-expressions/#activemeasurements
type TableIDFields struct {
	// SignalIDFieldName defines the field name of the signal ID field, type Guid.
	// Common value is "SignalID".
	SignalIDFieldName string
	// MeasurementKeyFieldName defines the name of the measurement key field
	// (format like "instance:id"), type string. Common value is "ID".
	MeasurementKeyFieldName string
	// PointTagFieldName defines the name of the point tag field, type string.
	// Common value is "PointTag".
	PointTagFieldName string
}

// DefaultTableIDFields defines the common default table ID field names.
var DefaultTableIDFields = &TableIDFields{
	SignalIDFieldName:       "SignalID",
	MeasurementKeyFieldName: "ID",
	PointTagFieldName:       "PointTag",
}
