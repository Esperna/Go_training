// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//
func MToFT(m Meter) Feet { return Feet(3.28084 * m) }

//
func FTToMT(ft Feet) Meter { return Meter(ft / 3.28084)}

//
func PDToKG(pd Pond) Kilogram { return Kilogram(pd / 2.20462)}

//
func KGToPD(kg Kilogram) Pond { return Pond(2.20462 * kg)}
//!-
