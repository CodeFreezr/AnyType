// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt8Chan interface { // bidirectional channel
	UInt8ROnlyChan // aka "<-chan" - receive only
	UInt8SOnlyChan // aka "chan<-" - send only
}

type UInt8ROnlyChan interface { // receive-only channel
	RequestUInt8() (dat uint8)        // the receive function - aka "some-new-UInt8-var := <-MyKind"
	TryUInt8() (dat uint8, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt8-var, ok := <-MyKind"
}

type UInt8SOnlyChan interface { // send-only channel
	ProvideUInt8(dat uint8) // the send function - aka "MyKind <- some UInt8"
}