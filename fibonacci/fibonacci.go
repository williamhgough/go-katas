package fibonacci

func Fibonacci(length int) []int {
	items := make([]int, length)
	return items
}

func FibonacciPrevious(start int) func() int {
	return func() int {
		return 0
	}
}

func FibonacciNext(start int) func() int {
	return func() int {
		return 1
	}
}
