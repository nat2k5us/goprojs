package main

import (
	"fmt"
	"net/http"
)

func httpRequest(site string) string {
	r, e := http.Get(site)
	if e != nil {
		panic(e)
	}
	fmt.Printf("requesting %s \n", site)
	return fmt.Sprintf("[%d] %s \n", r.StatusCode, site)
}

func TestChannelsUsingHttp() {
	c1 := make(chan string)
	c2 := make(chan string)

	websites1 := []string{"google.com", "amazon.com", "youtube.com", "apple.com", "microsoft.com"}
	websites2 := []string{"whatismyip.com", "cnn.com", "cnbc.com", "cloudflare.com", "webmd.com"}
	go func() {
		for _, site := range websites1 {
			c1 <- httpRequest("http://www." + site)
		}

	}()
	go func() {
		for _, site := range websites2 {
			c2 <- httpRequest("http://www." + site)
		}

	}()
	counter := 0
	for {
		if counter > (len(websites1) + len(websites2)) {
			break
		}

		select {

		case <-c2:
			fmt.Printf("%d => %s", counter, <-c2)
			counter++
		case <-c1:
			fmt.Printf("%d => %s", counter, <-c1)
			counter++
		default:
			//fmt.Print("no activity")
		}

	}

}

func TestSyncCallHttpRequest(){
	websites1 := []string{"google.com", "amazon.com", "youtube.com", "apple.com", "microsoft.com","whatismyip.com", "cnn.com", "cnbc.com", "cloudflare.com", "webmd.com"}
	
	go func() {
		for _, site := range websites1 {
			fmt.Printf("Response => %s", httpRequest("http://www." + site))
		}

	}()
}

func main() {
	fmt.Printf("Response => %s", httpRequest("http://www.google.com"))
}
