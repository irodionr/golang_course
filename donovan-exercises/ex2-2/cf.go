// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../ex2-1/tempconv"
	"./lenconv"
	"./wgconv"
)

func convert(args []string) {
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		m := lenconv.Meters(t)
		ft := lenconv.Feet(t)
		k := wgconv.Kilograms(t)
		p := wgconv.Pounds(t)

		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToF(c), f, tempconv.FToC(f))
		fmt.Printf("%s = %s, %s = %s\n",
			m, lenconv.MToF(m), ft, lenconv.FToM(ft))
		fmt.Printf("%s = %s, %s = %s\n",
			k, wgconv.KToP(k), p, wgconv.PToK(p))
	}
}

func main() {
	if len(os.Args) > 1 {
		convert(os.Args[1:])
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter numbers: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		args := strings.Fields(input)
		convert(args)
	}
}
