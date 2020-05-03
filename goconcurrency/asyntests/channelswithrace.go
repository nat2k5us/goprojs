package asyntests

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func asyncUsingChannelsToAvoidRaceConditions() {
	a := 1
	b := 2

	operationDone := make(chan bool) // creating the channel
	go func() {                      // this is how you make the async call in go
		b = a * b

		operationDone <- true
	}()

	// this tells the main thread to wait for a message to be pushed into channel (i.e bool)
	<-operationDone

	a = b * b

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()

	fmt.Printf("a = %d, b = %d\n", a, b)
}

func WaitingOnAllRoutinesToFinish() {
	rand.Seed(time.Now().UTC().UnixNano())

	respond := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	go checkDNS(respond, &wg, "pragmacoders.com", "ns1.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns2.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns3.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns4.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns5.nameserver.com")

	wg.Wait()
	close(respond)

	for queryResp := range respond {
		fmt.Printf("Got Response:\t %s\n", queryResp)
	}
}

func checkDNS(respond chan<- string, wg *sync.WaitGroup, query string, ns string) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	respond <- fmt.Sprintf("%s responded to query: %s", ns, query)
}
