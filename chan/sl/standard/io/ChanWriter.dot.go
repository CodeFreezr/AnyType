// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriterChan represents a
// bidirectional
// channel
type WriterChan interface {
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

// WriterROnlyChan represents a
// receive-only
// channel
type WriterROnlyChan interface {
	RequestWriter() (dat io.Writer)        // the receive function - aka "MyWriter := <-MyWriterROnlyChan"
	TryWriter() (dat io.Writer, open bool) // the multi-valued comma-ok receive function - aka "MyWriter, ok := <-MyWriterROnlyChan"
}

// WriterSOnlyChan represents a
// send-only
// channel
type WriterSOnlyChan interface {
	ProvideWriter(dat io.Writer) // the send function - aka "MyKind <- some Writer"
}
