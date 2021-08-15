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

// merge takes two or more  channels and merges them into a single channel
// func GiveMeStream(ctx context.Context, genericInterface <-chan interface{}) <-chan interface{} {
// 	outputCh := make(chan interface{})
// 	res := MyStream{Square: 0, Rectangle: 0, Circle: 0}
// 	i1 := square{}
// 	i2 := rectangle{}
// 	i3 := circle{}
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				log.Debug("Give me stream ctx.Done")
// 				return
// 			case unk := <-genericInterface:
// 				switch asTyp := unk.(type) {
// 				default:
// 					fmt.Println("Testing type")
// 					continue
// 				case nil:
// 					log.Warn("generic interface chan was nil")
// 					outputCh <- errors.New("stream closed")
// 					return
// 				case error:
// 					log.WithError(asTyp).Warn("err from generic interface")
// 					outputCh <- errors.New("error from generic interface channel")
// 					return
// 				case square:
// 					res.Square = i1.area((rand.Float64() * 5) + 25)
// 					outputCh <- res
// 					continue
// 				case rectangle:
// 					res.Rectangle = i2.area((rand.Float64()*5)+25, (rand.Float64()*5)+25)
// 					outputCh <- res
// 					continue
// 				case circle:
// 					res.Circle = i3.area((rand.Float64() * 5) + 10)
// 					outputCh <- res
// 					continue
// 				}
// 			}
// 		}
// 	}()
// 	return outputCh
// }

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
	i1 := square{}
	i2 := rectangle{}
	i3 := circle{}

	stream := make(chan interface{})
	defer close(stream)
	geo := make(chan interface{})
	defer close(geo)

	initInstance := MyStream{Square: 0, Rectangle: 0, Circle: 0}

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

	go func() {
		for {
			select {
			case o1 := <-geo:
				// fmt.Printf("Underlying Type: %T\n", o1)
				switch o1.(type) {
				case *square:
					initInstance.Square = i1.area((rand.Float64() * 5) + 25)
					stream <- initInstance
				case *rectangle:
					initInstance.Rectangle = i2.area((rand.Float64()*5)+25, (rand.Float64()*5)+25)
					stream <- initInstance
				case *circle:
					initInstance.Circle = i3.area((rand.Float64() * 5) + 10)
					stream <- initInstance

				}

			}
		}
	}()

	for {
		fmt.Printf("MyStream : %v\n", <-stream)
	}

	// stream <- GiveMeStream(ctx, geo)
	// defer cancel()

}
