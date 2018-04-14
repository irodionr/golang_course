// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius temperature - °C
type Celsius float64

// Fahrenheit temperature - °F
type Fahrenheit float64

// Kelvin temperature - K
type Kelvin float64

// AbsoluteZeroC - absolute zero, FreezingC - freezing, BoilingC - boiling
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
