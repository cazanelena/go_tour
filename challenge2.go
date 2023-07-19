package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words = strings.Fields(s)
	count := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}


// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.
