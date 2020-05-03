package asyntests

import (
	"fmt"
	"sync"
	"time"
)

var wgGlobal sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wgGlobal.Done() // keep channel open till done - might be decrementing wg index when done
	c <- someValue * 5
}

func WaitGroupTest() {
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wgGlobal.Add(1)
		go foo(fooVal, i)
	}
	wgGlobal.Wait() // blocks till wait group no longer needs to wait
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
}

func WGTest2() {

	var wg sync.WaitGroup
	wg.Add(1) // increment index
	go func() {
		fmt.Println("")
		time.Sleep(time.Second)
		fmt.Println("delay")
		wg.Done() //release block
	}()
	fmt.Println("Waiting。。。")
	wg.Wait() //block
}
