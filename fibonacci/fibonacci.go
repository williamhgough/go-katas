package fibonacci

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-2) + Fibonacci(n-1)
	}
}

func FibonacciSeries(start, end int) []int {
	var fib []int
	for i := start; i < end+1; i++ {
		fib = append(fib, Fibonacci(i))
	}
	return fib
}
