// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsPathSChan returns a new open channel
// (simply a 'chan fs.FsPathS' that is).
//
// Note: No 'FsPathS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsPathSPipelineStartsHere := MakeFsPathSChan()
//	// ... lot's of code to design and build Your favourite "myFsPathSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsPathSPipelineStartsHere <- drop
//	}
//	close(myFsPathSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsPathSBuffer) the channel is unbuffered.
//
func MakeFsPathSChan() chan fs.FsPathS {
	return make(chan fs.FsPathS)
}

// ChanFsPathS returns a channel to receive all inputs before close.
func ChanFsPathS(inp ...fs.FsPathS) chan fs.FsPathS {
	out := make(chan fs.FsPathS)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsPathSSlice returns a channel to receive all inputs before close.
func ChanFsPathSSlice(inp ...[]fs.FsPathS) chan fs.FsPathS {
	out := make(chan fs.FsPathS)
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

// JoinFsPathS
func JoinFsPathS(out chan<- fs.FsPathS, inp ...fs.FsPathS) chan struct{} {
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

// JoinFsPathSSlice
func JoinFsPathSSlice(out chan<- fs.FsPathS, inp ...[]fs.FsPathS) chan struct{} {
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

// JoinFsPathSChan
func JoinFsPathSChan(out chan<- fs.FsPathS, inp <-chan fs.FsPathS) chan struct{} {
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

// DoneFsPathS returns a channel to receive one signal before close after inp has been drained.
func DoneFsPathS(inp <-chan fs.FsPathS) chan struct{} {
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

// DoneFsPathSSlice returns a channel which will receive a slice
// of all the FsPathSs received on inp channel before close.
// Unlike DoneFsPathS, a full slice is sent once, not just an event.
func DoneFsPathSSlice(inp <-chan fs.FsPathS) chan []fs.FsPathS {
	done := make(chan []fs.FsPathS)
	go func() {
		defer close(done)
		FsPathSS := []fs.FsPathS{}
		for i := range inp {
			FsPathSS = append(FsPathSS, i)
		}
		done <- FsPathSS
	}()
	return done
}

// DoneFsPathSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsPathSFunc(inp <-chan fs.FsPathS, act func(a fs.FsPathS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.FsPathS) { return }
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

// PipeFsPathSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsPathSBuffer(inp <-chan fs.FsPathS, cap int) chan fs.FsPathS {
	out := make(chan fs.FsPathS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsPathSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsPathSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsPathSFunc(inp <-chan fs.FsPathS, act func(a fs.FsPathS) fs.FsPathS) chan fs.FsPathS {
	out := make(chan fs.FsPathS)
	if act == nil {
		act = func(a fs.FsPathS) fs.FsPathS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsPathSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsPathSFork(inp <-chan fs.FsPathS) (chan fs.FsPathS, chan fs.FsPathS) {
	out1 := make(chan fs.FsPathS)
	out2 := make(chan fs.FsPathS)
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