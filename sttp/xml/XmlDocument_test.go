//******************************************************************************************************
//  XmlDocument_test.go - Gbtc
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
//  09/25/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package xml

import (
	"fmt"
	"testing"
)

var doc XmlDocument

const (
	XmlSchemaNamespace        = "http://www.w3.org/2001/XMLSchema"
	ExtXmlSchemaDataNamespace = "urn:schemas-microsoft-com:xml-msdata"
)

func init() {
	doc.LoadXmlFromFile("../../test/SampleMetadata.xml")
}

func TestRootLevel(t *testing.T) {
	if doc.Root.Level != 0 {
		t.Fatalf("Root level in document tree should be zero")
	}
}

func TestValue(t *testing.T) {
	schemaVersion := doc.Root.Item["SchemaVersion"]
	versionNumber := schemaVersion.Item["VersionNumber"].Value()

	if versionNumber != "9" {
		t.Fatalf("SampleMetadata.xml expected to have schema version of 9, received: %s", versionNumber)
	}
}

func TestPath(t *testing.T) {
	schemaVersion := doc.Root.Item["SchemaVersion"]
	versionNumberPath := schemaVersion.Item["VersionNumber"].Path()

	if versionNumberPath != "//DataSet/SchemaVersion/VersionNumber" {
		t.Fatalf("SampleMetadata.xml expected to have schema version number path of \"//DataSet/SchemaVersion/VersionNumber\", received: %s", versionNumberPath)
	}
}

func TestPrefix(t *testing.T) {
	schema := doc.Root.Item["schema"]
	prefix := schema.Prefix()

	if prefix != "xs" {
		t.Fatalf("SampleMetadata.xml expected to have schema prefix of \"xs\", received: \"%s\"", prefix)
	}
}

func TestChildNodeLoad(t *testing.T) {
	root := doc.Root

	if !root.HasChildNodes() {
		t.Fatalf("SampleMetadata.xml expected to have child nodes")
	}

	if len(root.ChildNodes) != 138 {
		t.Fatalf("SampleMetadata.xml expected to have 138 root child nodes")
	}
}

func TestMaxDepthLoad(t *testing.T) {
	if doc.MaxDepth() != 9 {
		t.Fatalf("SampleMetadata.xml expected to have max depth of 9")
	}
}

func TestNamespaceLoad(t *testing.T) {
	schema := doc.Root.Item["schema"]

	if schema.Namespace != XmlSchemaNamespace {
		t.Fatalf("SampleMetadata.xml expected to have schema namespace of \"%s\", received: \"%s\"", XmlSchemaNamespace, schema.Namespace)
	}
}

func TestAttributesLoad(t *testing.T) {
	schema := doc.Root.Item["schema"]

	id, found := schema.Attributes["id"]

	if !found {
		t.Fatalf("SampleMetadata.xml \"schema\" element expected to have attribute \"id\" = \"DataSet\", found none")
	}

	if id != "DataSet" {
		t.Fatalf("SampleMetadata.xml \"schema\" element expected to have attribute \"id\" = \"DataSet\", received: \"%s\"", id)
	}

	if len(schema.Attributes) != 3 {
		t.Fatalf("SampleMetadata.xml \"schema\" element expected to have 3 attributes, received: %d", len(schema.Attributes))
	}
}

func TestAttributeNamespacesLoad(t *testing.T) {
	schema := doc.Root.Item["schema"]
	tableNodes := schema.SelectNodes("element/complexType/choice/element")
	guidFieldNodes := SelectNodes(tableNodes, "complexType/sequence/element[@DataType='System.Guid']")

	if len(guidFieldNodes) != 3 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 3 fields with Guid type, received: %d", len(guidFieldNodes))
	}

	for _, node := range guidFieldNodes {
		if namespace, found := node.AttributeNamespaces["DataType"]; found {
			if namespace != ExtXmlSchemaDataNamespace {
				t.Fatalf("SampleMetadata.xml Guid type fields expected to have namespace of \"%s\", received: \"%s\"", ExtXmlSchemaDataNamespace, namespace)
			}
		} else {
			t.Fatalf("SampleMetadata.xml failed to find attribute namespace for Guid type in field node: \"%s\"", node.Name)
		}
	}
}

func TestSelectNodes(t *testing.T) {
	nodes := doc.SelectNodes("schema/element/complexType/choice/element")

	if len(nodes) != 4 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 4 table definitions, received: %d", len(nodes))
	}
}

func TestSelectNodesFromRoot(t *testing.T) {
	tableNodes := doc.SelectNodes("//DataSet/schema/element/complexType/choice/element")

	if len(tableNodes) != 4 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 4 table definitions, received: %d", len(tableNodes))
	}

	records := doc.SelectNodes("SchemaVersion[VersionNumber]")

	if len(records) != 1 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 1 SchemaVersion record, received: %d", len(records))
	}
}

func TestSelectNodesWithWildcards(t *testing.T) {
	schema := doc.Root.Item["schema"]
	tableNodes := schema.SelectNodes("element/complexType/choice/*")

	if len(tableNodes) != 4 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 4 table definitions, received: %d", len(tableNodes))
	}

	guidFieldNodes := SelectNodes(tableNodes, "complexType/sequence/*[@DataType]")

	if len(guidFieldNodes) != 3 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 3 fields with Guid type, received: %d", len(guidFieldNodes))
	}

	guidFieldNodes = SelectNodes(tableNodes, "complexType/sequence/element[@*='System.Guid']")

	if len(guidFieldNodes) != 3 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 3 fields with Guid type, received: %d", len(guidFieldNodes))
	}

	records := doc.SelectNodes("SchemaVersion[*]")

	if len(records) != 1 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 1 SchemaVersion record, received: %d", len(records))
	}
}

func TestSelectNodesWithSubExprPredicateAttributes(t *testing.T) {
	schema := doc.Root.Item["schema"]
	guidFieldNodes := schema.SelectNodes("element/complexType/choice/element/complexType/sequence/element[@DataType='System.Guid']")

	if len(guidFieldNodes) != 3 {
		t.Fatalf("SampleMetadata.xml schema expected to contain 3 fields with Guid type, received: %d", len(guidFieldNodes))
	}
}

func TestItemLoad(t *testing.T) {
	_, found := doc.Root.Item["SchemaVersion"]

	if !found {
		t.Fatalf("SampleMetadata.xml expected to have \"SchemaVersion\" node, found none")
	}
}

func TestItemsLoad(t *testing.T) {
	root := doc.Root
	node := root.FirstChild()
	var lastNodeName string
	var nameCount int

	for node != nil {
		if lastNodeName != node.Name {
			if len(lastNodeName) > 0 {
				if nameCount != len(root.Items[lastNodeName]) {
					t.Fatalf("Items count for \"%s\" elements is %d, expected %d", lastNodeName, len(root.Items[lastNodeName]), nameCount)
				}
			}

			fmt.Println(node.Name)
			lastNodeName = node.Name
			nameCount = 1
		} else {
			nameCount++
		}

		node = node.Next
	}
}

func TestReverseEnumeration(t *testing.T) {
	node := doc.Root.LastChild()
	count := 0

	for node != nil {
		count++
		node = node.Previous
	}

	if count != 138 {
		t.Fatalf("SampleMetadata.xml expected to have 138 root child nodes")
	}
}
