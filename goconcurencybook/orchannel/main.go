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
	// The Or-Channel
	// At times you may find yourself wanting to combine one or more done
	// channels into a single done channel that closes if any of its component
	// channels close. It is perfectly acceptable, albeit verbose, to write a
	// select statement that performs this coupling; however, sometimes you
	// can’t know the number of done channels you’re working with at runtime.
	// In this case, or if you just prefer a one-liner, you can combine these
	// channels together using the or-channel pattern.
	var or func(channels ...<-chan interface{}) <-chan interface{}
	// Here we have our function, 'or', which takes in a variadic
	// slice of channels and returns a single channel.
	or = func(channels ...<-chan interface{}) <-chan interface{} {

		switch len(channels) {
			// Since this is a recursive function, we must set up termination 
			// criteria. The first is that if the variadic slice is empty, we simply
			//  return a nil channel. This is consistant with the idea of passing
			//   in no channels; we wouldn’t expect a composite channel to do anything.
		case 0:
			return nil
			// Our second termination criteria states that if our variadic slice only 
			// contains one element, we just return that element.
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})
		// Here is the main body of the function, and where the recursion happens. 
		// We create a goroutine so that we can wait for messages on our channels without blocking.
		go func() {
			defer close(orDone)
			switch len(channels) {
				// Because of how we’re recursing, every recursive call to "or" will at least have 
				// two channels. As an optimization to keep the number of goroutines constrained, 
				// we place a special case here for calls to or with only two channels.
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
				// Here we recursively create an or-channel from all the channels in our 
				// slice after the third index, and then select from this. This recurrence 
				// relation will de-structure the rest of the slice into or-channels to form
				// a tree from which the first signal will return. We also pass in the 
				// orDone channel so that when goroutines up the tree exit, goroutines down
				// the tree also exit.
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

}
