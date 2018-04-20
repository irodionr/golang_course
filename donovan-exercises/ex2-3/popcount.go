package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount2 is PopCount with a loop instead of a single expression.
func PopCount2(x uint64) int {
	var s byte

	for i := uint64(0); i < 8; i++ {
		s += pc[byte(x>>(i*8))]
	}

	return int(s)
}

func main() {
	arr := []uint64{7, 10, 64, 127, 256}

	start := time.Now()
	for _, x := range arr {
		fmt.Printf("PopCount2(%d) = %d\n", x, PopCount2(x))
	}
	fmt.Printf("Time elapsed: %v s\n", time.Since(start).Seconds())
}
