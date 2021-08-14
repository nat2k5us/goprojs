package main

import (
	"context"
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

func ContextWithForceCancel(){
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3*time.Second))

	go func(ctx context.Context) {
		// simulate a process that takes 2 second to complete
		time.Sleep(5 * time.Second)

		// cancel context by force, assuming the whole process is complete
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("context timeout exceeded")
		case context.Canceled:
			fmt.Println("context cancelled by force." +
				"whole process is complete")
		}
	}
}
func main() {
	ContextWithForceCancel()
	
}
