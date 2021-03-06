// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// ByteChan represents a
// bidirectional
// channel
type ByteChan interface {
	ByteROnlyChan // aka "<-chan" - receive only
	ByteSOnlyChan // aka "chan<-" - send only
}

// ByteROnlyChan represents a
// receive-only
// channel
type ByteROnlyChan interface {
	RequestByte() (dat byte)        // the receive function - aka "MyByte := <-MyByteROnlyChan"
	TryByte() (dat byte, open bool) // the multi-valued comma-ok receive function - aka "MyByte, ok := <-MyByteROnlyChan"
}

// ByteSOnlyChan represents a
// send-only
// channel
type ByteSOnlyChan interface {
	ProvideByte(dat byte) // the send function - aka "MyKind <- some Byte"
}

// DChByte is a demand channel
type DChByte struct {
	dat chan byte
	req chan struct{}
}

// MakeDemandByteChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandByteChan() *DChByte {
	d := new(DChByte)
	d.dat = make(chan byte)
	d.req = make(chan struct{})
	return d
}

// MakeDemandByteBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandByteBuff(cap int) *DChByte {
	d := new(DChByte)
	d.dat = make(chan byte, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByte is the send function - aka "MyKind <- some Byte"
func (c *DChByte) ProvideByte(dat byte) {
	<-c.req
	c.dat <- dat
}

// RequestByte is the receive function - aka "some Byte <- MyKind"
func (c *DChByte) RequestByte() (dat byte) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByte is the comma-ok multi-valued form of RequestByte and
// reports whether a received value was sent before the Byte channel was closed.
func (c *DChByte) TryByte() (dat byte, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
