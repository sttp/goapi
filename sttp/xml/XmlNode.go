//******************************************************************************************************
//  XmlNode.go - Gbtc
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
	"encoding/xml"
)

// XmlNode represents a single node in an XmlDocument tree.
type XmlNode struct {
	// XMLName gets the Go encoding xml.Name of this node.
	XMLName xml.Name // This is a specially named required field name in Go XML parsing
	// Name is the name of this node.
	Name string
	// NameSpace is the namespace URI of this node, if any.
	Namespace string
	// InnerXml is the XML content of the child nodes of this node.
	InnerXml []byte `xml:",innerxml"`
	// ChildNodes is a the collection of child nodes of this node.
	ChildNodes []XmlNode `xml:",any"`
	// Parent is the parent of this node. Value will be nil if node has no parent.
	Parent *XmlNode
	// Next gets the sibling-level node immediately following this node, or nil if there is none.
	Next *XmlNode
	// Previous gets the sibling-level node immediately preceding this node, or nil if there is none.
	Previous *XmlNode
	// Attributes is used to access the attributes of this node.
	Attributes map[string]string
	// Item is used to access the first child node of this node with the specified name.
	Item map[string]*XmlNode
	// Items is used to access all the child nodes of this node with the specified name.
	Items map[string][]*XmlNode
	// Level is the current node depth.
	Level int
	// Owner is the XmlDocument to which the current node belongs.
	Owner *XmlDocument
}

// HasChildNodes gets a flag indicating whether this node has any child nodes.
func (xn *XmlNode) HasChildNodes() bool {
	return len(xn.ChildNodes) > 0
}

// FirstChild gets the first child of this node, or nil if there are no child nodes.
func (xn *XmlNode) FirstChild() *XmlNode {
	if xn.HasChildNodes() {
		return &xn.ChildNodes[0]
	}

	return nil
}

// LastChild gets the last child of this node, or nil if there are not child nodes.
func (xn *XmlNode) LastChild() *XmlNode {
	if xn.HasChildNodes() {
		return &xn.ChildNodes[len(xn.ChildNodes)-1]
	}

	return nil
}
