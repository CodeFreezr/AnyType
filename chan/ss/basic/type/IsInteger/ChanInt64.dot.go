// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt64Chan returns a new open channel
// (simply a 'chan int64' that is).
//
// Note: No 'Int64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt64PipelineStartsHere := MakeInt64Chan()
//	// ... lot's of code to design and build Your favourite "myInt64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt64PipelineStartsHere <- drop
//	}
//	close(myInt64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt64Buffer) the channel is unbuffered.
//
func MakeInt64Chan() (out chan int64) {
	return make(chan int64)
}

func sendInt64(out chan<- int64, inp ...int64) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanInt64 returns a channel to receive all inputs before close.
func ChanInt64(inp ...int64) (out <-chan int64) {
	cha := make(chan int64)
	go sendInt64(cha, inp...)
	return cha
}

func sendInt64Slice(out chan<- int64, inp ...[]int64) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanInt64Slice returns a channel to receive all inputs before close.
func ChanInt64Slice(inp ...[]int64) (out <-chan int64) {
	cha := make(chan int64)
	go sendInt64Slice(cha, inp...)
	return cha
}

func joinInt64(done chan<- struct{}, out chan<- int64, inp ...int64) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64(out chan<- int64, inp ...int64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt64(cha, out, inp...)
	return cha
}

func joinInt64Slice(done chan<- struct{}, out chan<- int64, inp ...[]int64) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinInt64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64Slice(out chan<- int64, inp ...[]int64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt64Slice(cha, out, inp...)
	return cha
}

func joinInt64Chan(done chan<- struct{}, out chan<- int64, inp <-chan int64) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64Chan(out chan<- int64, inp <-chan int64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt64Chan(cha, out, inp)
	return cha
}

func doitInt64(done chan<- struct{}, inp <-chan int64) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneInt64 returns a channel to receive one signal before close after inp has been drained.
func DoneInt64(inp <-chan int64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitInt64(cha, inp)
	return cha
}

func doitInt64Slice(done chan<- ([]int64), inp <-chan int64) {
	defer close(done)
	Int64S := []int64{}
	for i := range inp {
		Int64S = append(Int64S, i)
	}
	done <- Int64S
}

// DoneInt64Slice returns a channel which will receive a slice
// of all the Int64s received on inp channel before close.
// Unlike DoneInt64, a full slice is sent once, not just an event.
func DoneInt64Slice(inp <-chan int64) (done <-chan ([]int64)) {
	cha := make(chan ([]int64))
	go doitInt64Slice(cha, inp)
	return cha
}

func doitInt64Func(done chan<- struct{}, inp <-chan int64, act func(a int64)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneInt64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt64Func(inp <-chan int64, act func(a int64)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int64) { return }
	}
	go doitInt64Func(cha, inp, act)
	return cha
}

func pipeInt64Buffer(out chan<- int64, inp <-chan int64) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeInt64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt64Buffer(inp <-chan int64, cap int) (out <-chan int64) {
	cha := make(chan int64, cap)
	go pipeInt64Buffer(cha, inp)
	return cha
}

func pipeInt64Func(out chan<- int64, inp <-chan int64, act func(a int64) int64) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeInt64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt64Func(inp <-chan int64, act func(a int64) int64) (out <-chan int64) {
	cha := make(chan int64)
	if act == nil {
		act = func(a int64) int64 { return a }
	}
	go pipeInt64Func(cha, inp, act)
	return cha
}

func pipeInt64Fork(out1, out2 chan<- int64, inp <-chan int64) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeInt64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt64Fork(inp <-chan int64) (out1, out2 <-chan int64) {
	cha1 := make(chan int64)
	cha2 := make(chan int64)
	go pipeInt64Fork(cha1, cha2, inp)
	return cha1, cha2
}

// MergeInt64 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int64channel is returned.
func MergeInt64(inps ...<-chan int64) (out <-chan int64) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int64)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt642(inps[0], MergeInt64(inps[1:]...))
	}
}

// mergeInt642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt642(i1, i2 <-chan int64) (out <-chan int64) {
	cha := make(chan int64)
	go func(out chan<- int64, i1, i2 <-chan int64) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int64 // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
