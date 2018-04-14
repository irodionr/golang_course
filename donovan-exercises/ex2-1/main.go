package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)                 // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))                     // "212°F"
	fmt.Println(tempconv.CToK(tempconv.FreezingC))                    // "273.15K"
	fmt.Println(tempconv.FToK(tempconv.CToF(tempconv.AbsoluteZeroC))) // "0K"
}
