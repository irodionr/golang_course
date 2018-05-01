package main

import (
	"fmt"
)

const size = 6

// reverse reverses an array of ints in place.
func reverse(s *[size]int) {
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var s [size]int
	for i := 0; i < size; i++ {
		s[i] = i
	}

	reverse(&s)

	fmt.Println(s)
}
