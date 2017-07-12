// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeComplex128Chan returns a new open channel
// (simply a 'chan complex128' that is).
//
// Note: No 'Complex128-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myComplex128PipelineStartsHere := MakeComplex128Chan()
//	// ... lot's of code to design and build Your favourite "myComplex128WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myComplex128PipelineStartsHere <- drop
//	}
//	close(myComplex128PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeComplex128Buffer) the channel is unbuffered.
//
func MakeComplex128Chan() (out chan complex128) {
	return make(chan complex128)
}

func sendComplex128(out chan<- complex128, inp ...complex128) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanComplex128 returns a channel to receive all inputs before close.
func ChanComplex128(inp ...complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	go sendComplex128(cha, inp...)
	return cha
}

func sendComplex128Slice(out chan<- complex128, inp ...[]complex128) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanComplex128Slice returns a channel to receive all inputs before close.
func ChanComplex128Slice(inp ...[]complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	go sendComplex128Slice(cha, inp...)
	return cha
}

func joinComplex128(done chan<- struct{}, out chan<- complex128, inp ...complex128) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinComplex128
func JoinComplex128(out chan<- complex128, inp ...complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex128(cha, out, inp...)
	return cha
}

func joinComplex128Slice(done chan<- struct{}, out chan<- complex128, inp ...[]complex128) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinComplex128Slice
func JoinComplex128Slice(out chan<- complex128, inp ...[]complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex128Slice(cha, out, inp...)
	return cha
}

func joinComplex128Chan(done chan<- struct{}, out chan<- complex128, inp <-chan complex128) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinComplex128Chan
func JoinComplex128Chan(out chan<- complex128, inp <-chan complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex128Chan(cha, out, inp)
	return cha
}

func doitComplex128(done chan<- struct{}, inp <-chan complex128) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneComplex128 returns a channel to receive one signal before close after inp has been drained.
func DoneComplex128(inp <-chan complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitComplex128(cha, inp)
	return cha
}

func doitComplex128Slice(done chan<- ([]complex128), inp <-chan complex128) {
	defer close(done)
	Complex128S := []complex128{}
	for i := range inp {
		Complex128S = append(Complex128S, i)
	}
	done <- Complex128S
}

// DoneComplex128Slice returns a channel which will receive a slice
// of all the Complex128s received on inp channel before close.
// Unlike DoneComplex128, a full slice is sent once, not just an event.
func DoneComplex128Slice(inp <-chan complex128) (done <-chan ([]complex128)) {
	cha := make(chan ([]complex128))
	go doitComplex128Slice(cha, inp)
	return cha
}

func doitComplex128Func(done chan<- struct{}, inp <-chan complex128, act func(a complex128)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneComplex128Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneComplex128Func(inp <-chan complex128, act func(a complex128)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a complex128) { return }
	}
	go doitComplex128Func(cha, inp, act)
	return cha
}

func pipeComplex128Buffer(out chan<- complex128, inp <-chan complex128) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeComplex128Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeComplex128Buffer(inp <-chan complex128, cap int) (out <-chan complex128) {
	cha := make(chan complex128, cap)
	go pipeComplex128Buffer(cha, inp)
	return cha
}

func pipeComplex128Func(out chan<- complex128, inp <-chan complex128, act func(a complex128) complex128) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeComplex128Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeComplex128Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeComplex128Func(inp <-chan complex128, act func(a complex128) complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	if act == nil {
		act = func(a complex128) complex128 { return a }
	}
	go pipeComplex128Func(cha, inp, act)
	return cha
}

func pipeComplex128Fork(out1, out2 chan<- complex128, inp <-chan complex128) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeComplex128Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeComplex128Fork(inp <-chan complex128) (out1, out2 <-chan complex128) {
	cha1 := make(chan complex128)
	cha2 := make(chan complex128)
	go pipeComplex128Fork(cha1, cha2, inp)
	return cha1, cha2
}