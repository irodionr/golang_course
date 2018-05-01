package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func sha(flag string, data []byte) []byte {
	var hash []byte

	switch flag {
	case "-sha256":
		arr := sha256.Sum256(data)
		hash = arr[:]
	case "-sha384":
		arr := sha512.Sum384(data)
		hash = arr[:]
	case "-sha512":
		arr := sha512.Sum512(data)
		hash = arr[:]
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag:", flag)
	}

	return hash
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println(sha(os.Args[1], []byte(strings.Join(os.Args[2:], " "))))
	} else {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments:", len(os.Args)-1)
	}
}
