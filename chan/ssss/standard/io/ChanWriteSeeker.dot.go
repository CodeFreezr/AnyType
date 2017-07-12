// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriteSeekerChan returns a new open channel
// (simply a 'chan io.WriteSeeker' that is).
//
// Note: No 'WriteSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriteSeekerPipelineStartsHere := MakeWriteSeekerChan()
//	// ... lot's of code to design and build Your favourite "myWriteSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriteSeekerPipelineStartsHere <- drop
//	}
//	close(myWriteSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriteSeekerBuffer) the channel is unbuffered.
//
func MakeWriteSeekerChan() chan io.WriteSeeker {
	return make(chan io.WriteSeeker)
}

// ChanWriteSeeker returns a channel to receive all inputs before close.
func ChanWriteSeeker(inp ...io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanWriteSeekerSlice returns a channel to receive all inputs before close.
func ChanWriteSeekerSlice(inp ...[]io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
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

// JoinWriteSeeker
func JoinWriteSeeker(out chan<- io.WriteSeeker, inp ...io.WriteSeeker) chan struct{} {
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

// JoinWriteSeekerSlice
func JoinWriteSeekerSlice(out chan<- io.WriteSeeker, inp ...[]io.WriteSeeker) chan struct{} {
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

// JoinWriteSeekerChan
func JoinWriteSeekerChan(out chan<- io.WriteSeeker, inp <-chan io.WriteSeeker) chan struct{} {
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

// DoneWriteSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneWriteSeeker(inp <-chan io.WriteSeeker) chan struct{} {
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

// DoneWriteSeekerSlice returns a channel which will receive a slice
// of all the WriteSeekers received on inp channel before close.
// Unlike DoneWriteSeeker, a full slice is sent once, not just an event.
func DoneWriteSeekerSlice(inp <-chan io.WriteSeeker) chan []io.WriteSeeker {
	done := make(chan []io.WriteSeeker)
	go func() {
		defer close(done)
		WriteSeekerS := []io.WriteSeeker{}
		for i := range inp {
			WriteSeekerS = append(WriteSeekerS, i)
		}
		done <- WriteSeekerS
	}()
	return done
}

// DoneWriteSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriteSeekerFunc(inp <-chan io.WriteSeeker, act func(a io.WriteSeeker)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.WriteSeeker) { return }
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

// PipeWriteSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriteSeekerBuffer(inp <-chan io.WriteSeeker, cap int) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriteSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriteSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriteSeekerFunc(inp <-chan io.WriteSeeker, act func(a io.WriteSeeker) io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
	if act == nil {
		act = func(a io.WriteSeeker) io.WriteSeeker { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriteSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriteSeekerFork(inp <-chan io.WriteSeeker) (chan io.WriteSeeker, chan io.WriteSeeker) {
	out1 := make(chan io.WriteSeeker)
	out2 := make(chan io.WriteSeeker)
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