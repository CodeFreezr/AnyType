// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

type BufferChan interface { // bidirectional channel
	BufferROnlyChan // aka "<-chan" - receive only
	BufferSOnlyChan // aka "chan<-" - send only
}

type BufferROnlyChan interface { // receive-only channel
	RequestBuffer() (dat bytes.Buffer)        // the receive function - aka "some-new-Buffer-var := <-MyKind"
	TryBuffer() (dat bytes.Buffer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Buffer-var, ok := <-MyKind"
}

type BufferSOnlyChan interface { // send-only channel
	ProvideBuffer(dat bytes.Buffer) // the send function - aka "MyKind <- some Buffer"
}

type DChBuffer struct { // demand channel
	dat chan bytes.Buffer
	req chan struct{}
}

func MakeDemandBufferChan() *DChBuffer {
	d := new(DChBuffer)
	d.dat = make(chan bytes.Buffer)
	d.req = make(chan struct{})
	return d
}

func MakeDemandBufferBuff(cap int) *DChBuffer {
	d := new(DChBuffer)
	d.dat = make(chan bytes.Buffer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideBuffer is the send function - aka "MyKind <- some Buffer"
func (c *DChBuffer) ProvideBuffer(dat bytes.Buffer) {
	<-c.req
	c.dat <- dat
}

// RequestBuffer is the receive function - aka "some Buffer <- MyKind"
func (c *DChBuffer) RequestBuffer() (dat bytes.Buffer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryBuffer is the comma-ok multi-valued form of RequestBuffer and
// reports whether a received value was sent before the Buffer channel was closed.
func (c *DChBuffer) TryBuffer() (dat bytes.Buffer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len