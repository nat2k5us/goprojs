package main

import (
	"fmt"
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
func main() {

	chanOwner := func() <-chan int { // func returns <-chan int
		results := make(chan int, 5) // make a buffered channel of 5
		go func() {                  // go routine
			defer close(results) // close the channel when fn goes out of scope
			for i := 0; i <= 5; i++ {
				results <- i // write to channel
			}
		}()
		return results
	}

	// channel is populated with data
	results := chanOwner()

	// read the data that was populated
	for x := range results {
		fmt.Printf("x: %d\n", x)
	}
	// try to read it again will give empty channel
	for x := range results {
		fmt.Printf("x: %d\n", x)
	}

}
