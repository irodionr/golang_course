package wgconv

// KToP converts a weight in kilograms to weight in pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 2.2046) }

// PToK converts a weight in pounds to weight in kilograms.
func PToK(p Pounds) Kilograms { return Kilograms(p / 2.2046) }
