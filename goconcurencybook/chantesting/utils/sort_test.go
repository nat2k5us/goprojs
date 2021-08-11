package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 12, 8, 6, 7, 13, 14, 15, 16}
	BubbleSort(elements)
}
