// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
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
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	if s == "" {
		return ""
	}
	var buf bytes.Buffer
	const seprateNum = 3
	var i, j, k int
	n := len(s)
	hasDecimal := false
	if s[i] == '.' {
		buf.WriteString(s[i:])
		return buf.String()
	} else if s[i] == '+' || s[i] == '-' {
		buf.WriteString(s[i : i+1])
		i++
		j = i
	}
	for i < len(s) {
		if s[i] == '.' {
			k = i
			hasDecimal = true
			break
		}
		i++
	}
	if k > j {
		n = k - j
	}
	i = n%seprateNum + j
	if j < i {
		buf.WriteString(s[j:i])
	}
	for n >= seprateNum {
		if i != j {
			buf.WriteString(",")
		}
		if i+seprateNum < len(s) {
			buf.WriteString(s[i : i+seprateNum])
		} else {
			buf.WriteString(s[i:])
		}

		i += seprateNum
		n -= seprateNum
	}
	if hasDecimal {
		buf.WriteString(s[k:])
	}
	return buf.String()
}

//!-
