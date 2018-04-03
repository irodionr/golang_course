package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree sending all values from the tree to the channel
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees contain the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(3), tree.New(5)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
}
