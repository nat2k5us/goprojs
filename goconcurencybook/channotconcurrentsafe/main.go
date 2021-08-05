package main

import (
	"bytes"
	"fmt"
	"sync"
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

func main() {

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}
