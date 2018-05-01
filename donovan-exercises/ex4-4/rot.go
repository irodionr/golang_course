package main

import "fmt"

func rotate(s []int, n int) {
	s1 := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		s1[i] = s[(i+n)%(len(s))]
	}

	copy(s, s1)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	rotate(s, 2)

	fmt.Println(s)
}
