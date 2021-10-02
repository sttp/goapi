//******************************************************************************************************
//  OrderByTerm.go - Gbtc
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
//  10/02/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package filterexpressions

import "github.com/sttp/goapi/sttp/data"

// OrderByTerm represents the elements parsed from a column specified in the "ORDER BY" keyword.
type OrderByTerm struct {
	// Column is the data column reference of the OrderByTerm.
	Column *data.DataColumn

	// Ascending is a flag that determines if the OrderByTerm is sorted in ascending order.
	Ascending bool

	// ExactMatch is a flag that determines if the OrderByTerm used an exact match comparison.
	ExactMatch bool
}
