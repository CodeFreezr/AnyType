// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriterAtChan returns a new open channel
// (simply a 'chan io.WriterAt' that is).
//
// Note: No 'WriterAt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterAtPipelineStartsHere := MakeWriterAtChan()
//	// ... lot's of code to design and build Your favourite "myWriterAtWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterAtPipelineStartsHere <- drop
//	}
//	close(myWriterAtPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterAtBuffer) the channel is unbuffered.
//
func MakeWriterAtChan() (out chan io.WriterAt) {
	return make(chan io.WriterAt)
}

// ChanWriterAt returns a channel to receive all inputs before close.
func ChanWriterAt(inp ...io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	go func(out chan<- io.WriterAt, inp ...io.WriterAt) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanWriterAtSlice returns a channel to receive all inputs before close.
func ChanWriterAtSlice(inp ...[]io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	go func(out chan<- io.WriterAt, inp ...[]io.WriterAt) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinWriterAt sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterAt(out chan<- io.WriterAt, inp ...io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.WriterAt, inp ...io.WriterAt) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinWriterAtSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterAtSlice(out chan<- io.WriterAt, inp ...[]io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.WriterAt, inp ...[]io.WriterAt) {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinWriterAtChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterAtChan(out chan<- io.WriterAt, inp <-chan io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.WriterAt, inp <-chan io.WriterAt) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneWriterAt returns a channel to receive one signal before close after inp has been drained.
func DoneWriterAt(inp <-chan io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.WriterAt) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneWriterAtSlice returns a channel which will receive a slice
// of all the WriterAts received on inp channel before close.
// Unlike DoneWriterAt, a full slice is sent once, not just an event.
func DoneWriterAtSlice(inp <-chan io.WriterAt) (done <-chan []io.WriterAt) {
	cha := make(chan []io.WriterAt)
	go func(inp <-chan io.WriterAt, done chan<- []io.WriterAt) {
		defer close(done)
		WriterAtS := []io.WriterAt{}
		for i := range inp {
			WriterAtS = append(WriterAtS, i)
		}
		done <- WriterAtS
	}(inp, cha)
	return cha
}

// DoneWriterAtFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterAtFunc(inp <-chan io.WriterAt, act func(a io.WriterAt)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.WriterAt) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.WriterAt, act func(a io.WriterAt)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeWriterAtBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterAtBuffer(inp <-chan io.WriterAt, cap int) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt, cap)
	go func(out chan<- io.WriterAt, inp <-chan io.WriterAt) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeWriterAtFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterAtMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterAtFunc(inp <-chan io.WriterAt, act func(a io.WriterAt) io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	if act == nil {
		act = func(a io.WriterAt) io.WriterAt { return a }
	}
	go func(out chan<- io.WriterAt, inp <-chan io.WriterAt, act func(a io.WriterAt) io.WriterAt) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeWriterAtFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterAtFork(inp <-chan io.WriterAt) (out1, out2 <-chan io.WriterAt) {
	cha1 := make(chan io.WriterAt)
	cha2 := make(chan io.WriterAt)
	go func(out1, out2 chan<- io.WriterAt, inp <-chan io.WriterAt) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// WriterAtTube is the signature for a pipe function.
type WriterAtTube func(inp <-chan io.WriterAt, out <-chan io.WriterAt)

// WriterAtdaisy returns a channel to receive all inp after having passed thru tube.
func WriterAtdaisy(inp <-chan io.WriterAt, tube WriterAtTube) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	go tube(inp, cha)
	return cha
}

// WriterAtDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterAtDaisyChain(inp <-chan io.WriterAt, tubes ...WriterAtTube) (out <-chan io.WriterAt) {
	cha := inp
	for _, tube := range tubes {
		cha = WriterAtdaisy(cha, tube)
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
