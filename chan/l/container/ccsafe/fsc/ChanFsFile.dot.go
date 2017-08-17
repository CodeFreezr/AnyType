// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsFileChan interface { // bidirectional channel
	FsFileROnlyChan // aka "<-chan" - receive only
	FsFileSOnlyChan // aka "chan<-" - send only
}

type FsFileROnlyChan interface { // receive-only channel
	RequestFsFile() (dat *fs.FsFile)        // the receive function - aka "some-new-FsFile-var := <-MyKind"
	TryFsFile() (dat *fs.FsFile, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFile-var, ok := <-MyKind"
}

type FsFileSOnlyChan interface { // send-only channel
	ProvideFsFile(dat *fs.FsFile) // the send function - aka "MyKind <- some FsFile"
}

type DChFsFile struct { // demand channel
	dat chan *fs.FsFile
	req chan struct{}
}

func MakeDemandFsFileChan() *DChFsFile {
	d := new(DChFsFile)
	d.dat = make(chan *fs.FsFile)
	d.req = make(chan struct{})
	return d
}

func MakeDemandFsFileBuff(cap int) *DChFsFile {
	d := new(DChFsFile)
	d.dat = make(chan *fs.FsFile, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsFile is the send function - aka "MyKind <- some FsFile"
func (c *DChFsFile) ProvideFsFile(dat *fs.FsFile) {
	<-c.req
	c.dat <- dat
}

// RequestFsFile is the receive function - aka "some FsFile <- MyKind"
func (c *DChFsFile) RequestFsFile() (dat *fs.FsFile) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsFile is the comma-ok multi-valued form of RequestFsFile and
// reports whether a received value was sent before the FsFile channel was closed.
func (c *DChFsFile) TryFsFile() (dat *fs.FsFile, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
