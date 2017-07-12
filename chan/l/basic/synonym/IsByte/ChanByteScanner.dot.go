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

type DChByteScanner struct { // demand channel
	dat chan io.ByteScanner
	req chan struct{}
}

func MakeDemandByteScannerChan() *DChByteScanner {
	d := new(DChByteScanner)
	d.dat = make(chan io.ByteScanner)
	d.req = make(chan struct{})
	return d
}

func MakeDemandByteScannerBuff(cap int) *DChByteScanner {
	d := new(DChByteScanner)
	d.dat = make(chan io.ByteScanner, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByteScanner is the send function - aka "MyKind <- some ByteScanner"
func (c *DChByteScanner) ProvideByteScanner(dat io.ByteScanner) {
	<-c.req
	c.dat <- dat
}

// RequestByteScanner is the receive function - aka "some ByteScanner <- MyKind"
func (c *DChByteScanner) RequestByteScanner() (dat io.ByteScanner) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByteScanner is the comma-ok multi-valued form of RequestByteScanner and
// reports whether a received value was sent before the ByteScanner channel was closed.
func (c *DChByteScanner) TryByteScanner() (dat io.ByteScanner, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len