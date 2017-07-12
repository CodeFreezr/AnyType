// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package pat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeChan returns a new open channel
// (simply a 'chan struct{}' that is).
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
func MakeChan() (out chan struct{}) {
	return make(chan struct{})
}

func send(out chan<- struct{}, inp ...struct{}) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...struct{}) (out <-chan struct{}) {
	cha := make(chan struct{})
	go send(cha, inp...)
	return cha
}

func sendSlice(out chan<- struct{}, inp ...[]struct{}) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]struct{}) (out <-chan struct{}) {
	cha := make(chan struct{})
	go sendSlice(cha, inp...)
	return cha
}

func join(done chan<- struct{}, out chan<- struct{}, inp ...struct{}) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// Join
func Join(out chan<- struct{}, inp ...struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go join(cha, out, inp...)
	return cha
}

func joinSlice(done chan<- struct{}, out chan<- struct{}, inp ...[]struct{}) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinSlice
func JoinSlice(out chan<- struct{}, inp ...[]struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSlice(cha, out, inp...)
	return cha
}

func joinChan(done chan<- struct{}, out chan<- struct{}, inp <-chan struct{}) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinChan
func JoinChan(out chan<- struct{}, inp <-chan struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinChan(cha, out, inp)
	return cha
}

func doit(done chan<- struct{}, inp <-chan struct{}) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doit(cha, inp)
	return cha
}

func doitSlice(done chan<- ([]struct{}), inp <-chan struct{}) {
	defer close(done)
	S := []struct{}{}
	for i := range inp {
		S = append(S, i)
	}
	done <- S
}

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan struct{}) (done <-chan ([]struct{})) {
	cha := make(chan ([]struct{}))
	go doitSlice(cha, inp)
	return cha
}

func doitFunc(done chan<- struct{}, inp <-chan struct{}, act func(a struct{})) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan struct{}, act func(a struct{})) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a struct{}) { return }
	}
	go doitFunc(cha, inp, act)
	return cha
}

func pipeBuffer(out chan<- struct{}, inp <-chan struct{}) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan struct{}, cap int) (out <-chan struct{}) {
	cha := make(chan struct{}, cap)
	go pipeBuffer(cha, inp)
	return cha
}

func pipeFunc(out chan<- struct{}, inp <-chan struct{}, act func(a struct{}) struct{}) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan struct{}, act func(a struct{}) struct{}) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a struct{}) struct{} { return a }
	}
	go pipeFunc(cha, inp, act)
	return cha
}

func pipeFork(out1, out2 chan<- struct{}, inp <-chan struct{}) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan struct{}) (out1, out2 <-chan struct{}) {
	cha1 := make(chan struct{})
	cha2 := make(chan struct{})
	go pipeFork(cha1, cha2, inp)
	return cha1, cha2
}