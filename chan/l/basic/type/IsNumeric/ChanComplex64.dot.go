// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Complex64Chan interface { // bidirectional channel
	Complex64ROnlyChan // aka "<-chan" - receive only
	Complex64SOnlyChan // aka "chan<-" - send only
}

type Complex64ROnlyChan interface { // receive-only channel
	RequestComplex64() (dat complex64)        // the receive function - aka "some-new-Complex64-var := <-MyKind"
	TryComplex64() (dat complex64, open bool) // the multi-valued comma-ok receive function - aka "some-new-Complex64-var, ok := <-MyKind"
}

type Complex64SOnlyChan interface { // send-only channel
	ProvideComplex64(dat complex64) // the send function - aka "MyKind <- some Complex64"
}

type DChComplex64 struct { // demand channel
	dat chan complex64
	req chan struct{}
}

func MakeDemandComplex64Chan() *DChComplex64 {
	d := new(DChComplex64)
	d.dat = make(chan complex64)
	d.req = make(chan struct{})
	return d
}

func MakeDemandComplex64Buff(cap int) *DChComplex64 {
	d := new(DChComplex64)
	d.dat = make(chan complex64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideComplex64 is the send function - aka "MyKind <- some Complex64"
func (c *DChComplex64) ProvideComplex64(dat complex64) {
	<-c.req
	c.dat <- dat
}

// RequestComplex64 is the receive function - aka "some Complex64 <- MyKind"
func (c *DChComplex64) RequestComplex64() (dat complex64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryComplex64 is the comma-ok multi-valued form of RequestComplex64 and
// reports whether a received value was sent before the Complex64 channel was closed.
func (c *DChComplex64) TryComplex64() (dat complex64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len