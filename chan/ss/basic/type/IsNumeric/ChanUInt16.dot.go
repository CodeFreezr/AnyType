// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt16Chan returns a new open channel
// (simply a 'chan uint16' that is).
//
// Note: No 'UInt16-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt16PipelineStartsHere := MakeUInt16Chan()
//	// ... lot's of code to design and build Your favourite "myUInt16WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt16PipelineStartsHere <- drop
//	}
//	close(myUInt16PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt16Buffer) the channel is unbuffered.
//
func MakeUInt16Chan() (out chan uint16) {
	return make(chan uint16)
}

func sendUInt16(out chan<- uint16, inp ...uint16) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanUInt16 returns a channel to receive all inputs before close.
func ChanUInt16(inp ...uint16) (out <-chan uint16) {
	cha := make(chan uint16)
	go sendUInt16(cha, inp...)
	return cha
}

func sendUInt16Slice(out chan<- uint16, inp ...[]uint16) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanUInt16Slice returns a channel to receive all inputs before close.
func ChanUInt16Slice(inp ...[]uint16) (out <-chan uint16) {
	cha := make(chan uint16)
	go sendUInt16Slice(cha, inp...)
	return cha
}

func joinUInt16(done chan<- struct{}, out chan<- uint16, inp ...uint16) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt16 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16(out chan<- uint16, inp ...uint16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt16(cha, out, inp...)
	return cha
}

func joinUInt16Slice(done chan<- struct{}, out chan<- uint16, inp ...[]uint16) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinUInt16Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16Slice(out chan<- uint16, inp ...[]uint16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt16Slice(cha, out, inp...)
	return cha
}

func joinUInt16Chan(done chan<- struct{}, out chan<- uint16, inp <-chan uint16) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt16Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16Chan(out chan<- uint16, inp <-chan uint16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt16Chan(cha, out, inp)
	return cha
}

func doitUInt16(done chan<- struct{}, inp <-chan uint16) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneUInt16 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt16(inp <-chan uint16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitUInt16(cha, inp)
	return cha
}

func doitUInt16Slice(done chan<- ([]uint16), inp <-chan uint16) {
	defer close(done)
	UInt16S := []uint16{}
	for i := range inp {
		UInt16S = append(UInt16S, i)
	}
	done <- UInt16S
}

// DoneUInt16Slice returns a channel which will receive a slice
// of all the UInt16s received on inp channel before close.
// Unlike DoneUInt16, a full slice is sent once, not just an event.
func DoneUInt16Slice(inp <-chan uint16) (done <-chan ([]uint16)) {
	cha := make(chan ([]uint16))
	go doitUInt16Slice(cha, inp)
	return cha
}

func doitUInt16Func(done chan<- struct{}, inp <-chan uint16, act func(a uint16)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneUInt16Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt16Func(inp <-chan uint16, act func(a uint16)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint16) { return }
	}
	go doitUInt16Func(cha, inp, act)
	return cha
}

func pipeUInt16Buffer(out chan<- uint16, inp <-chan uint16) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeUInt16Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt16Buffer(inp <-chan uint16, cap int) (out <-chan uint16) {
	cha := make(chan uint16, cap)
	go pipeUInt16Buffer(cha, inp)
	return cha
}

func pipeUInt16Func(out chan<- uint16, inp <-chan uint16, act func(a uint16) uint16) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeUInt16Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt16Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt16Func(inp <-chan uint16, act func(a uint16) uint16) (out <-chan uint16) {
	cha := make(chan uint16)
	if act == nil {
		act = func(a uint16) uint16 { return a }
	}
	go pipeUInt16Func(cha, inp, act)
	return cha
}

func pipeUInt16Fork(out1, out2 chan<- uint16, inp <-chan uint16) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeUInt16Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt16Fork(inp <-chan uint16) (out1, out2 <-chan uint16) {
	cha1 := make(chan uint16)
	cha2 := make(chan uint16)
	go pipeUInt16Fork(cha1, cha2, inp)
	return cha1, cha2
}
