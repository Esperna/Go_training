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
	var buf bytes.Buffer
	i := 0
	j := 0
	k := 0
	n := len(s)
	if s[i] == '+' || s[i] == '-' {
		buf.WriteString(s[i:i])
		i++
		j = i
	}
	for i < len(s) {
		if s[i] == '.' {
			k = i
			break
		}
		i++
	}
	n = k - j
	i = n%3 + j
	if j < i {
		buf.WriteString(s[j:i])
	}
	for n >= 3 {
		if i != j {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
		i += 3
		n -= 3
	}
	buf.WriteString(s[k:])
	return buf.String()
}

//!-
