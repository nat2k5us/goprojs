package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var i int

	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
	go func() {
		i++
		for {
			fmt.Printf("i = %v ", i)
		}
	}()
}

////////////////// OUTPUT ////////////////////
//hello