package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
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
			fmt.Printf("\t  1: %v\n", o1)
		case o2 := <-ch2:
			fmt.Printf("\t  2: %v\n", o2)
		}
	}
}
