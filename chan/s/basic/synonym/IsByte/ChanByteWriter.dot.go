// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteWriterChan represents a
// bidirectional
// channel
type ByteWriterChan interface {
	ByteWriterROnlyChan // aka "<-chan" - receive only
	ByteWriterSOnlyChan // aka "chan<-" - send only
}

// ByteWriterROnlyChan represents a
// receive-only
// channel
type ByteWriterROnlyChan interface {
	RequestByteWriter() (dat io.ByteWriter)        // the receive function - aka "MyByteWriter := <-MyByteWriterROnlyChan"
	TryByteWriter() (dat io.ByteWriter, open bool) // the multi-valued comma-ok receive function - aka "MyByteWriter, ok := <-MyByteWriterROnlyChan"
}

// ByteWriterSOnlyChan represents a
// send-only
// channel
type ByteWriterSOnlyChan interface {
	ProvideByteWriter(dat io.ByteWriter) // the send function - aka "MyKind <- some ByteWriter"
}

// SChByteWriter is a supply channel
type SChByteWriter struct {
	dat chan io.ByteWriter
	// req chan struct{}
}

// MakeSupplyByteWriterChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyByteWriterChan() *SChByteWriter {
	d := new(SChByteWriter)
	d.dat = make(chan io.ByteWriter)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyByteWriterBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyByteWriterBuff(cap int) *SChByteWriter {
	d := new(SChByteWriter)
	d.dat = make(chan io.ByteWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByteWriter is the send function - aka "MyKind <- some ByteWriter"
func (c *SChByteWriter) ProvideByteWriter(dat io.ByteWriter) {
	// .req
	c.dat <- dat
}

// RequestByteWriter is the receive function - aka "some ByteWriter <- MyKind"
func (c *SChByteWriter) RequestByteWriter() (dat io.ByteWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByteWriter is the comma-ok multi-valued form of RequestByteWriter and
// reports whether a received value was sent before the ByteWriter channel was closed.
func (c *SChByteWriter) TryByteWriter() (dat io.ByteWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
