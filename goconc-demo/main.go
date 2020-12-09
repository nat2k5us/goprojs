package main

import (
	"fmt"
	"time"
)

func main() {
	go dosomething("order received ")
	go dosomething("prepare pizza ")
	fmt.Scanln()
	fmt.Println("program exit..")
}

func dosomething(job string) {
	for i := 0; true; i++ {
		fmt.Println(job, i)
		time.Sleep(time.Millisecond * 500)
	}

}
