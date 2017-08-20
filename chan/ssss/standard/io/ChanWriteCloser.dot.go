// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriteCloserChan returns a new open channel
// (simply a 'chan io.WriteCloser' that is).
//
// Note: No 'WriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriteCloserPipelineStartsHere := MakeWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriteCloserPipelineStartsHere <- drop
//	}
//	close(myWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriteCloserBuffer) the channel is unbuffered.
//
func MakeWriteCloserChan() chan io.WriteCloser {
	return make(chan io.WriteCloser)
}

// ChanWriteCloser returns a channel to receive all inputs before close.
func ChanWriteCloser(inp ...io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanWriteCloserSlice returns a channel to receive all inputs before close.
func ChanWriteCloserSlice(inp ...[]io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
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

// JoinWriteCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloser(out chan<- io.WriteCloser, inp ...io.WriteCloser) chan struct{} {
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

// JoinWriteCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloserSlice(out chan<- io.WriteCloser, inp ...[]io.WriteCloser) chan struct{} {
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

// JoinWriteCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloserChan(out chan<- io.WriteCloser, inp <-chan io.WriteCloser) chan struct{} {
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

// DoneWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneWriteCloser(inp <-chan io.WriteCloser) chan struct{} {
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

// DoneWriteCloserSlice returns a channel which will receive a slice
// of all the WriteClosers received on inp channel before close.
// Unlike DoneWriteCloser, a full slice is sent once, not just an event.
func DoneWriteCloserSlice(inp <-chan io.WriteCloser) chan []io.WriteCloser {
	done := make(chan []io.WriteCloser)
	go func() {
		defer close(done)
		WriteCloserS := []io.WriteCloser{}
		for i := range inp {
			WriteCloserS = append(WriteCloserS, i)
		}
		done <- WriteCloserS
	}()
	return done
}

// DoneWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.WriteCloser) { return }
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

// PipeWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriteCloserBuffer(inp <-chan io.WriteCloser, cap int) chan io.WriteCloser {
	out := make(chan io.WriteCloser, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser) io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
	if act == nil {
		act = func(a io.WriteCloser) io.WriteCloser { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriteCloserFork(inp <-chan io.WriteCloser) (chan io.WriteCloser, chan io.WriteCloser) {
	out1 := make(chan io.WriteCloser)
	out2 := make(chan io.WriteCloser)
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
