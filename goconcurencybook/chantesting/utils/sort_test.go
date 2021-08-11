package sort

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 12, 8, 6, 7, 13, 14, 15, 16}
	BubbleSort(elements)
	log.Printf("%v elements", elements)
	if elements[0] != 16 {
		t.Errorf("first element should be 16")
	}
	if elements[len(elements)-1] != 1 {
		t.Errorf("last element should be 1")
	}
}
func TestBubbleSortWithTimeout(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 12, 8, 6, 7, 13, 14, 15, 16}
	timeOut := make(chan bool, 1)
	go func() {
		BubbleSort(elements)
		timeOut <- false
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		timeOut <- true
	}()

	if <-timeOut {
		assert.Fail(t, "Timed out")
		return
	}
	log.Printf("%v elements", elements)
	if elements[0] != 16 {
		t.Errorf("first element should be 16")
	}
	if elements[len(elements)-1] != 1 {
		t.Errorf("last element should be 1")
	}
}

func TestFib(t *testing.T) {

	tests := []struct {
		testItem int
		expected int
	}{
		{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
	}
	for _, tt := range tests {

		actual := Fib(tt.testItem)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.testItem, tt.expected, actual)
		}
	}

}
