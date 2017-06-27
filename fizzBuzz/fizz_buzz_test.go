package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzCollection(t *testing.T) {
	assert.Equal(t, FizzBuzzCollection(5), map[int]string{1: "1", 2: "2", 3: "fizz", 4: "4", 5: "buzz"},
		"If given size 5, counts to 5 returning fizz/buzz value for each")
}

func TestFizzBuzzItem(t *testing.T) {
	assert.Equal(t, FizzBuzzItem(4), "4", "Given 4, returns 4")
	assert.Equal(t, FizzBuzzItem(99), "fizz", "Given 99, returns fizz")
	assert.Equal(t, FizzBuzzItem(245), "buzz", "Given 245, returns buzz")
	assert.Equal(t, FizzBuzzItem(15), "fizz-buzz", "Given 15, returns fizz-buzz")
}
