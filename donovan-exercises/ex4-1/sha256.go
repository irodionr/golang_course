package main

import (
	"crypto/sha256"
	"fmt"
)

func bitsDiff(a, b byte) int {
	var count int

	for a != 0 {
		if a&1 != b&1 {
			count++
		}

		a >>= 1
		b >>= 1
	}

	return count
}

func digestsDiff(a, b *[32]byte) int {
	var count int

	for i := 0; i < len(a); i++ {
		count += bitsDiff(a[i], b[i])
	}

	return count
}

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n", a, b, a == b)

	fmt.Println("Different bits: ", digestsDiff(&a, &b))
}
