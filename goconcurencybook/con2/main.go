package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex // guards
	value int
}

var wg sync.WaitGroup

func printSum(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()
	time.Sleep(1000 * time.Millisecond)
	v2.mu.Lock()
	defer v2.mu.Unlock()
	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

func main() {
	var a, b value
	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

////////////////// OUTPUT ////////////////////
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0x118f3e0)
//         /usr/local/go/src/runtime/sema.go:56 +0x45
// sync.(*WaitGroup).Wait(0x118f3d8)
//         /usr/local/go/src/sync/waitgroup.go:130 +0x65
// main.main()
//         /Users/nbontha/dev/goprojs/goconcurencybook/con2/main.go:33 +0xd6

// goroutine 6 [semacquire]:
// sync.runtime_SemacquireMutex(0xc0000140e4, 0x1061600, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc0000140e0)
//         /usr/local/go/src/sync/mutex.go:138 +0x105
// sync.(*Mutex).Lock(...)
//         /usr/local/go/src/sync/mutex.go:81
// main.PrintDum(0xc0000140d0, 0xc0000140e0)
//         /Users/nbontha/dev/goprojs/goconcurencybook/con2/main.go:21 +0x1be
// created by main.main
//         /Users/nbontha/dev/goprojs/goconcurencybook/con2/main.go:30 +0x9a

// goroutine 7 [semacquire]:
// sync.runtime_SemacquireMutex(0xc0000140d4, 0x1061600, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc0000140d0)
//         /usr/local/go/src/sync/mutex.go:138 +0x105
// sync.(*Mutex).Lock(...)
//         /usr/local/go/src/sync/mutex.go:81
// main.PrintDum(0xc0000140e0, 0xc0000140d0)
//         /Users/nbontha/dev/goprojs/goconcurencybook/con2/main.go:21 +0x1be
// created by main.main
//         /Users/nbontha/dev/goprojs/goconcurencybook/con2/main.go:31 +0xc6
// exit status 2