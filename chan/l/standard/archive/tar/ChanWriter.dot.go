// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

type WriterChan interface { // bidirectional channel
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

type WriterROnlyChan interface { // receive-only channel
	RequestWriter() (dat *tar.Writer)        // the receive function - aka "some-new-Writer-var := <-MyKind"
	TryWriter() (dat *tar.Writer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Writer-var, ok := <-MyKind"
}

type WriterSOnlyChan interface { // send-only channel
	ProvideWriter(dat *tar.Writer) // the send function - aka "MyKind <- some Writer"
}

type DChWriter struct { // demand channel
	dat chan *tar.Writer
	req chan struct{}
}

func MakeDemandWriterChan() *DChWriter {
	d := new(DChWriter)
	d.dat = make(chan *tar.Writer)
	d.req = make(chan struct{})
	return d
}

func MakeDemandWriterBuff(cap int) *DChWriter {
	d := new(DChWriter)
	d.dat = make(chan *tar.Writer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideWriter is the send function - aka "MyKind <- some Writer"
func (c *DChWriter) ProvideWriter(dat *tar.Writer) {
	<-c.req
	c.dat <- dat
}

// RequestWriter is the receive function - aka "some Writer <- MyKind"
func (c *DChWriter) RequestWriter() (dat *tar.Writer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryWriter is the comma-ok multi-valued form of RequestWriter and
// reports whether a received value was sent before the Writer channel was closed.
func (c *DChWriter) TryWriter() (dat *tar.Writer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len