package main

import (
	"fmt"
	"time"
)

func main() {
    var Ball int
    table := make(chan int)
    go player(table)
    go player(table)

    table <- Ball
    time.Sleep(1 * time.Second)
    <-table
}

func player(table chan int) {
    for {
		// Read the value fromtable and assign to ball - initial val will be 0
        ball := <-table
		fmt.Println(ball)
        ball++
        time.Sleep(100 * time.Millisecond)
		// put the increased value of ball back on table
        table <- ball
		
    }
}