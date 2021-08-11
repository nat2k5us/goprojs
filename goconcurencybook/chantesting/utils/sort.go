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

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}