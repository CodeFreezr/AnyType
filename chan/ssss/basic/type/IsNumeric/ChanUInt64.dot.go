// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt64Chan returns a new open channel
// (simply a 'chan uint64' that is).
//
// Note: No 'UInt64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt64PipelineStartsHere := MakeUInt64Chan()
//	// ... lot's of code to design and build Your favourite "myUInt64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt64PipelineStartsHere <- drop
//	}
//	close(myUInt64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt64Buffer) the channel is unbuffered.
//
func MakeUInt64Chan() chan uint64 {
	return make(chan uint64)
}

// ChanUInt64 returns a channel to receive all inputs before close.
func ChanUInt64(inp ...uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanUInt64Slice returns a channel to receive all inputs before close.
func ChanUInt64Slice(inp ...[]uint64) chan uint64 {
	out := make(chan uint64)
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

// JoinUInt64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64(out chan<- uint64, inp ...uint64) chan struct{} {
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

// JoinUInt64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64Slice(out chan<- uint64, inp ...[]uint64) chan struct{} {
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

// JoinUInt64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64Chan(out chan<- uint64, inp <-chan uint64) chan struct{} {
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

// DoneUInt64 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt64(inp <-chan uint64) chan struct{} {
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

// DoneUInt64Slice returns a channel which will receive a slice
// of all the UInt64s received on inp channel before close.
// Unlike DoneUInt64, a full slice is sent once, not just an event.
func DoneUInt64Slice(inp <-chan uint64) chan []uint64 {
	done := make(chan []uint64)
	go func() {
		defer close(done)
		UInt64S := []uint64{}
		for i := range inp {
			UInt64S = append(UInt64S, i)
		}
		done <- UInt64S
	}()
	return done
}

// DoneUInt64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt64Func(inp <-chan uint64, act func(a uint64)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a uint64) { return }
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

// PipeUInt64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt64Buffer(inp <-chan uint64, cap int) chan uint64 {
	out := make(chan uint64, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeUInt64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt64Func(inp <-chan uint64, act func(a uint64) uint64) chan uint64 {
	out := make(chan uint64)
	if act == nil {
		act = func(a uint64) uint64 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeUInt64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt64Fork(inp <-chan uint64) (chan uint64, chan uint64) {
	out1 := make(chan uint64)
	out2 := make(chan uint64)
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
