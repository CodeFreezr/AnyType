// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeSeekerChan returns a new open channel
// (simply a 'chan io.Seeker' that is).
//
// Note: No 'Seeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySeekerPipelineStartsHere := MakeSeekerChan()
//	// ... lot's of code to design and build Your favourite "mySeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySeekerPipelineStartsHere <- drop
//	}
//	close(mySeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSeekerBuffer) the channel is unbuffered.
//
func MakeSeekerChan() chan io.Seeker {
	return make(chan io.Seeker)
}

// ChanSeeker returns a channel to receive all inputs before close.
func ChanSeeker(inp ...io.Seeker) chan io.Seeker {
	out := make(chan io.Seeker)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanSeekerSlice returns a channel to receive all inputs before close.
func ChanSeekerSlice(inp ...[]io.Seeker) chan io.Seeker {
	out := make(chan io.Seeker)
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

// JoinSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeeker(out chan<- io.Seeker, inp ...io.Seeker) chan struct{} {
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

// JoinSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeekerSlice(out chan<- io.Seeker, inp ...[]io.Seeker) chan struct{} {
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

// JoinSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeekerChan(out chan<- io.Seeker, inp <-chan io.Seeker) chan struct{} {
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

// DoneSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneSeeker(inp <-chan io.Seeker) chan struct{} {
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

// DoneSeekerSlice returns a channel which will receive a slice
// of all the Seekers received on inp channel before close.
// Unlike DoneSeeker, a full slice is sent once, not just an event.
func DoneSeekerSlice(inp <-chan io.Seeker) chan []io.Seeker {
	done := make(chan []io.Seeker)
	go func() {
		defer close(done)
		SeekerS := []io.Seeker{}
		for i := range inp {
			SeekerS = append(SeekerS, i)
		}
		done <- SeekerS
	}()
	return done
}

// DoneSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.Seeker) { return }
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

// PipeSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSeekerBuffer(inp <-chan io.Seeker, cap int) chan io.Seeker {
	out := make(chan io.Seeker, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker) io.Seeker) chan io.Seeker {
	out := make(chan io.Seeker)
	if act == nil {
		act = func(a io.Seeker) io.Seeker { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSeekerFork(inp <-chan io.Seeker) (chan io.Seeker, chan io.Seeker) {
	out1 := make(chan io.Seeker)
	out2 := make(chan io.Seeker)
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
