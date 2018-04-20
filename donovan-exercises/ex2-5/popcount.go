package main

import (
	"fmt"
	"time"
)

// PopCount4 is PopCount that counts bits by using the fact that x&(x-1) clears the rightmost non-zero bit of x.
func PopCount4(x uint64) int {
	var c int

	for x != 0 {
		x &= (x - 1)
		c++
	}

	return c
}

func main() {
	arr := []uint64{7, 10, 64, 127, 256}

	start := time.Now()
	for _, x := range arr {
		fmt.Printf("PopCount4(%d) = %d\n", x, PopCount4(x))
	}
	fmt.Printf("Time elapsed: %v s\n", time.Since(start).Seconds())
}
