// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type charCount struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
	err     error
}

func main() {
	charcount := countChar(bufio.NewReader(os.Stdin))
	if charcount.err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", charcount.err)
		os.Exit(1)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range charcount.counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range charcount.utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if charcount.invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", charcount.invalid)
	}
}

func countChar(rd io.Reader) charCount {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(rd)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return charCount{counts, utflen, invalid, err}
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return charCount{counts, utflen, invalid, nil}
}

//!-
