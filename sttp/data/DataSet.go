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

package data

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/xml"
)

const (
	// XmlSchemaNamespace defines the schema namespace for the W3C XML Schema Definition Language (XSD)
	// used by STTP metadata tables.
	XmlSchemaNamespace = "http://www.w3.org/2001/XMLSchema"

	// ExtXmlSchemaDataNamespace is used to define extended types for XSD elements, e.g., Guid and expression data types.
	ExtXmlSchemaDataNamespace = "urn:schemas-microsoft-com:xml-msdata"

	// DateTimeFormat defines the format of date/time values in an XSD formatted XML schema.
	DateTimeFormat = "2006-01-02T15:04:05.99-07:00"
)

// DataSet represents an in-memory cache of records that is structured similarly to information
// defined in a database. The data set object consists of a collection of data table objects.
// See https://sttp.github.io/documentation/data-sets/ for more information.
// Note that this implementation uses a case-insensitive map for DataTable name lookups.
// Internally, case-insensitive lookups are accomplished using `strings.ToUpper`.
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
	tableNames := make([]string, 0, len(ds.tables))

	for _, table := range ds.tables {
		tableNames = append(tableNames, table.Name())
	}

	return tableNames
}

// Tables gets the DataTable instances defined in the DataSet.
func (ds *DataSet) Tables() []*DataTable {
	tables := make([]*DataTable, 0, len(ds.tables))

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
// Lookup is case-insensitive.
func (ds *DataSet) RemoveTable(tableName string) bool {
	tableName = strings.ToUpper(tableName)

	if _, ok := ds.tables[tableName]; ok {
		delete(ds.tables, tableName)
		return true
	}

	return false
}

// String get a representation of the DataSet as a string.
func (ds *DataSet) String() string {
	var image strings.Builder

	image.WriteString(ds.Name)
	image.WriteString(" [")
	i := 0

	for _, table := range ds.tables {
		if i > 0 {
			image.WriteString(", ")
		}

		image.WriteString(table.Name())
		i++
	}

	image.WriteRune(']')

	return image.String()
}

// ParseXml loads the DataSet from the XML in the specified buffer.
func (ds *DataSet) ParseXml(data []byte) error {
	var doc xml.XmlDocument

	if err := doc.LoadXml(data); err != nil {
		return err
	}

	return ds.ParseXmlDocument(&doc)
}

// ParseXmlDocument loads the DataSet from an existing XmlDocument.
func (ds *DataSet) ParseXmlDocument(doc *xml.XmlDocument) error {
	root := doc.Root

	// Find schema node
	schema, found := root.Item["schema"]

	if !found {
		return errors.New("failed to parse DataSet XML: Cannot find schema node")
	}

	id, found := schema.Attributes["id"]

	if !found || id != root.Name {
		return errors.New("failed to parse DataSet XML: Cannot find schema node matching \"" + root.Name + "\"")
	}

	// Validate schema namespace
	if schema.Namespace != XmlSchemaNamespace {
		return errors.New("failed to parse DataSet XML: cannot find schema namespace \"" + XmlSchemaNamespace + "\"")
	}

	// Populate DataSet schema
	ds.loadSchema(schema)

	// Populate DataSet records
	ds.loadRecords(&root)

	return nil
}

//gocyclo:ignore
func (ds *DataSet) loadSchema(schema *xml.XmlNode) {
	schemaPrefix := schema.Prefix()

	if len(schemaPrefix) > 0 {
		schemaPrefix += ":"
	}

	// Find choice elements representing schema table definitions
	tableNodes := schema.SelectNodes("element/complexType/choice/element")

	for _, tableNode := range tableNodes {
		tableName, found := tableNode.Attributes["name"]

		if !found || len(tableName) == 0 {
			continue
		}

		dataTable := ds.CreateTable(tableName)

		// Find sequence elements representing schema table field definitions
		fieldNodes := tableNode.SelectNodes("complexType/sequence/element")

		dataTable.InitColumns(len(fieldNodes))

		for _, fieldNode := range fieldNodes {
			fieldName, found := fieldNode.Attributes["name"]

			if !found || len(fieldName) == 0 {
				continue
			}

			typeName, found := fieldNode.Attributes["type"]

			if !found || len(typeName) == 0 {
				continue
			}

			typeName = strings.TrimPrefix(typeName, schemaPrefix)

			// Check for extended data type (allows XSD Guid field definitions)
			extDataType, found := fieldNode.Attributes["DataType"]

			if found && len(extDataType) > 0 {
				// Ignore DataType attributes that do not target desired namespace
				if fieldNode.AttributeNamespaces["DataType"] != ExtXmlSchemaDataNamespace {
					extDataType = ""
				}
			}

			dataType, found := ParseXsdDataType(typeName, extDataType)

			// Columns with unsupported XSD data types are skipped
			if !found {
				continue
			}

			// Check for computed expression
			expression, found := fieldNode.Attributes["Expression"]

			if found && len(expression) > 0 {
				// Ignore Expression attributes that do not target desired namespace
				if fieldNode.AttributeNamespaces["Expression"] != ExtXmlSchemaDataNamespace {
					expression = ""
				}
			}

			dataColumn := dataTable.CreateColumn(fieldName, dataType, expression)
			dataTable.AddColumn(dataColumn)
		}

		ds.AddTable(dataTable)
	}
}

//gocyclo:ignore
func (ds *DataSet) loadRecords(root *xml.XmlNode) {
	// Each root node child that matches a table name represents a record
	for _, table := range ds.Tables() {
		records := root.Items[table.Name()]

		table.InitRows(len(records))

		for _, record := range records {
			dataRow := table.CreateRow()

			// Each child node of a record represents a field value
			for _, field := range record.ChildNodes {
				column := table.ColumnByName(field.Name)

				if column == nil {
					continue
				}

				columnIndex := column.Index()
				value := field.Value()

				switch column.Type() {
				case DataType.String:
					dataRow.SetValue(columnIndex, value)
				case DataType.Boolean:
					dataRow.SetValue(columnIndex, value == "true")
				case DataType.DateTime:
					dt, _ := time.Parse(DateTimeFormat, value)
					dataRow.SetValue(columnIndex, dt)
				case DataType.Single:
					f32, _ := strconv.ParseFloat(value, 32)
					dataRow.SetValue(columnIndex, float32(f32))
				case DataType.Double:
					f64, _ := strconv.ParseFloat(value, 64)
					dataRow.SetValue(columnIndex, f64)
				case DataType.Decimal:
					d, _ := decimal.NewFromString(value)
					dataRow.SetValue(columnIndex, d)
				case DataType.Guid:
					g, _ := guid.Parse(value)
					dataRow.SetValue(columnIndex, g)
				case DataType.Int8:
					i8, _ := strconv.ParseInt(value, 0, 8)
					dataRow.SetValue(columnIndex, int8(i8))
				case DataType.Int16:
					i16, _ := strconv.ParseInt(value, 0, 16)
					dataRow.SetValue(columnIndex, int16(i16))
				case DataType.Int32:
					i32, _ := strconv.ParseInt(value, 0, 32)
					dataRow.SetValue(columnIndex, int32(i32))
				case DataType.Int64:
					i64, _ := strconv.ParseInt(value, 0, 64)
					dataRow.SetValue(columnIndex, i64)
				case DataType.UInt8:
					ui8, _ := strconv.ParseUint(value, 0, 8)
					dataRow.SetValue(columnIndex, uint8(ui8))
				case DataType.UInt16:
					ui16, _ := strconv.ParseUint(value, 0, 16)
					dataRow.SetValue(columnIndex, uint16(ui16))
				case DataType.UInt32:
					ui32, _ := strconv.ParseUint(value, 0, 32)
					dataRow.SetValue(columnIndex, uint32(ui32))
				case DataType.UInt64:
					ui64, _ := strconv.ParseUint(value, 0, 64)
					dataRow.SetValue(columnIndex, ui64)
				}
			}

			table.AddRow(dataRow)
		}
	}
}

// // WriteXML saves the DataSet information as XML into the specified buffer.
// func (ds *DataSet) WriteXml(data *[]byte, dataSetName string) {
// // TODO: Will be needed by DataPublisher
// }

// FromXml creates a new DataSet as read from the XML in the specified buffer.
func FromXml(buffer []byte) *DataSet {
	dataSet := NewDataSet()
	dataSet.ParseXml(buffer)
	return dataSet
}
