// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriteSeekerChan interface { // bidirectional channel
	WriteSeekerROnlyChan // aka "<-chan" - receive only
	WriteSeekerSOnlyChan // aka "chan<-" - send only
}

type WriteSeekerROnlyChan interface { // receive-only channel
	RequestWriteSeeker() (dat io.WriteSeeker)        // the receive function - aka "some-new-WriteSeeker-var := <-MyKind"
	TryWriteSeeker() (dat io.WriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriteSeeker-var, ok := <-MyKind"
}

type WriteSeekerSOnlyChan interface { // send-only channel
	ProvideWriteSeeker(dat io.WriteSeeker) // the send function - aka "MyKind <- some WriteSeeker"
}