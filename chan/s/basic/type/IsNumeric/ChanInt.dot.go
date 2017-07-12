// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type IntChan interface { // bidirectional channel
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

type IntROnlyChan interface { // receive-only channel
	RequestInt() (dat int)        // the receive function - aka "some-new-Int-var := <-MyKind"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int-var, ok := <-MyKind"
}

type IntSOnlyChan interface { // send-only channel
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}

type SChInt struct { // supply channel
	dat chan int
	// req chan struct{}
}

func MakeSupplyIntChan() *SChInt {
	d := new(SChInt)
	d.dat = make(chan int)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyIntBuff(cap int) *SChInt {
	d := new(SChInt)
	d.dat = make(chan int, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt is the send function - aka "MyKind <- some Int"
func (c *SChInt) ProvideInt(dat int) {
	// .req
	c.dat <- dat
}

// RequestInt is the receive function - aka "some Int <- MyKind"
func (c *SChInt) RequestInt() (dat int) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt is the comma-ok multi-valued form of RequestInt and
// reports whether a received value was sent before the Int channel was closed.
func (c *SChInt) TryInt() (dat int, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len