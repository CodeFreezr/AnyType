// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
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
	RequestReader() (dat *strings.Reader)        // the receive function - aka "MyReader := <-MyReaderROnlyChan"
	TryReader() (dat *strings.Reader, open bool) // the multi-valued comma-ok receive function - aka "MyReader, ok := <-MyReaderROnlyChan"
}

// ReaderSOnlyChan represents a
// send-only
// channel
type ReaderSOnlyChan interface {
	ProvideReader(dat *strings.Reader) // the send function - aka "MyKind <- some Reader"
}

// DChReader is a demand channel
type DChReader struct {
	dat chan *strings.Reader
	req chan struct{}
}

// MakeDemandReaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReaderChan() *DChReader {
	d := new(DChReader)
	d.dat = make(chan *strings.Reader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReaderBuff(cap int) *DChReader {
	d := new(DChReader)
	d.dat = make(chan *strings.Reader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReader is the send function - aka "MyKind <- some Reader"
func (c *DChReader) ProvideReader(dat *strings.Reader) {
	<-c.req
	c.dat <- dat
}

// RequestReader is the receive function - aka "some Reader <- MyKind"
func (c *DChReader) RequestReader() (dat *strings.Reader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReader is the comma-ok multi-valued form of RequestReader and
// reports whether a received value was sent before the Reader channel was closed.
func (c *DChReader) TryReader() (dat *strings.Reader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChReader is a supply channel
type SChReader struct {
	dat chan *strings.Reader
	// req chan struct{}
}

// MakeSupplyReaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReaderChan() *SChReader {
	d := new(SChReader)
	d.dat = make(chan *strings.Reader)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReaderBuff(cap int) *SChReader {
	d := new(SChReader)
	d.dat = make(chan *strings.Reader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReader is the send function - aka "MyKind <- some Reader"
func (c *SChReader) ProvideReader(dat *strings.Reader) {
	// .req
	c.dat <- dat
}

// RequestReader is the receive function - aka "some Reader <- MyKind"
func (c *SChReader) RequestReader() (dat *strings.Reader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReader is the comma-ok multi-valued form of RequestReader and
// reports whether a received value was sent before the Reader channel was closed.
func (c *SChReader) TryReader() (dat *strings.Reader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
