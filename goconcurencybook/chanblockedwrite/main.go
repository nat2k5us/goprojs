package main

import (
	"fmt"
	"math/rand"
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
// 		 do some work on channels
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
	// //memory leaked version
	// newRandStream := func() <-chan int {
	// 	randStream := make(chan int)
	// 	go func() {
	// 		defer fmt.Println("newRandStream closure exited.")
	// 		defer close(randStream)
	// 		for {
	// 			randStream <- rand.Int()
	// 		}
	// 	}()
	// 	return randStream
	// }
	// randStream := newRandStream()
	// fmt.Println("3 random ints:")
	// for i := 1; i <= 3; i++ {
	// 	fmt.Printf("%d: %d\n", i, <-randStream)
	// }
// /// Output //////////////////////////////////
// 3 random ints:
// 1: 5577006791947779410
// 2: 8674665223082153551
// 3: 6129484611666145821

	// Fixed Version
	// The 'done' channel is used to prevent memory leaks
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}
	done := make(chan interface{}) // done channel
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)

	// /// Output //////////////////////////////////
	// 3 random ints:
	// 1: 5577006791947779410
	// 2: 8674665223082153551
	// 3: 6129484611666145821
	// newRandStream closure exited.
}
