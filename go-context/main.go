package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func doSomethingLong(site string) string {
	result, err := http.Get("http://www." + site)
	if err != nil {
		panic(err)
	}
	time.Sleep(5* time.Second)
	return fmt.Sprintf("[%d] %s \n", result.StatusCode, site)
}
func doSomethingUsingContext(ctx context.Context, site string) (string, error) {
	res := make(chan string)
	go func() {
		res <- doSomethingLong(site)
		close(res)
	}()
	// Wait for events to come
	for {
		select {
		case result := <-res:
			return result, nil
		case <-ctx.Done():
			return "done", ctx.Err()
		}
	}
}
func main() {
	rootContext := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(rootContext, time.Duration(3000)*time.Millisecond)
	defer cancel()
	response, err := doSomethingUsingContext(ctxWithTimeout, "google.com")
	if err != nil {
		fmt.Println("got err: ", err)
		return
	}
	fmt.Println("response from:", response)

}
