package asyntests

import "fmt"

// Group 2
func firstDBCall(resultSlice *[]string, doneChannel chan bool) {
	(*resultSlice)[0] = "1"
	doneChannel <- true
}

func secondDBCall(resultSlice *[]string, doneChannel chan bool) {
	(*resultSlice)[1] = "2"
	doneChannel <- true
}

func thirdDBCall(resultSlice *[]string, doneChannel chan bool) {
	(*resultSlice)[2] = "3"
	doneChannel <- true
}

func AsyncSequenciallyUsingChannelsAndSlices() {
	resultSlice := make([]string, 3)
	doneChannel := make(chan bool)
	go firstDBCall(&resultSlice, doneChannel)
	go secondDBCall(&resultSlice, doneChannel)
	go thirdDBCall(&resultSlice, doneChannel)

	for i := 0; i < 3; i++ {
		<-doneChannel
	}
	fmt.Println(resultSlice)
}
