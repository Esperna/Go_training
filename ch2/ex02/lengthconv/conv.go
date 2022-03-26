package lengthconv

// MToFt converts a Celsius temperature to Fahrenheit.
func MToFt(m Meter) Feet { return Feet(m / 0.3408) }

// FtToM converts a Feet to Meter.
func FtToM(ft Feet) Meter { return Meter(ft * 0.3408) }

//!-
