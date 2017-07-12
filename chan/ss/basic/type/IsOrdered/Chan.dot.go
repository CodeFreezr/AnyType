// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeChan returns a new open channel
// (simply a 'chan string' that is).
//
// Note: No '-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipelineStartsHere := MakeChan()
//	// ... lot's of code to design and build Your favourite "myWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipelineStartsHere <- drop
//	}
//	close(myPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBuffer) the channel is unbuffered.
//
func MakeChan() (out chan string) {
	return make(chan string)
}

func send(out chan<- string, inp ...string) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...string) (out <-chan string) {
	cha := make(chan string)
	go send(cha, inp...)
	return cha
}

func sendSlice(out chan<- string, inp ...[]string) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]string) (out <-chan string) {
	cha := make(chan string)
	go sendSlice(cha, inp...)
	return cha
}

func join(done chan<- struct{}, out chan<- string, inp ...string) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// Join
func Join(out chan<- string, inp ...string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go join(cha, out, inp...)
	return cha
}

func joinSlice(done chan<- struct{}, out chan<- string, inp ...[]string) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinSlice
func JoinSlice(out chan<- string, inp ...[]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSlice(cha, out, inp...)
	return cha
}

func joinChan(done chan<- struct{}, out chan<- string, inp <-chan string) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinChan
func JoinChan(out chan<- string, inp <-chan string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinChan(cha, out, inp)
	return cha
}

func doit(done chan<- struct{}, inp <-chan string) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doit(cha, inp)
	return cha
}

func doitSlice(done chan<- ([]string), inp <-chan string) {
	defer close(done)
	S := []string{}
	for i := range inp {
		S = append(S, i)
	}
	done <- S
}

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan string) (done <-chan ([]string)) {
	cha := make(chan ([]string))
	go doitSlice(cha, inp)
	return cha
}

func doitFunc(done chan<- struct{}, inp <-chan string, act func(a string)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan string, act func(a string)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a string) { return }
	}
	go doitFunc(cha, inp, act)
	return cha
}

func pipeBuffer(out chan<- string, inp <-chan string) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan string, cap int) (out <-chan string) {
	cha := make(chan string, cap)
	go pipeBuffer(cha, inp)
	return cha
}

func pipeFunc(out chan<- string, inp <-chan string, act func(a string) string) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan string, act func(a string) string) (out <-chan string) {
	cha := make(chan string)
	if act == nil {
		act = func(a string) string { return a }
	}
	go pipeFunc(cha, inp, act)
	return cha
}

func pipeFork(out1, out2 chan<- string, inp <-chan string) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan string) (out1, out2 <-chan string) {
	cha1 := make(chan string)
	cha2 := make(chan string)
	go pipeFork(cha1, cha2, inp)
	return cha1, cha2
}

// Merge2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func Merge2(i1, i2 <-chan string) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, i1, i2 <-chan string) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 string // what we've read
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