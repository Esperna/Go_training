package weightconv

// KgToLb converts a Celsius temperature to Fahrenheit.
func KgToLb(kg Kilogram) Pound { return Pound(kg / 2.20462) }

// LbToKg converts a Feet to Meter.
func LbToKg(lb Pound) Kilogram { return Kilogram(lb * 2.20642) }

//!-
