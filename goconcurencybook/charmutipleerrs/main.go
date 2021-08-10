package main

import (
	"fmt"
	"net/http"
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

// Here we create a type that encompasses both the *http.Response
//and the error possible from an iteration of the loop within our
//goroutine.
type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	// !Error handling
	//In concurrent programs, error handling can be difficult to get right.
	//Sometimes, we spend so much time thinking about how our various processes
	//will be sharing information and coordinating, we forget to consider
	//how they’ll gracefully handle error-ed states. When Go eschewed the
	//popular exception model of errors, it made a statement that error
	// handling was important, and that as we develop our programs, we should
	// give our error paths the same attention we give our algorithms.
	// In that spirit, let’s take a look at how we do that when working with
	// multiple concurrent processes.
	//The most fundamental question when thinking about error handling is,
	//“Who should be responsible for handling the error?” At some point,
	//the program needs to stop ferrying the error up the stack and actually
	// do something with it. What is responsible for this?
	//With concurrent processes, this question becomes a little more
	// complex. Because a concurrent process is operating independently of
	// its parent or siblings, it can be difficult for it to reason about
	// what the right thing to do with the error is. Take a look at the
	// following code for an example of this issue:
	checkStatus := func(done <-chan interface{}, urls ...string,
	) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}

				select {
				case <-done:
					return
				case results <- result: // This is where we write the Result to our channel.
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		// Here, in our main goroutine, we are able to deal with 
		// errors coming out of the goroutine started by checkStatus
		// intelligently, and with the full context of the larger program.
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
