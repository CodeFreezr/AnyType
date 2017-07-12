// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// MakeWriterChan returns a new open channel
// (simply a 'chan *bufio.Writer' that is).
//
// Note: No 'Writer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterPipelineStartsHere := MakeWriterChan()
//	// ... lot's of code to design and build Your favourite "myWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterPipelineStartsHere <- drop
//	}
//	close(myWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterBuffer) the channel is unbuffered.
//
func MakeWriterChan() (out chan *bufio.Writer) {
	return make(chan *bufio.Writer)
}

// ChanWriter returns a channel to receive all inputs before close.
func ChanWriter(inp ...*bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go func(out chan<- *bufio.Writer, inp ...*bufio.Writer) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanWriterSlice returns a channel to receive all inputs before close.
func ChanWriterSlice(inp ...[]*bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go func(out chan<- *bufio.Writer, inp ...[]*bufio.Writer) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinWriter
func JoinWriter(out chan<- *bufio.Writer, inp ...*bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Writer, inp ...*bufio.Writer) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinWriterSlice
func JoinWriterSlice(out chan<- *bufio.Writer, inp ...[]*bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Writer, inp ...[]*bufio.Writer) {
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

// JoinWriterChan
func JoinWriterChan(out chan<- *bufio.Writer, inp <-chan *bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneWriter returns a channel to receive one signal before close after inp has been drained.
func DoneWriter(inp <-chan *bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *bufio.Writer) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneWriterSlice returns a channel which will receive a slice
// of all the Writers received on inp channel before close.
// Unlike DoneWriter, a full slice is sent once, not just an event.
func DoneWriterSlice(inp <-chan *bufio.Writer) (done <-chan []*bufio.Writer) {
	cha := make(chan []*bufio.Writer)
	go func(inp <-chan *bufio.Writer, done chan<- []*bufio.Writer) {
		defer close(done)
		WriterS := []*bufio.Writer{}
		for i := range inp {
			WriterS = append(WriterS, i)
		}
		done <- WriterS
	}(inp, cha)
	return cha
}

// DoneWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterFunc(inp <-chan *bufio.Writer, act func(a *bufio.Writer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.Writer) { return }
	}
	go func(done chan<- struct{}, inp <-chan *bufio.Writer, act func(a *bufio.Writer)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterBuffer(inp <-chan *bufio.Writer, cap int) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer, cap)
	go func(out chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterFunc(inp <-chan *bufio.Writer, act func(a *bufio.Writer) *bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	if act == nil {
		act = func(a *bufio.Writer) *bufio.Writer { return a }
	}
	go func(out chan<- *bufio.Writer, inp <-chan *bufio.Writer, act func(a *bufio.Writer) *bufio.Writer) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterFork(inp <-chan *bufio.Writer) (out1, out2 <-chan *bufio.Writer) {
	cha1 := make(chan *bufio.Writer)
	cha2 := make(chan *bufio.Writer)
	go func(out1, out2 chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// WriterTube is the signature for a pipe function.
type WriterTube func(inp <-chan *bufio.Writer, out <-chan *bufio.Writer)

// Writerdaisy returns a channel to receive all inp after having passed thru tube.
func Writerdaisy(inp <-chan *bufio.Writer, tube WriterTube) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go tube(inp, cha)
	return cha
}

// WriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterDaisyChain(inp <-chan *bufio.Writer, tubes ...WriterTube) (out <-chan *bufio.Writer) {
	cha := inp
	for _, tube := range tubes {
		cha = Writerdaisy(cha, tube)
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