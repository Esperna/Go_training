// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gFt", ft) }

//!-
