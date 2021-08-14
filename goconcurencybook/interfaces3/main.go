package main

import (
	"fmt"
)

func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It is a string")
	case int:
		fmt.Println("It is a int")
	case float64:
		fmt.Println("It is a float64")
	default:
		fmt.Println("It is not a known type")
	}
}

func Interfaces() {
	CheckType("string")
	CheckType(5)
	CheckType(0xFFFF)
	CheckType(false)

	var i interface{}
	i = "Super"
	if val, ok := i.(string); ok {
		fmt.Println("val is ", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("handle the failed case of val")
	}
}

func main() {

	Interfaces()
}
