package main

import (
	"bytes"
	"fmt"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign, mant string
	var buf bytes.Buffer

	if s[0] == '-' {
		sign = "-"
		s = s[1:]
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		mant = s[dot:]
		s = s[:dot]
	}

	n := len(s)

	for i := 0; i < n; i++ {
		if ((i-n%3)%3 == 0 || i == n%3) && i != 0 {
			buf.WriteString(",")
		}

		fmt.Fprintf(&buf, "%c", s[i])
	}

	return sign + buf.String() + mant
}

func main() {
	fmt.Println(comma("-1234567899.12345"))
	fmt.Println(comma("-123456789"))
	fmt.Println(comma("-12345678.12"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("1234.1234567"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12"))
	fmt.Println(comma("-1"))
}
