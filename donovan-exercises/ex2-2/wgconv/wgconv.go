// Package wgconv performs kilograms and pounds conversions.
package wgconv

import "fmt"

// Kilograms weight - kg
type Kilograms float64

// Pounds weight - lb
type Pounds float64

func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%glb", p) }
