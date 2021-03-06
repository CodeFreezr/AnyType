// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

// ReaderChan represents a
// bidirectional
// channel
type ReaderChan interface {
	ReaderROnlyChan // aka "<-chan" - receive only
	ReaderSOnlyChan // aka "chan<-" - send only
}

// ReaderROnlyChan represents a
// receive-only
// channel
type ReaderROnlyChan interface {
	RequestReader() (dat bytes.Reader)        // the receive function - aka "MyReader := <-MyReaderROnlyChan"
	TryReader() (dat bytes.Reader, open bool) // the multi-valued comma-ok receive function - aka "MyReader, ok := <-MyReaderROnlyChan"
}

// ReaderSOnlyChan represents a
// send-only
// channel
type ReaderSOnlyChan interface {
	ProvideReader(dat bytes.Reader) // the send function - aka "MyKind <- some Reader"
}

// DChReader is a demand channel
type DChReader struct {
	dat chan bytes.Reader
	req chan struct{}
}

// MakeDemandReaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReaderChan() *DChReader {
	d := new(DChReader)
	d.dat = make(chan bytes.Reader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReaderBuff(cap int) *DChReader {
	d := new(DChReader)
	d.dat = make(chan bytes.Reader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReader is the send function - aka "MyKind <- some Reader"
func (c *DChReader) ProvideReader(dat bytes.Reader) {
	<-c.req
	c.dat <- dat
}

// RequestReader is the receive function - aka "some Reader <- MyKind"
func (c *DChReader) RequestReader() (dat bytes.Reader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReader is the comma-ok multi-valued form of RequestReader and
// reports whether a received value was sent before the Reader channel was closed.
func (c *DChReader) TryReader() (dat bytes.Reader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
