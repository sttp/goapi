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
	"sync/atomic"
)

// Thread represents a thread-like wrapper for a Go routine.
type Thread struct {
	exec    func()
	mutex   sync.Mutex
	running int32
}

// NewThread creates a new Thread.
func NewThread(exec func()) *Thread {
	return &Thread{exec: exec}
}

// Start causes the thread function to be scheduled for execution via a new Go routine.
func (thread *Thread) Start() {
	if thread.exec == nil {
		panic("Thread has no execution function defined")
	}

	if thread.IsRunning() {
		panic("Thread is already running")
	}

	thread.mutex.Lock()
	go thread.run()
}

// Join blocks the calling thread until this Thread terminates.
func (thread *Thread) Join() {
	if thread.exec == nil || !thread.IsRunning() {
		return
	}

	thread.mutex.Lock()
	//lint:ignore SA2001 -- desired behavior
	thread.mutex.Unlock()
}

// IsRunning safely determines if the thread function is currently executing.
func (thread *Thread) IsRunning() bool {
	return atomic.LoadInt32(&(thread.running)) != 0
}

func (thread *Thread) setIsRunning(value bool) {
	var i int32 = 0

	if value {
		i = 1
	}

	atomic.StoreInt32(&(thread.running), i)
}

func (thread *Thread) run() {
	defer thread.setIsRunning(false)
	defer thread.mutex.Unlock()

	thread.setIsRunning(true)
	thread.exec()
}
