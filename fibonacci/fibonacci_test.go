package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(1), 1, "Given 1, return 1.")
	assert.Equal(t, Fibonacci(0), 0, "Given 0, return 0.")
	assert.Equal(t, Fibonacci(4), 3, "Given 4, return 3.")
}

func TestFibonacciSeries(t *testing.T) {
	assert.Equal(t, FibonacciSeries(0, 10), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, "When 10 numbers requested, returns correct sequence.")
	assert.Equal(t, FibonacciSeries(0, 20), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}, "When 20 numbers requested, returns correct 20 number sequence.")
	assert.Equal(t, FibonacciSeries(0, 5), []int{0, 1, 1, 2, 3, 5}, "Handles shorter sequences.")
}
