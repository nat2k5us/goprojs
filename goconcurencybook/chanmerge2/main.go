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

type geometry interface {
	area() float64
}

type square struct {
	side float64
}

func (s *square) area(si float64) float64 {
	s.side = si
	return s.side * s.side
}

type rectangle struct {
	width, height float64
}

func (r *rectangle) area(wi, he float64) float64 {
	r.width = wi
	r.height = he
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c *circle) area(r float64) float64 {
	c.radius = r
	return 3.14 * c.radius * c.radius
}

type MyStream struct {
	Square    float64 `json:"square,omitempty"`
	Rectangle float64 `json:"rectangle,omitempty"`
	Circle    float64 `json:"circle,omitempty"`
}

func UpdateStream(geo <-chan interface{}, stm *MyStream) <-chan interface{} {
	stream := make(chan interface{})
	i1 := square{}
	i2 := rectangle{}
	i3 := circle{}
	
	go func() {
		for {
			select {
			case o1 := <-geo:
				switch o1.(type) {
				case *square:
					stm.Square = i1.area((rand.Float64() * 5) + 25)
					stream <- stm
					return
				case *rectangle:
					stm.Rectangle = i2.area((rand.Float64()*5)+25, (rand.Float64()*5)+25)
					stream <- stm
					return
				case *circle:
					stm.Circle = i3.area((rand.Float64() * 5) + 10)
					stream <- stm
					return
				}
			}
		}
	}()
	return stream
}

func main() {

	stream := make(chan interface{})
	defer close(stream)
	geo := make(chan interface{})
	defer close(geo)

	go func() {
		for {
			time.Sleep(time.Second)
			obj := rand.Intn(3)
			switch obj {
			case 0:
				geo <- new(square)
			case 1:
				geo <- new(rectangle)
			case 2:
				geo <- new(circle)
			}
		}
	}()
	myStreamInstance := MyStream{Square:0, Rectangle:0, Circle:0}
	for {
		got := <-UpdateStream(geo, &myStreamInstance)
		fmt.Printf("got %#v\n", got.(MyStream))
	}
}
