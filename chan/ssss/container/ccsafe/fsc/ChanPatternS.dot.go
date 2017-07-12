// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakePatternSChan returns a new open channel
// (simply a 'chan fs.PatternS' that is).
//
// Note: No 'PatternS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatternSPipelineStartsHere := MakePatternSChan()
//	// ... lot's of code to design and build Your favourite "myPatternSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatternSPipelineStartsHere <- drop
//	}
//	close(myPatternSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatternSBuffer) the channel is unbuffered.
//
func MakePatternSChan() chan fs.PatternS {
	return make(chan fs.PatternS)
}

// ChanPatternS returns a channel to receive all inputs before close.
func ChanPatternS(inp ...fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanPatternSSlice returns a channel to receive all inputs before close.
func ChanPatternSSlice(inp ...[]fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
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

// JoinPatternS
func JoinPatternS(out chan<- fs.PatternS, inp ...fs.PatternS) chan struct{} {
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

// JoinPatternSSlice
func JoinPatternSSlice(out chan<- fs.PatternS, inp ...[]fs.PatternS) chan struct{} {
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

// JoinPatternSChan
func JoinPatternSChan(out chan<- fs.PatternS, inp <-chan fs.PatternS) chan struct{} {
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

// DonePatternS returns a channel to receive one signal before close after inp has been drained.
func DonePatternS(inp <-chan fs.PatternS) chan struct{} {
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

// DonePatternSSlice returns a channel which will receive a slice
// of all the PatternSs received on inp channel before close.
// Unlike DonePatternS, a full slice is sent once, not just an event.
func DonePatternSSlice(inp <-chan fs.PatternS) chan []fs.PatternS {
	done := make(chan []fs.PatternS)
	go func() {
		defer close(done)
		PatternSS := []fs.PatternS{}
		for i := range inp {
			PatternSS = append(PatternSS, i)
		}
		done <- PatternSS
	}()
	return done
}

// DonePatternSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.PatternS) { return }
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

// PipePatternSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternSBuffer(inp <-chan fs.PatternS, cap int) chan fs.PatternS {
	out := make(chan fs.PatternS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipePatternSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
	if act == nil {
		act = func(a fs.PatternS) fs.PatternS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipePatternSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternSFork(inp <-chan fs.PatternS) (chan fs.PatternS, chan fs.PatternS) {
	out1 := make(chan fs.PatternS)
	out2 := make(chan fs.PatternS)
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