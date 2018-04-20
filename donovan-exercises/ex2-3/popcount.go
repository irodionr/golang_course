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

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
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
	arr := []uint64{1, 3, 7, 15, 31, 63, 127, 255}

	start := time.Now()
	for _, x := range arr {
		fmt.Printf("PopCount(%d) = %d\n", x, PopCount(x))
	}
	fmt.Printf("Time elapsed: %v s\n", time.Since(start).Seconds())

	start = time.Now()
	for _, x := range arr {
		fmt.Printf("PopCount2(%d) = %d\n", x, PopCount2(x))
	}
	fmt.Printf("Time elapsed: %v s\n", time.Since(start).Seconds())
}
