package main

import (
	"fmt"
)

func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func removeDuplicates(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = remove(s, i+1)
			i--
		}
	}

	return s
}

func main() {
	s := []int{0, 0, 1, 2, 3, 3, 3, 4, 4, 3}
	fmt.Println(s)

	s = removeDuplicates(s)
	fmt.Println(s)
}
