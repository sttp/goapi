# STTP Data Sets

A [data set](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataSet.go) represents an in-memory cache of records that is structured similarly to information defined in a database. The data set object consists of a collection of [data table](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataTable.go) objects.

Data tables define of collection of [data columns](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataColumn.go) where each data column defines a name and [data type](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataType.go). Data columns can also be computed where its value would be derived from other columns and [functions](https://sttp.github.io/documentation/filter-expressions/#filter-expression-functions) defined in an expression.

Data tables also define a set of [data rows](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataRow.go) where each data row defines a record of information with a field value for each defined data column. Each field value can be `null` regardless of the defined data column type.

A data set schema and associated records can be read from and written to XML documents. The XML specification used for serialization is the standard for [W3C XML Schema Definition Language (XSD)](https://www.w3.org/TR/xmlschema/). See the [ParseXmlDocument and GenerateXmlDocument](https://github.com/sttp/goapi/blob/main/sttp/metadata/DataSet.go#L140) functions.

> :information_source: STTP requires that schema information be included with serialized XML data sets; the STTP API does not attempt to infer a schema from the data. Schema functionality also includes DataColumn expressions that operate similarly to the .NET [System.Data.DataColumn.Expression](https://docs.microsoft.com/en-us/dotnet/api/system.data.datacolumn.expression) to allow for computed columns; however, STTP defines more [functions](https://sttp.github.io/documentation/filter-expressions/#filter-expression-functions) than the .NET implementation, so a serialized STTP data set that includes column expressions using functions not available to a .NET DataColumn will fail to evaluate when accessed from within .NET.

> :small_blue_diamond: The STTP DataSet implementation is always case-insensitive for table and column name lookups as the primary use-case for STTP data sets is for use with [filter expressions](https://sttp.github.io/documentation/filter-expressions/). The code uses [`ToUpper`](https://pkg.go.dev/strings#ToUpper) for its case-insensitive lookups.
