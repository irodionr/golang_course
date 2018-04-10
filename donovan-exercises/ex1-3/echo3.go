package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	start := time.Now()
	var s1, sep1 string
	for i := 1; i < len(os.Args); i++ {
		s1 += sep1 + os.Args[i]
		sep1 = " "
	}
	fmt.Println(s1)
	fmt.Println(time.Since(start))
}

func echo2() {
	start := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	fmt.Println(time.Since(start))
}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start))
}

func main() {
	echo1()
	echo2()
	echo3()
}
