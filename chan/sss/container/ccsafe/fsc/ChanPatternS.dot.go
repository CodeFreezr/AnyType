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
func MakePatternSChan() (out chan fs.PatternS) {
	return make(chan fs.PatternS)
}

// ChanPatternS returns a channel to receive all inputs before close.
func ChanPatternS(inp ...fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go func(out chan<- fs.PatternS, inp ...fs.PatternS) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanPatternSSlice returns a channel to receive all inputs before close.
func ChanPatternSSlice(inp ...[]fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go func(out chan<- fs.PatternS, inp ...[]fs.PatternS) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinPatternS
func JoinPatternS(out chan<- fs.PatternS, inp ...fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.PatternS, inp ...fs.PatternS) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinPatternSSlice
func JoinPatternSSlice(out chan<- fs.PatternS, inp ...[]fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.PatternS, inp ...[]fs.PatternS) {
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

// JoinPatternSChan
func JoinPatternSChan(out chan<- fs.PatternS, inp <-chan fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.PatternS, inp <-chan fs.PatternS) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DonePatternS returns a channel to receive one signal before close after inp has been drained.
func DonePatternS(inp <-chan fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan fs.PatternS) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DonePatternSSlice returns a channel which will receive a slice
// of all the PatternSs received on inp channel before close.
// Unlike DonePatternS, a full slice is sent once, not just an event.
func DonePatternSSlice(inp <-chan fs.PatternS) (done <-chan []fs.PatternS) {
	cha := make(chan []fs.PatternS)
	go func(inp <-chan fs.PatternS, done chan<- []fs.PatternS) {
		defer close(done)
		PatternSS := []fs.PatternS{}
		for i := range inp {
			PatternSS = append(PatternSS, i)
		}
		done <- PatternSS
	}(inp, cha)
	return cha
}

// DonePatternSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.PatternS) { return }
	}
	go func(done chan<- struct{}, inp <-chan fs.PatternS, act func(a fs.PatternS)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipePatternSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternSBuffer(inp <-chan fs.PatternS, cap int) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS, cap)
	go func(out chan<- fs.PatternS, inp <-chan fs.PatternS) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipePatternSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	if act == nil {
		act = func(a fs.PatternS) fs.PatternS { return a }
	}
	go func(out chan<- fs.PatternS, inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipePatternSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternSFork(inp <-chan fs.PatternS) (out1, out2 <-chan fs.PatternS) {
	cha1 := make(chan fs.PatternS)
	cha2 := make(chan fs.PatternS)
	go func(out1, out2 chan<- fs.PatternS, inp <-chan fs.PatternS) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// PatternSTube is the signature for a pipe function.
type PatternSTube func(inp <-chan fs.PatternS, out <-chan fs.PatternS)

// PatternSdaisy returns a channel to receive all inp after having passed thru tube.
func PatternSdaisy(inp <-chan fs.PatternS, tube PatternSTube) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go tube(inp, cha)
	return cha
}

// PatternSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PatternSDaisyChain(inp <-chan fs.PatternS, tubes ...PatternSTube) (out <-chan fs.PatternS) {
	cha := inp
	for _, tube := range tubes {
		cha = PatternSdaisy(cha, tube)
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