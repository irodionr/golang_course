package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt - square root of negative number
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// Sqrt - square root
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
