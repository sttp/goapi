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

	"github.com/sttp/goapi/sttp/metadata"
)

var doc XmlDocument

func init() {
	doc.LoadXmlFromFile("../../test/SampleMetadata.xml")
}

func TestRootLevel(t *testing.T) {
	if doc.Root.Level != 0 {
		t.Fatalf("Root level in document tree should be zero")
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

	if schema.Namespace != metadata.XmlSchemaNamespace {
		t.Fatalf("SampleMetadata.xml expected to have schema namespace of \"%s\", received: \"%s\"", metadata.XmlSchemaNamespace, schema.Namespace)
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
