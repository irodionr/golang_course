package lenconv

// MToF converts a length in meters to feet.
func MToF(m Meters) Feet { return Feet(m / 0.3048) }

// FToM converts a length in feet to meters.
func FToM(f Feet) Meters { return Meters(f * 0.3048) }
