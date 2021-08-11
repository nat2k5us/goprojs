package sort

func BubbleSort(elements []int){
	keepGoing := true
	for keepGoing {
		keepGoing = false
		for  i :=0 ;i < len(elements)-1; i++ {
			if elements[i] <  elements[i+1] {
				keepGoing = true
				elements[i],elements[i+1] = elements[i+1],elements[i]

			}
		}
	}
}