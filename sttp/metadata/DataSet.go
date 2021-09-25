//******************************************************************************************************
//  DataSet.go - Gbtc
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
//  09/23/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package metadata

import (
	"strings"
)

const (
	// XmlSchemaNamespace defines the schema namespace for the W3C XML Schema Definition Language (XSD)
	// used by STTP metadata tables.
	XmlSchemaNamespace = "http://www.w3.org/2001/XMLSchema"

	// ExtXmlSchemaDataNamespace is used to define an WSD extended type element as a Guid value.
	ExtXmlSchemaDataNamespace = "urn:schemas-microsoft-com:xml-msdata"
)

// DataSet represents an in-memory cache of records that is structured similarly to information
// defined in a database. The data set object consists of a collection of data table objects.
// See https://sttp.github.io/documentation/data-sets/ for more information.
// Note that this implementation uses a case-insensitive map for DataTable name lookups. Internally
// this is accomplished using ToUpper to keep things simple and efficient, however, this implies that
// case-insensitivity will be effectively restricted to ASCII-based table names.
type DataSet struct {
	tables map[string]*DataTable

	// Name defines the name of the DataSet.
	Name string
}

// NewDataSet creates a new DataSet.
func NewDataSet() *DataSet {
	return &DataSet{
		tables: make(map[string]*DataTable),
		Name:   "DataSet",
	}
}

// AddTable adds the specified table to the DataSet.
func (ds *DataSet) AddTable(table *DataTable) {
	ds.tables[strings.ToUpper(table.Name())] = table
}

// Table gets the DataTable for the specified tableName if the name exists;
// otherwise, nil is returned. Lookup is case-insensitive.
func (ds *DataSet) Table(tableName string) *DataTable {
	if table, ok := ds.tables[strings.ToUpper(tableName)]; ok {
		return table
	}

	return nil
}

// TableNames gets the table names defined in the DataSet.
func (ds *DataSet) TableNames() []string {
	tableNames := make([]string, len(ds.tables))

	for _, table := range ds.tables {
		tableNames = append(tableNames, table.Name())
	}

	return tableNames
}

// Tables gets the DataTable instances defined in the DataSet.
func (ds *DataSet) Tables() []*DataTable {
	tables := make([]*DataTable, len(ds.tables))

	for _, table := range ds.tables {
		tables = append(tables, table)
	}

	return tables
}

// CreateTable creates a new DataTable associated with the DataSet.
// Use AddTable to add the new table to the DataSet.
func (ds *DataSet) CreateTable(name string) *DataTable {
	return newDataTable(ds, name)
}

// TableCount gets the total number of tables defined in the DataSet.
func (ds *DataSet) TableCount() int {
	return len(ds.tables)
}

// RemoveTable removes the specified tableName from the DataSet. Returns
// true if table was removed; otherwise, false if it did not exist.
func (ds *DataSet) RemoveTable(tableName string) bool {
	if _, ok := ds.tables[tableName]; ok {
		delete(ds.tables, tableName)
		return true
	}

	return false
}

// ReadXML loads the DataSet from the XML in the specified buffer.
func (ds *DataSet) ReadXML(data []byte) {

}

// WriteXML saves the DataSet information as XML into the specified buffer.
func (ds *DataSet) WriteXml(data *[]byte, dataSetName string) {
	// TODO: Will be needed by DataPublisher
}

// FromXml creates a new DataSet as read from the XML in the specified buffer.
func FromXml(buffer []byte) *DataSet {
	dataSet := NewDataSet()
	dataSet.ReadXML(buffer)
	return dataSet
}
