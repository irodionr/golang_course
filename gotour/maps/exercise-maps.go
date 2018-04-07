package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount - map of the counts of each "word" in the string
func WordCount(s string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(s) {
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
