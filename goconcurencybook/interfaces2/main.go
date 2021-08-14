package main

import (
	"fmt"
	"math"
)

// abstract
type geometry interface {
	area() float64
	perimeter() float64
}

// concrete 1
type rect struct {
	width, height float64
}

// constructor function in golang
// func "New""ConcImplName"() "interface" as returned by
// for rect for example
func NewRectConstructor() geometry {
	return &rect{width: 1, height: 1}
}
// constructor 2
// note constructor with same names are disallowed
func NewRectConstructorWithValues(w float64, h float64) geometry {
	return &rect{width: w, height: h}
}

// concrete 2
type circle struct {
	radius float64
}

// embedded types
// is when we create a new type with all the functionality of its parent type
type ellipse struct {
	*circle // ellipse will get all the methods and properties of circle
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}



func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	eli := ellipse{circle: &circle{radius: 15}}
	measure(eli)
}
