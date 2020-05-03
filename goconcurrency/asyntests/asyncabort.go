package asyntests

import "sync"

func AsyncAbortRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)

	wg.Wait()
}
