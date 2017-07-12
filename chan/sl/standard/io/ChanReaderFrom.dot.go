// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReaderFromChan interface { // bidirectional channel
	ReaderFromROnlyChan // aka "<-chan" - receive only
	ReaderFromSOnlyChan // aka "chan<-" - send only
}

type ReaderFromROnlyChan interface { // receive-only channel
	RequestReaderFrom() (dat io.ReaderFrom)        // the receive function - aka "some-new-ReaderFrom-var := <-MyKind"
	TryReaderFrom() (dat io.ReaderFrom, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReaderFrom-var, ok := <-MyKind"
}

type ReaderFromSOnlyChan interface { // send-only channel
	ProvideReaderFrom(dat io.ReaderFrom) // the send function - aka "MyKind <- some ReaderFrom"
}