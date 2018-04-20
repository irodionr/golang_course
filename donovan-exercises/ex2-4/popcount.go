package main

import (
	"fmt"
	"time"
)

// PopCount3 counts bits by shifting its argument through 64 bits positions, testing rightmost bit each time.
func PopCount3(x uint64) int {
	var c int

	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			c++
		}
	}

	return c
}

func main() {
	arr := []uint64{7, 10, 64, 127, 256}

	start := time.Now()
	for _, x := range arr {
		fmt.Printf("PopCount3(%d) = %d\n", x, PopCount3(x))
	}
	fmt.Printf("Time elapsed: %v s\n", time.Since(start).Seconds())
}
