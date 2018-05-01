package main

import (
	"fmt"
	"unicode"
)

func remove(s []byte, i int) []byte {
	return append(s[:i], s[i+1:]...)
}

func squashSpaces(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && s[i] == s[i+1] {
			s = remove(s, i)
			i--
		}
	}

	return s
}

func main() {
	s := "Hello,   W o  r l    d !  "

	fmt.Printf("%s\n%s\n", s, squashSpaces([]byte(s)))
}
