// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	data, err := os.Open("top-1m.csv") // http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		if err := data.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	reader := csv.NewReader(data)
	ch := make(chan string)
	lineCount := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		lineCount++

		go fetch("http://"+record[1], ch) // start a goroutine
	}

	out, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		if err := out.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for i := 0; i < lineCount; i++ {
		fmt.Fprintln(out, <-ch) // receive from channel ch
	}

	fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
