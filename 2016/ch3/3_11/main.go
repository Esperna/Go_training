// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign, integer, decimal string

	//check sign(+/-) of input
	if strings.HasPrefix(s, "+") {
		sign = "+"
		s = s[1:]
	} else if strings.HasPrefix(s, "-") {
		sign = "-"
		s = s[1:]
	} else {
		sign = "+"
	}

	//seperate input into integer and decimal part
	if strings.Contains(s, ".") {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '.' {
				integer = s[:i]
				decimal = s[i:]
				break
			}
		}
	} else {
		integer = s
		decimal = ""
	}

	//insert "," into integer part
	n := len(integer)
	if n <= 3 {
		return sign + integer + decimal
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if (i > 0) && ((((n - 3) - i) % 3) == 0) {
			buf.WriteString(",")
		}
		buf.WriteByte(integer[i])
	}
	integer = buf.String()

	return sign + integer + decimal
}
