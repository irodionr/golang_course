package main

import (
	"fmt"
	"strings"
)

func anagram(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)

		chars1 := make(map[byte]int)
		chars2 := make(map[byte]int)

		for i := 0; i < len(s1); i++ {
			chars1[s1[i]]++
			chars2[s2[i]]++
		}

		for c, count := range chars1 {
			if chars2[c] != count {
				return false
			}
		}

		return true
	}

	return false
}

func main() {
	s1 := "Hello"
	s2 := "Aloha"
	s3 := "Hell"
	s4 := "Haalo"
	s5 := "Heloo"

	fmt.Println(s1, s2, anagram(s1, s2))
	fmt.Println(s1, s3, anagram(s1, s3))
	fmt.Println(s1, s4, anagram(s1, s4))
	fmt.Println(s1, s5, anagram(s1, s5))
	fmt.Println(s2, s4, anagram(s2, s4))
}
