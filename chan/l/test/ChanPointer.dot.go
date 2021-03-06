// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PointerChan represents a
// bidirectional
// channel
type PointerChan interface {
	PointerROnlyChan // aka "<-chan" - receive only
	PointerSOnlyChan // aka "chan<-" - send only
}

// PointerROnlyChan represents a
// receive-only
// channel
type PointerROnlyChan interface {
	RequestPointer() (dat *SomeType)        // the receive function - aka "MyPointer := <-MyPointerROnlyChan"
	TryPointer() (dat *SomeType, open bool) // the multi-valued comma-ok receive function - aka "MyPointer, ok := <-MyPointerROnlyChan"
}

// PointerSOnlyChan represents a
// send-only
// channel
type PointerSOnlyChan interface {
	ProvidePointer(dat *SomeType) // the send function - aka "MyKind <- some Pointer"
}

// DChPointer is a demand channel
type DChPointer struct {
	dat chan *SomeType
	req chan struct{}
}

// MakeDemandPointerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandPointerChan() *DChPointer {
	d := new(DChPointer)
	d.dat = make(chan *SomeType)
	d.req = make(chan struct{})
	return d
}

// MakeDemandPointerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandPointerBuff(cap int) *DChPointer {
	d := new(DChPointer)
	d.dat = make(chan *SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePointer is the send function - aka "MyKind <- some Pointer"
func (c *DChPointer) ProvidePointer(dat *SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestPointer is the receive function - aka "some Pointer <- MyKind"
func (c *DChPointer) RequestPointer() (dat *SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPointer is the comma-ok multi-valued form of RequestPointer and
// reports whether a received value was sent before the Pointer channel was closed.
func (c *DChPointer) TryPointer() (dat *SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
