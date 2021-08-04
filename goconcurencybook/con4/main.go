package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// for _, line := range []string{"hello", "greetings", "good day","you", "can", "omit", "semicolons", "if", "there", "is", "only", "a", "condition"} {
	// 	wg.Add(1)
	// 	go func() {
	// 		wg.Done()
	// 		fmt.Println(line)
	// 	}()
	// }
	// wg.Wait()

	// fixed version 
	fmt.Println(" -------- fixed version -----------")
	for _, line := range []string{"hello", "greetings", "good day","you", "can", "omit", "semicolons", "if", "there", "is", "only", "a", "condition"} {
		wg.Add(1)
		go func(line string) {
			wg.Done()
			fmt.Println(line)
		}(line)
	}
	wg.Wait()
}

////////////////// OUTPUT ////////////////////
