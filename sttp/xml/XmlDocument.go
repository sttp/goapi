//******************************************************************************************************
//  XmlDocument.go - Gbtc
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
//  09/24/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package xml

import (
	"bytes"
	"encoding/xml"
	"os"
	"strings"
)

// XmlDocument represents XML data as a tree of XmlNode instances.
type XmlDocument struct {
	// Root is the starting XmlNode in a set of parsed XML data, or nil if there is not one.
	// This is the root node of the XMLDocument tree.
	Root XmlNode

	maxLevel int
}

func (xd *XmlDocument) traverse(nodes []XmlNode, parent *XmlNode) {
	length := len(nodes)
	parent.Item = make(map[string]*XmlNode, length)
	parent.Items = make(map[string][]*XmlNode, length)

	for i := 0; i < length; i++ {
		level := parent.Level + 1

		if level > xd.maxLevel {
			xd.maxLevel = level
		}

		node := &nodes[i]

		node.Name = node.XMLName.Local
		node.Namespace = node.XMLName.Space
		node.Parent = parent
		node.Level = level
		node.Owner = xd

		if _, found := parent.Item[node.Name]; !found {
			parent.Item[node.Name] = node
			parent.Items[node.Name] = []*XmlNode{node}
		} else {
			parent.Items[node.Name] = append(parent.Items[node.Name], node)
		}

		if i > 0 {
			node.Previous = &nodes[i-1]
			node.Previous.Next = node
		}

		xd.traverse(node.ChildNodes, node)
	}
}

// UnmarshalXML decodes a single XML element beginning with the given start element.
// This XmlNode targeted overriding implementation turns attributes into a map.
func (xn *XmlNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	xn.Attributes = make(map[string]string, len(start.Attr))
	xn.AttributeNamespaces = make(map[string]string, len(start.Attr))

	for _, v := range start.Attr {
		xn.Attributes[v.Name.Local] = v.Value
		xn.AttributeNamespaces[v.Name.Local] = v.Name.Space
	}

	// Separate type here prevents recursing into current UnmarshalXML for Decoder
	type node XmlNode
	return d.DecodeElement((*node)(xn), &start)
}

// LoadXml loads the XmlDocument from the specified XML data.
func (xd *XmlDocument) LoadXml(data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := xml.NewDecoder(buffer)

	err := decoder.Decode(&xd.Root)

	if err != nil {
		return err
	}

	xd.Root.Name = xd.Root.XMLName.Local
	xd.Root.Namespace = xd.Root.XMLName.Space
	xd.Root.Owner = xd

	xd.traverse(xd.Root.ChildNodes, &xd.Root)

	return nil
}

// LoadXmlFromFile loads the XmlDocument from the specified file name containing XML data.
func (xd *XmlDocument) LoadXmlFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	return xd.LoadXml(data)
}

// MaxDepth gets the maximum node depth for XmlNode instances in this XmlDocument tree.
func (xd *XmlDocument) MaxDepth() int {
	return xd.maxLevel + 1
}

// SelectNodes finds all nodes matching xpath expression starting at XmlDocument root.
// Predicates currently only support "=" matching.
func (xn *XmlDocument) SelectNodes(xpath string) []*XmlNode {

	if strings.HasPrefix(xpath, "//") && strings.HasPrefix(xpath[2:], xn.Root.Name+"/") {
		return xn.Root.SelectNodes(xpath[3+len(xn.Root.Name):])
	}

	return xn.Root.SelectNodes(xpath)
}
