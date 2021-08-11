package main

import (
	"fmt"
	"time"
)

func waitAndSend(v, i int) chan int {
	channelReturn := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		channelReturn <- v
	}()
	return channelReturn
}

type testinterface interface {
	Say()
	Do()
}

type testSpeechImpl struct {}


func main() {
	select {
	case v1 := <-waitAndSend(3, 2):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1):
		fmt.Println(v2)
	}

}

////////////////// OUTPUT ////////////////////
