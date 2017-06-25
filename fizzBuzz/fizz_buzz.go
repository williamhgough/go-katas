package main

import (
	"fmt"
	"strconv"
)

func main() {
	FizzBuzzCollection(5)
}

// FizzBuzzCollection returns a map of int keys and their fizz/buzz string value
func FizzBuzzCollection(size int) map[int]string {
	items := make(map[int]string, size)

	for i := 1; i <= size; i++ {
		if i%2 == 0 && i%5 == 0 {
			items[i] = "fizz-buzz"
		} else if i%3 == 0 {
			items[i] = "fizz"
		} else if i%5 == 0 {
			items[i] = "buzz"
		} else {
			items[i] = strconv.Itoa(i)
		}
	}
	fmt.Println(items)
	return items
}

// FizzBuzzItem will return the fizz/buzz value for given int.
func FizzBuzzItem(item int) string {
	if item%3 == 0 && item%5 == 0 {
		return "fizz-buzz"
	} else if item%3 == 0 {
		return "fizz"
	} else if item%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(item)
}
