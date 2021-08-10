package main

import (
	"fmt"
	"time"
)

// Like a river, a channel serves as a conduit for a stream of information;
// values may be passed along the channel, and then read out downstream.
// For this reason I usually end my chan variable names with the word “Stream.”
// When using channels, you’ll pass a value into a chan variable, and then
// somewhere else in your program read it off the channel. The disparate
// parts of your program don’t require knowledge of each other, only a
// reference to the same place in memory where the channel resides.
// This can be done by passing references of channels around your program.

// Read and Write Channel example
// var dataStream chan interface{}
// 	dataStream = make(chan interface{})

// Read or Receive only Channel example
// var dataStream <- chan interface{}
// dataStream := make( <-chan interface{})

// Send or Write only Channel example
// var dataStream chan <- interface{}
// 	dataStream := make(chan <- interface{})

// Notes: Once a channel is read - it will be empty

// The for-select loop
// for { // either loop infinitely or range over something
// 	select {
// 		// do some work on channels
// 	}
// }

// Send iteration variables to channel
// stringStream := make(chan interface{})
// 	defer close(stringStream)
// 	for _, s := range []string{"a", "b", "c"} {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case stringStream <- s:
// 		}
// 	}
func main() {

	done := make(chan interface{})
	defer close(done)
	go func() {
		for _, c := range []int{1, 2, 3, 5, 4} {
			done <- c
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case <-time.After(time.Second * 5):
			fmt.Println("5 seconds passed")
		case val := <-done:
			fmt.Printf("%v", val)
		default:
			fmt.Print("U")

		}
		// Do non-preemptable work
		// fmt.Println("non-preemptive work")
	}

}
