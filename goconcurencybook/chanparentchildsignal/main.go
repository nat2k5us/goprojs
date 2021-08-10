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
	// Here we pass the done channel to the doWork function. 
	doWork := func(done <-chan interface{},
		strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done:
					// On this line we see the ubiquitous for-select pattern in use. 
					// One of our case state‐ ments is checking whether our done channel
					//  has been signaled. If it has, we return from the goroutine.
					fmt.Println("channel closed - done")
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	// Here we create another goroutine that will cancel the 
	// goroutine spawned in doWork if more than one second passes.
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	// This is where we join the goroutine spawned from doWork 
	// with the main goroutine.
	<-terminated 
	fmt.Println("Done.")

	// You can see that despite passing in nil for our strings 
	// channel, our goroutine still exits successfully. Unlike the 
	// example before it, in this example we do join the two goroutines,
	//  and yet do not receive a deadlock. This is because before we 
	//  join the two goroutines, we create a third goroutine to cancel
	//   the goroutine within doWork after a second. We have 
	//   successfully eliminated our goroutine leak!
}
