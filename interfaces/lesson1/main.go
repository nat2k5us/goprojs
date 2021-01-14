package main

import (
	"fmt"
)

type human interface {
	sayHello() string
	sayBye() string
}

type man struct {
	Greeting string
	Bye string
}

type woman struct {
	Greeting string
	Bye string
}

func (m man) sayHello() string {
	return m.Greeting
}

func (m man) sayBye() string {
	return m.Bye
}

func (w woman) sayHello() string {
	return w.Greeting
}
func (w woman) sayBye() string {
	return w.Bye
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
