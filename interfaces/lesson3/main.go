package main

import (
	"fmt"
	"math"
)

type square struct {
	width float64
}

type circle struct {
	radius float64
}

type shapeMetrics interface {
	area() float64
	perimeter() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) perimeter() float64 {
	return 4 * s.width
}



func main() {
   squareInstance := square{width: 10}
   circleInstance := circle{radius: 5}

  // fmt.Println(squareInstance.area())
   fmt.Println(squareInstance.perimeter())

   fmt.Println(circleInstance.area())
   fmt.Println(circleInstance.perimeter())
}
