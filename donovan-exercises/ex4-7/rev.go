package main

import (
	"fmt"
)

// reverse reverses an array of ints in place.
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	s := "Hello, World!"

	fmt.Printf("%s\n%s\n", s, reverse([]byte(s)))
}
