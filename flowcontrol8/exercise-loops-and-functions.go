package main

import (
	"fmt"
	"math"
)

// Sqrt - square root
func Sqrt(x float64) float64 {
	z := 1.0 // 1.0, x, x / 2
	prev := 0.0
	eps := 0.001
	i := 0

	fmt.Println(z)

	for math.Abs(z-prev) > eps {
		prev = z
		z -= (z*z - x) / (2 * z)
		i++

		fmt.Println(z)
	}

	fmt.Println(i, "iterations")

	return z
}

func main() {
	x := 64.0
	mySqrtX := Sqrt(x)
	mathSqrtX := math.Sqrt(x)

	fmt.Println("Result:", mySqrtX)
	fmt.Println("math.Sqrt():", mathSqrtX)
	fmt.Println("Difference:", mySqrtX-mathSqrtX)
}
