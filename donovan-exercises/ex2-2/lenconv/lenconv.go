// Package lenconv performs meters and feet conversions.
package lenconv

import "fmt"

// Meters length - m
type Meters float64

// Feet length - ft
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }
