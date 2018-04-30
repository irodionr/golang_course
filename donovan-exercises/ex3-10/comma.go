package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if ((i-n%3)%3 == 0 || i == n%3) && i != 0 {
			buf.WriteString(",")
		}

		fmt.Fprintf(&buf, "%c", s[i])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234567899"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12"))
	fmt.Println(comma("1"))
}
