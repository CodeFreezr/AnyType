// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsFileSChan returns a new open channel
// (simply a 'chan fs.FsFileS' that is).
//
// Note: No 'FsFileS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFileSPipelineStartsHere := MakeFsFileSChan()
//	// ... lot's of code to design and build Your favourite "myFsFileSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFileSPipelineStartsHere <- drop
//	}
//	close(myFsFileSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFileSBuffer) the channel is unbuffered.
//
func MakeFsFileSChan() chan fs.FsFileS {
	return make(chan fs.FsFileS)
}

// ChanFsFileS returns a channel to receive all inputs before close.
func ChanFsFileS(inp ...fs.FsFileS) chan fs.FsFileS {
	out := make(chan fs.FsFileS)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsFileSSlice returns a channel to receive all inputs before close.
func ChanFsFileSSlice(inp ...[]fs.FsFileS) chan fs.FsFileS {
	out := make(chan fs.FsFileS)
	go func() {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}()
	return out
}

// JoinFsFileS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileS(out chan<- fs.FsFileS, inp ...fs.FsFileS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFsFileSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileSSlice(out chan<- fs.FsFileS, inp ...[]fs.FsFileS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFsFileSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileSChan(out chan<- fs.FsFileS, inp <-chan fs.FsFileS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFsFileS returns a channel to receive one signal before close after inp has been drained.
func DoneFsFileS(inp <-chan fs.FsFileS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFsFileSSlice returns a channel which will receive a slice
// of all the FsFileSs received on inp channel before close.
// Unlike DoneFsFileS, a full slice is sent once, not just an event.
func DoneFsFileSSlice(inp <-chan fs.FsFileS) chan []fs.FsFileS {
	done := make(chan []fs.FsFileS)
	go func() {
		defer close(done)
		FsFileSS := []fs.FsFileS{}
		for i := range inp {
			FsFileSS = append(FsFileSS, i)
		}
		done <- FsFileSS
	}()
	return done
}

// DoneFsFileSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.FsFileS) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeFsFileSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFileSBuffer(inp <-chan fs.FsFileS, cap int) chan fs.FsFileS {
	out := make(chan fs.FsFileS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsFileSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFileSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS) fs.FsFileS) chan fs.FsFileS {
	out := make(chan fs.FsFileS)
	if act == nil {
		act = func(a fs.FsFileS) fs.FsFileS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsFileSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFileSFork(inp <-chan fs.FsFileS) (chan fs.FsFileS, chan fs.FsFileS) {
	out1 := make(chan fs.FsFileS)
	out2 := make(chan fs.FsFileS)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}
