package asyntests

import (
	"sync"
	"time"
)

func AsyncSequentialCallToRoutines(wg sync.WaitGroup) {
	wg.Add(3)
	var firstUserData, secondUserData, thirdUserData string
	// First DB call
	go func() {

		firstUserData = firstDbCall()
		defer wg.Done()
	}()

	// Second DB call
	go func() {

		secondUserData = secondDbCall()
		defer wg.Done()
	}()

	// Third DB call
	go func() {

		thirdUserData = thirdDbCall()
		defer wg.Done()
	}()

	wg.Wait()

	println(firstUserData, secondUserData, thirdUserData)
}

// Group 1
func firstDbCall() string {

	time.Sleep(time.Duration(15 * time.Second))
	return "UserId1"
}

func secondDbCall() string {

	time.Sleep(time.Duration(10 * time.Second))
	return "UserId2"
}

func thirdDbCall() string {

	time.Sleep(time.Duration(5 * time.Second))
	return "UserId3"
}
