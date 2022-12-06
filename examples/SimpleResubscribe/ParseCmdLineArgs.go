//******************************************************************************************************
//  ParseCmdLineArgs.go - Gbtc
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
//  09/30/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func parseCmdLineArgs() string {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("    SimpleSubscribe HOSTNAME PORT")
		os.Exit(1)
	}

	hostname := args[1]
	port, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("Invalid port number \"%s\": %s\n", args[1], err.Error())
		os.Exit(2)
	}

	if port < 1 || port > math.MaxUint16 {
		fmt.Printf("Port number \"%s\" is out of range: must be 1 to %d\n", args[1], math.MaxUint16)
		os.Exit(2)
	}

	return hostname + ":" + strconv.Itoa(port)
}

func readKey() rune {
	r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
	return r
}
