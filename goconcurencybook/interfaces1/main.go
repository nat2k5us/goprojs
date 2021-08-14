package main

import "fmt"

// Create the interface
type testinterface interface {
	SayHello()
	Say(s string)
	Increment()
	Decrement()
	GetInternalValue()
	Concat()
}

// implementer of the interface - 1
type testinterfaceImplInt struct {
	i int
}

// implementer the interface 2
type testinterfaceImplString struct {
	s string
}

func (t *testinterfaceImplInt) SayHello() {
	fmt.Println("Hello")
}

func (t *testinterfaceImplInt) Concat(a string) {
	fmt.Println("Concat")
}

func (t *testinterfaceImplString) Concat(a string) {
	t.s + a
}
func (t testinterfaceImplInt) Say(s string) {
	fmt.Printf("Hello %s\n", s)
}

// Important to use pointer - else the values are
// local variables scoped and not propagated
func (t *testinterfaceImplInt) Increment() {
	t.i++
}

// Important to use pointer - else the values are
// local variables scoped and not propagated
func (t *testinterfaceImplInt) GetInternalValue() {
	fmt.Printf("i = %d\n", t.i)
}

// Important to use pointer - else the values are
// local variables scoped and not propagated
func (t *testinterfaceImplInt) Decrement() {
	t.i--
}

func main() {
	var myTestInterface testinterface
	myTestInterface = &testinterfaceImplInt{}
	myTestInterface.Say("Natraj")
	myTestInterface.SayHello()

	myTestInterface.Increment()
	myTestInterface.Increment()
	myTestInterface.GetInternalValue()
	myTestInterface.Decrement()
	myTestInterface.Decrement()
	myTestInterface.GetInternalValue()

	var myTestInterfaceString testinterfaceImplString
	myTestInterfaceString = testinterfaceImplString{}
	myTestInterfaceString.Concat("hello")

}
