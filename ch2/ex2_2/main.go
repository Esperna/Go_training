// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"ch2/ex2_2/lengthconv"
	"ch2/ex2_2/tempconv"
	"ch2/ex2_2/weightconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s \n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
			ft := lengthconv.Feet(t)
			m := lengthconv.Meter(t)
			fmt.Printf("%s = %s, %s = %s \n",
				ft, lengthconv.FtToM(ft), m, lengthconv.MToFt(m))
			lb := weightconv.Pound(t)
			kg := weightconv.Kilogram(t)
			fmt.Printf("%s = %s, %s = %s \n",
				lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))
		}
	} else {
		var i int
		fmt.Scan(&i)
		a := strconv.Itoa(i)
		t, err := strconv.ParseFloat(a, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s \n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		ft := lengthconv.Feet(t)
		m := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s \n",
			ft, lengthconv.FtToM(ft), m, lengthconv.MToFt(m))
		lb := weightconv.Pound(t)
		kg := weightconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s \n",
			lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))
	}
}

//!-
