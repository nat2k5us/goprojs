package main

import (
	"fmt"
)

type human interface {
	sayHello() string
}

type man struct {
	Greeting string
}

type woman struct {
	Greeting string
}

func (m man) sayHello() string {
	return m.Greeting
}

func (w woman) sayHello() string {
	return w.Greeting
}

type Greeting interface {
	Write([]byte) (int, error)
}

func printGreeting(h human) {
	fmt.Println(h.sayHello())
}

func main() {
	fmt.Println("hello interfaces")
	fooMan := man{Greeting: "whats up dude!!"}
	barWoman := woman{Greeting: "Hi *** Hugs"}

	printGreeting(fooMan)
	printGreeting(barWoman)
}
