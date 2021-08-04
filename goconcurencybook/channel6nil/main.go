package main

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
	// read from a nil data stream
	var datastream chan interface{}
	<-datastream
	// ----- OUTPUT --------------------------------
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive (nil chan)]:

	// write to a nil data stream
	var dataStream chan interface{}
	dataStream <- struct{}{}
	// ----- OUTPUT --------------------------------
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send (nil chan)]:

	// close a nil channel
	var dataStream chan interface{}
	close(dataStream)
	// ----- OUTPUT -----------------------------
	// panic: close of nil channel

}


