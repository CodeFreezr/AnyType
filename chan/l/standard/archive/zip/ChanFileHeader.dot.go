// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

type FileHeaderChan interface { // bidirectional channel
	FileHeaderROnlyChan // aka "<-chan" - receive only
	FileHeaderSOnlyChan // aka "chan<-" - send only
}

type FileHeaderROnlyChan interface { // receive-only channel
	RequestFileHeader() (dat zip.FileHeader)        // the receive function - aka "some-new-FileHeader-var := <-MyKind"
	TryFileHeader() (dat zip.FileHeader, open bool) // the multi-valued comma-ok receive function - aka "some-new-FileHeader-var, ok := <-MyKind"
}

type FileHeaderSOnlyChan interface { // send-only channel
	ProvideFileHeader(dat zip.FileHeader) // the send function - aka "MyKind <- some FileHeader"
}

type DChFileHeader struct { // demand channel
	dat chan zip.FileHeader
	req chan struct{}
}

func MakeDemandFileHeaderChan() *DChFileHeader {
	d := new(DChFileHeader)
	d.dat = make(chan zip.FileHeader)
	d.req = make(chan struct{})
	return d
}

func MakeDemandFileHeaderBuff(cap int) *DChFileHeader {
	d := new(DChFileHeader)
	d.dat = make(chan zip.FileHeader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFileHeader is the send function - aka "MyKind <- some FileHeader"
func (c *DChFileHeader) ProvideFileHeader(dat zip.FileHeader) {
	<-c.req
	c.dat <- dat
}

// RequestFileHeader is the receive function - aka "some FileHeader <- MyKind"
func (c *DChFileHeader) RequestFileHeader() (dat zip.FileHeader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFileHeader is the comma-ok multi-valued form of RequestFileHeader and
// reports whether a received value was sent before the FileHeader channel was closed.
func (c *DChFileHeader) TryFileHeader() (dat zip.FileHeader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len