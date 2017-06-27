package n8s

import (
	"strconv"
	"strings"
)

// N8n will numerize the given word.
func N8s(word string) string {
	// Handle numbers
	_, err := strconv.ParseInt(word, 10, 64)
	if err == nil {
		return ""
	}

	// split string if spaces
	word = strings.Join(strings.Split(word, " "), "")

	// grab first char
	first := word[0:1]

	// grab last char
	last := word[len(word)-1 : len(word)]

	// count middle char count
	count := strconv.Itoa(len(word) - 2)

	return first + count + last
}
