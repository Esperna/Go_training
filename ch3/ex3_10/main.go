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
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	i := n % 3
	//	sum := s[:i]
	buf.WriteString(s[:i])
	for n >= 3 {
		if i > 0 {
			//	sum += ","
			buf.WriteString(",")
		}
		//		sum += s[i : i+3]
		buf.WriteString(s[i : i+3])
		i += 3
		n -= 3
	}
	//return sum
	return buf.String()
}

//!-
