// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }

//!-
