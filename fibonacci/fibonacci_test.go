package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(10), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, "Given 10, return the first 10 numbers of fibo sequence.")
}

func TestFibonacciPrevious(t *testing.T) {
	assert.Equal(t, FibonacciPrevious(5), 3, "Given 5, returns previous sequence number 3.")
}

func TestFibonacciNext(t *testing.T) {
	assert.Equal(t, FibonacciNext(5), 3, "Given 5, returns next sequence number 8.")
}
