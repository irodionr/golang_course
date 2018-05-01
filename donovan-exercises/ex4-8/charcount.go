// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	var letterCount, digitCount int // count letters and digits
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF || r == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letterCount++
		} else if unicode.IsDigit(r) {
			digitCount++
		}

		utflen[n]++
	}

	fmt.Printf("%d letters, %d digits\n", letterCount, digitCount)

	fmt.Print("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
