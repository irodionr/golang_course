package main

import (
	"fmt"
)

// KB - kilobytes, ...
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	//fmt.Println(YB) // overflows int
}
