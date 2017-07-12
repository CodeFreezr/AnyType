// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt16Chan returns a new open channel
// (simply a 'chan int16' that is).
//
// Note: No 'Int16-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt16PipelineStartsHere := MakeInt16Chan()
//	// ... lot's of code to design and build Your favourite "myInt16WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt16PipelineStartsHere <- drop
//	}
//	close(myInt16PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt16Buffer) the channel is unbuffered.
//
func MakeInt16Chan() (out chan int16) {
	return make(chan int16)
}

// ChanInt16 returns a channel to receive all inputs before close.
func ChanInt16(inp ...int16) (out <-chan int16) {
	cha := make(chan int16)
	go func(out chan<- int16, inp ...int16) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanInt16Slice returns a channel to receive all inputs before close.
func ChanInt16Slice(inp ...[]int16) (out <-chan int16) {
	cha := make(chan int16)
	go func(out chan<- int16, inp ...[]int16) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinInt16
func JoinInt16(out chan<- int16, inp ...int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int16, inp ...int16) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinInt16Slice
func JoinInt16Slice(out chan<- int16, inp ...[]int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int16, inp ...[]int16) {
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

// JoinInt16Chan
func JoinInt16Chan(out chan<- int16, inp <-chan int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int16, inp <-chan int16) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneInt16 returns a channel to receive one signal before close after inp has been drained.
func DoneInt16(inp <-chan int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan int16) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneInt16Slice returns a channel which will receive a slice
// of all the Int16s received on inp channel before close.
// Unlike DoneInt16, a full slice is sent once, not just an event.
func DoneInt16Slice(inp <-chan int16) (done <-chan []int16) {
	cha := make(chan []int16)
	go func(inp <-chan int16, done chan<- []int16) {
		defer close(done)
		Int16S := []int16{}
		for i := range inp {
			Int16S = append(Int16S, i)
		}
		done <- Int16S
	}(inp, cha)
	return cha
}

// DoneInt16Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt16Func(inp <-chan int16, act func(a int16)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int16) { return }
	}
	go func(done chan<- struct{}, inp <-chan int16, act func(a int16)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeInt16Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt16Buffer(inp <-chan int16, cap int) (out <-chan int16) {
	cha := make(chan int16, cap)
	go func(out chan<- int16, inp <-chan int16) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeInt16Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt16Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt16Func(inp <-chan int16, act func(a int16) int16) (out <-chan int16) {
	cha := make(chan int16)
	if act == nil {
		act = func(a int16) int16 { return a }
	}
	go func(out chan<- int16, inp <-chan int16, act func(a int16) int16) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeInt16Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt16Fork(inp <-chan int16) (out1, out2 <-chan int16) {
	cha1 := make(chan int16)
	cha2 := make(chan int16)
	go func(out1, out2 chan<- int16, inp <-chan int16) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// Int16Tube is the signature for a pipe function.
type Int16Tube func(inp <-chan int16, out <-chan int16)

// Int16daisy returns a channel to receive all inp after having passed thru tube.
func Int16daisy(inp <-chan int16, tube Int16Tube) (out <-chan int16) {
	cha := make(chan int16)
	go tube(inp, cha)
	return cha
}

// Int16DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Int16DaisyChain(inp <-chan int16, tubes ...Int16Tube) (out <-chan int16) {
	cha := inp
	for _, tube := range tubes {
		cha = Int16daisy(cha, tube)
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
// MergeInt16 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int16channel is returned.
func MergeInt16(inps ...<-chan int16) (out <-chan int16) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int16)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return merge2(inps[0], Merge(inps[1:]...))
	}
}

// mergeInt162 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt162(i1, i2 <-chan int16) (out <-chan int16) {
	cha := make(chan int16)
	go func(out chan<- int16, i1, i2 <-chan int16) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int16 // what we've read
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