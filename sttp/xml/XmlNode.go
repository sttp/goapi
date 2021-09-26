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
	"strings"
)

// XmlNode represents a single node in an XmlDocument tree.
type XmlNode struct {
	// XMLName gets the Go encoding xml.Name of this node.
	XMLName xml.Name // This is a specially named required field in Go XML parsing
	// Name is the name of this node.
	Name string
	// NameSpace is the namespace URI of this node, if any.
	Namespace string
	// InnerXml is the XML content of the child nodes of this node.
	InnerXml []byte `xml:",innerxml"`
	// ChildNodes is a the collection of child nodes of this node.
	ChildNodes []XmlNode `xml:",any"`
	// Parent is the parent of this node. Value will be nil if node has no parent, e.g., XmlDocument.Root.
	Parent *XmlNode
	// Next gets the sibling-level node immediately following this node, or nil if there is none.
	Next *XmlNode
	// Previous gets the sibling-level node immediately preceding this node, or nil if there is none.
	Previous *XmlNode
	// Attributes is used to access the attributes of this node.
	Attributes map[string]string
	// AttributeNamespaces is used to access the attribute namespaces of this node.
	AttributeNamespaces map[string]string
	// Item is used to access the first child node of this node with the specified name.
	Item map[string]*XmlNode
	// Items is used to access the collection of child nodes of this node with the specified name.
	Items map[string][]*XmlNode
	// Level is the current node depth.
	Level int
	// Owner is the XmlDocument to which this current node belongs.
	Owner *XmlDocument
}

// Path gets the full path of this node within its XmlDocument tree.
func (xn *XmlNode) Path() string {
	var paths []string
	node := xn

	for node != nil {
		paths = append(paths, node.Name)
		node = node.Parent
	}

	var path strings.Builder
	path.WriteRune('/')

	for i := len(paths) - 1; i > -1; i-- {
		path.WriteRune('/')
		path.WriteString(paths[i])
	}

	return path.String()
}

// Value gets the InnerXml of this node as a string.
func (xn *XmlNode) Value() string {
	return string(xn.InnerXml)
}

// HasChildNodes gets a flag indicating if this node has any child nodes.
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

// GetChildNodes gets a slice of pointers to all child nodes of this node.
func (xn *XmlNode) GetChildNodes() []*XmlNode {
	count := len(xn.ChildNodes)
	childNodes := make([]*XmlNode, count)

	for i := 0; i < count; i++ {
		childNodes[i] = &xn.ChildNodes[i]
	}

	return childNodes
}

// SelectNodes finds all nodes matching xpath expression for each input node.
func SelectNodes(nodes []*XmlNode, xpath string) []*XmlNode {
	results := make([]*XmlNode, 0)

	for _, node := range nodes {
		results = append(results, node.SelectNodes(xpath)...)
	}

	return results
}

// SelectNodes finds all nodes matching xpath expression.
// Predicates currently only support "=" matching.
func (xn *XmlNode) SelectNodes(xpath string) []*XmlNode {
	// This is a simple XPath implementation, this can be expanded
	// as needed, perhaps using github.com/antchfx/xpath
	nodes := xn.GetChildNodes()
	exprs := strings.Split(xpath, "/")
	var results []*XmlNode

	for i := 0; i < len(exprs); i++ {
		expr := exprs[i]
		results = make([]*XmlNode, 0)

		for _, node := range nodes {
			// Look for predicate expression
			lbp := strings.IndexRune(expr, '[')
			rbp := strings.IndexRune(expr, ']')

			if lbp > -1 && rbp > -1 && rbp > lbp && len(expr) > 2 {
				name := expr[:lbp]
				predicate := expr[lbp+1 : rbp]

				if predicate[0] == '@' {
					if (node.Name == name || name == "*") && node.attributeMatch(predicate[1:]) {
						results = append(results, node)
					}
				} else {
					if (node.Name == name || name == "*") && node.childValueMatch(predicate) {
						results = append(results, node)
					}
				}
			} else if node.Name == expr || expr == "*" {
				results = append(results, node)
			}
		}

		if len(results) == 0 {
			break
		}

		if i < len(exprs)-1 {
			nodes = make([]*XmlNode, 0, len(results))

			// Check all macthing children of next level
			for _, result := range results {
				nodes = append(nodes, result.GetChildNodes()...)
			}
		}
	}

	return results
}

func (xn *XmlNode) attributeMatch(expr string) bool {
	parts := strings.Split(expr, "=")

	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}

	if parts[0] == "*" {
		if len(parts) == 1 {
			return true
		}

		// Check if any attribute value matches
		for _, value := range xn.Attributes {
			if value == removeQuotes(parts[1]) {
				return true
			}
		}
	} else {
		// Check if named attribute value matches
		if value, found := xn.Attributes[parts[0]]; found {
			if len(parts) == 2 {
				return value == removeQuotes(parts[1])
			}

			return true
		}
	}

	return false
}

func (xn *XmlNode) childValueMatch(expr string) bool {
	for _, node := range xn.GetChildNodes() {
		if node.valueMatch(expr) {
			return true
		}
	}

	return false
}

func (xn *XmlNode) valueMatch(expr string) bool {
	parts := strings.Split(expr, "=")

	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}

	if parts[0] == xn.Name || parts[0] == "*" {
		if len(parts) == 2 {
			return xn.Value() == removeQuotes(parts[1])
		}

		return true
	}

	return false
}

func removeQuotes(value string) string {
	if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") && len(value) > 2 {
		value = value[1 : len(value)-1]
	}

	return value
}
