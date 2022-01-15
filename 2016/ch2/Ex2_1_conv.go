// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CtoK converts a Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC)}

// KtoC converts a Kelvin temperature to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - AbsoluteZeroC)}
 
//!-
