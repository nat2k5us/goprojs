package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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

func parseInt(i interface{}) (int, error) {
	s, ok := i.(string)
	if !ok {
		return 0, errors.New("not string")
	}
	return strconv.Atoi(s)
}

//       ch
//------------------------------------------------------------------------------
//
//       ch <- 4                                         <-ch
//----------------------------------------------------------------
//       ^                                               ch = 4

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	sumit := 0
	go func() {
		for {
			time.Sleep(time.Second)
			t1 := rand.Intn(50)
			fmt.Println("Add to 1:", t1)
			ch1 <- t1
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			t2 := 1000 + rand.Intn(50)
			fmt.Println("Add to 2:", t2)
			ch2 <- t2
		}
	}()

	for {
		select {
		case o1 := <-ch1:
			fmt.Printf("\t  1: %v\n", o1)
			temp, _ := o1.(int)
			sumit += temp
			fmt.Printf("sumit %d\n", sumit)
		case o2 := <-ch2:
			fmt.Printf("\t  2: %v\n", o2)
			temp, _ := o2.(int)
			sumit += temp
			fmt.Printf("sumit %d\n", sumit)
		}
	}

}
