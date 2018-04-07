package main

import "golang.org/x/tour/reader"

// MyReader emits an infinite stream of the ASCII character 'A'
type MyReader struct{}

func (reader MyReader) Read(a []byte) (int, error) {
	for i := range a {
		a[i] = 'A'
	}

	return len(a), nil
}

func main() {
	reader.Validate(MyReader{})
}
