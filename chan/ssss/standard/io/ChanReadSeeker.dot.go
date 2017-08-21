// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadSeekerChan returns a new open channel
// (simply a 'chan io.ReadSeeker' that is).
//
// Note: No 'ReadSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadSeekerPipelineStartsHere := MakeReadSeekerChan()
//	// ... lot's of code to design and build Your favourite "myReadSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadSeekerPipelineStartsHere <- drop
//	}
//	close(myReadSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadSeekerBuffer) the channel is unbuffered.
//
func MakeReadSeekerChan() chan io.ReadSeeker {
	return make(chan io.ReadSeeker)
}

// ChanReadSeeker returns a channel to receive all inputs before close.
func ChanReadSeeker(inp ...io.ReadSeeker) chan io.ReadSeeker {
	out := make(chan io.ReadSeeker)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanReadSeekerSlice returns a channel to receive all inputs before close.
func ChanReadSeekerSlice(inp ...[]io.ReadSeeker) chan io.ReadSeeker {
	out := make(chan io.ReadSeeker)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanReadSeekerFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadSeekerFuncNok(act func() (io.ReadSeeker, bool)) <-chan io.ReadSeeker {
	out := make(chan io.ReadSeeker)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanReadSeekerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadSeekerFuncErr(act func() (io.ReadSeeker, error)) <-chan io.ReadSeeker {
	out := make(chan io.ReadSeeker)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinReadSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeeker(out chan<- io.ReadSeeker, inp ...io.ReadSeeker) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReadSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeekerSlice(out chan<- io.ReadSeeker, inp ...[]io.ReadSeeker) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReadSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeekerChan(out chan<- io.ReadSeeker, inp <-chan io.ReadSeeker) chan struct{} {
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

// DoneReadSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneReadSeeker(inp <-chan io.ReadSeeker) chan struct{} {
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

// DoneReadSeekerSlice returns a channel which will receive a slice
// of all the ReadSeekers received on inp channel before close.
// Unlike DoneReadSeeker, a full slice is sent once, not just an event.
func DoneReadSeekerSlice(inp <-chan io.ReadSeeker) chan []io.ReadSeeker {
	done := make(chan []io.ReadSeeker)
	go func() {
		defer close(done)
		ReadSeekerS := []io.ReadSeeker{}
		for i := range inp {
			ReadSeekerS = append(ReadSeekerS, i)
		}
		done <- ReadSeekerS
	}()
	return done
}

// DoneReadSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadSeekerFunc(inp <-chan io.ReadSeeker, act func(a io.ReadSeeker)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReadSeeker) { return }
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

// PipeReadSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadSeekerBuffer(inp <-chan io.ReadSeeker, cap int) chan io.ReadSeeker {
	out := make(chan io.ReadSeeker, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadSeekerFunc(inp <-chan io.ReadSeeker, act func(a io.ReadSeeker) io.ReadSeeker) chan io.ReadSeeker {
	out := make(chan io.ReadSeeker)
	if act == nil {
		act = func(a io.ReadSeeker) io.ReadSeeker { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadSeekerFork(inp <-chan io.ReadSeeker) (chan io.ReadSeeker, chan io.ReadSeeker) {
	out1 := make(chan io.ReadSeeker)
	out2 := make(chan io.ReadSeeker)
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

// ReadSeekerTube is the signature for a pipe function.
type ReadSeekerTube func(inp <-chan io.ReadSeeker, out <-chan io.ReadSeeker)

// ReadSeekerDaisy returns a channel to receive all inp after having passed thru tube.
func ReadSeekerDaisy(inp <-chan io.ReadSeeker, tube ReadSeekerTube) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go tube(inp, cha)
	return cha
}

// ReadSeekerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadSeekerDaisyChain(inp <-chan io.ReadSeeker, tubes ...ReadSeekerTube) (out <-chan io.ReadSeeker) {
	cha := inp
	for i := range tubes {
		cha = ReadSeekerDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
