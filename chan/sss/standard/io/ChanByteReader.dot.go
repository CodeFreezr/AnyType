// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteReaderChan returns a new open channel
// (simply a 'chan io.ByteReader' that is).
//
// Note: No 'ByteReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteReaderPipelineStartsHere := MakeByteReaderChan()
//	// ... lot's of code to design and build Your favourite "myByteReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteReaderPipelineStartsHere <- drop
//	}
//	close(myByteReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteReaderBuffer) the channel is unbuffered.
//
func MakeByteReaderChan() (out chan io.ByteReader) {
	return make(chan io.ByteReader)
}

// ChanByteReader returns a channel to receive all inputs before close.
func ChanByteReader(inp ...io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go func(out chan<- io.ByteReader, inp ...io.ByteReader) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanByteReaderSlice returns a channel to receive all inputs before close.
func ChanByteReaderSlice(inp ...[]io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go func(out chan<- io.ByteReader, inp ...[]io.ByteReader) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinByteReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReader(out chan<- io.ByteReader, inp ...io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteReader, inp ...io.ByteReader) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinByteReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderSlice(out chan<- io.ByteReader, inp ...[]io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteReader, inp ...[]io.ByteReader) {
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

// JoinByteReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderChan(out chan<- io.ByteReader, inp <-chan io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteReader, inp <-chan io.ByteReader) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneByteReader returns a channel to receive one signal before close after inp has been drained.
func DoneByteReader(inp <-chan io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.ByteReader) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneByteReaderSlice returns a channel which will receive a slice
// of all the ByteReaders received on inp channel before close.
// Unlike DoneByteReader, a full slice is sent once, not just an event.
func DoneByteReaderSlice(inp <-chan io.ByteReader) (done <-chan []io.ByteReader) {
	cha := make(chan []io.ByteReader)
	go func(inp <-chan io.ByteReader, done chan<- []io.ByteReader) {
		defer close(done)
		ByteReaderS := []io.ByteReader{}
		for i := range inp {
			ByteReaderS = append(ByteReaderS, i)
		}
		done <- ByteReaderS
	}(inp, cha)
	return cha
}

// DoneByteReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ByteReader) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.ByteReader, act func(a io.ByteReader)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeByteReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteReaderBuffer(inp <-chan io.ByteReader, cap int) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader, cap)
	go func(out chan<- io.ByteReader, inp <-chan io.ByteReader) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeByteReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader) io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	if act == nil {
		act = func(a io.ByteReader) io.ByteReader { return a }
	}
	go func(out chan<- io.ByteReader, inp <-chan io.ByteReader, act func(a io.ByteReader) io.ByteReader) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeByteReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteReaderFork(inp <-chan io.ByteReader) (out1, out2 <-chan io.ByteReader) {
	cha1 := make(chan io.ByteReader)
	cha2 := make(chan io.ByteReader)
	go func(out1, out2 chan<- io.ByteReader, inp <-chan io.ByteReader) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ByteReaderTube is the signature for a pipe function.
type ByteReaderTube func(inp <-chan io.ByteReader, out <-chan io.ByteReader)

// ByteReaderdaisy returns a channel to receive all inp after having passed thru tube.
func ByteReaderdaisy(inp <-chan io.ByteReader, tube ByteReaderTube) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go tube(inp, cha)
	return cha
}

// ByteReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ByteReaderDaisyChain(inp <-chan io.ByteReader, tubes ...ByteReaderTube) (out <-chan io.ByteReader) {
	cha := inp
	for _, tube := range tubes {
		cha = ByteReaderdaisy(cha, tube)
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
