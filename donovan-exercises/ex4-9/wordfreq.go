package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func wordfreq(filename string, counts map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	counts := make(map[string]int)

	wordfreq("input.txt", counts)
	fmt.Println(counts)
}
