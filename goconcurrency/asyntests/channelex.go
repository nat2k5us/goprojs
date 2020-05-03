package asyntests

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("executing go routine")
	c <- "Hello Channel" // see the ref kinda operation - value of channel var is set
	time.Sleep(1 * time.Second)
	fmt.Println("finished exuting go routine ..")
}

func TestParrallProc() {
	// <- is the channel operator
	// you make a channel using make (chan Type)
	c := make(chan string, 2) // Make a channel
	defer close(c)            // prevents interaction

	go SendValue(c)
	go SendValue(c)
	go SendValue(c)
	value := <-c
	fmt.Println(value)

	time.Sleep(2 * time.Second)
}



func asyncGo() {
	fmt.Println("This will happen first")

	go func() {
		fmt.Println("This will happen at some unknown time")
	}()

	fmt.Println("This will either happen second or third")

	fmt.Scanln()
	fmt.Println("done")
}

func asyncGo2() {
	a := 1
	b := 2

	go func() { // B2
		b = a * b
	}()

	a = b * b // B1

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()
	// prints a= 4 and b = 8 - as B1 executes before B2
	// this is what is called a race condition
	fmt.Printf("a = %d, b = %d\n", a, b)
}