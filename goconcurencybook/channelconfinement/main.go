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

// Unblocking multiple go routines at once
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
	consumer := func(results <-chan int) { // fn input param <-chan int
		for result := range results {
			result = 10
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	results := chanOwner()

	consumer(results) // consume is not allowed to modify channel data

	for x := range results {
		fmt.Println(x)
	}

	fmt.Printf("Received: %#v\n", results)
}
