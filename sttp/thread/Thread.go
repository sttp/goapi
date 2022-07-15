//******************************************************************************************************
//  Thread.go - Gbtc
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
//  09/13/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package thread

import (
	"sync"

	"github.com/tevino/abool/v2"
)

// Thread represents a thread-like wrapper for a Go routine.
type Thread struct {
	exec    func()
	mutex   sync.Mutex
	running abool.AtomicBool
}

// NewThread creates a new Thread.
func NewThread(exec func()) *Thread {
	return &Thread{exec: exec}
}

// Start causes the thread function to be scheduled for execution via a new Go routine.
func (thread *Thread) Start() {
	if thread.exec == nil {
		panic("thread has no execution function defined")
	}

	if thread.running.IsSet() {
		panic("thread is already running")
	}

	thread.mutex.Lock()
	go thread.run()
}

// TryStart attempts to cause the thread function to be scheduled for execution via a new Go routine.
// Returns true if the thread function was successfully scheduled for execution; otherwise, false.
func (thread *Thread) TryStart() bool {
	defer func() {
		// Ignore possible thread Start race panics, function will return false
		recover()
	}()

	if thread.exec != nil && thread.running.IsNotSet() {
		thread.Start()
		return true
	}

	return false
}

// Join blocks the calling thread until this Thread terminates.
func (thread *Thread) Join() {
	if thread.exec == nil || thread.running.IsNotSet() {
		return
	}

	thread.mutex.Lock()
	//lint:ignore SA2001 -- desired behavior
	thread.mutex.Unlock()
}

// IsRunning safely determines if the thread function is currently executing.
func (thread *Thread) IsRunning() bool {
	return thread.running.IsSet()
}

func (thread *Thread) run() {
	defer thread.running.UnSet()
	defer thread.mutex.Unlock()

	thread.running.Set()
	thread.exec()
}
