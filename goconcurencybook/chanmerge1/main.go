package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	// log "github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// merge takes two or more  channels and merges them into a single channel
func merge(cs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan interface{}) {
			for v := range c {
				switch v.(type) {
				case nil:
					fmt.Println("got chan nil")
					out <- errors.New("nil stream closed")
					return
				case error:
					fmt.Println("got chan error")
					out <- errors.New("error stream closed")
					return
				case string:
					out <- v
				case int:
					out <- fmt.Sprintf("%v", v)
				}
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		for {
			time.Sleep(time.Second)
			t1 := rand.Intn(50)
			fmt.Println("Add to 1:", t1)
			ch1 <- t1
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			t2 := RandStringBytes(8)
			fmt.Println("Add to 2:", t2)
			ch2 <- t2
		}
	}()

	for {
		select {
		case o1 := <-ch1:
			// The below merge does not work because
			fmt.Println("Merged", <-merge(ch1, ch2))
			fmt.Printf("\t  1: %v\n", o1)

		case o2 := <-ch2:
			fmt.Println("Merged", <-merge(ch1, ch2))
			fmt.Printf("\t 1: %v 2: %v\n",ch1,  o2)

		}
	}
}
