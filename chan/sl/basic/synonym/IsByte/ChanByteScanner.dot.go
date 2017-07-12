// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteScannerChan interface { // bidirectional channel
	ByteScannerROnlyChan // aka "<-chan" - receive only
	ByteScannerSOnlyChan // aka "chan<-" - send only
}

type ByteScannerROnlyChan interface { // receive-only channel
	RequestByteScanner() (dat io.ByteScanner)        // the receive function - aka "some-new-ByteScanner-var := <-MyKind"
	TryByteScanner() (dat io.ByteScanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteScanner-var, ok := <-MyKind"
}

type ByteScannerSOnlyChan interface { // send-only channel
	ProvideByteScanner(dat io.ByteScanner) // the send function - aka "MyKind <- some ByteScanner"
}