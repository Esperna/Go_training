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

type unicodeCount struct {
	letterCount, numberCount, spaceCount int
}

func main() {
	ct := countUnicode(os.Stdin)
	fmt.Printf("Letter count\t%d\n", ct.letterCount)
	fmt.Printf("Number count\t%d\n", ct.numberCount)
	fmt.Printf("Space count\t%d\n", ct.spaceCount)
}

func countUnicode(rd io.Reader) unicodeCount {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	var invalid, numberCount, letterCount, spaceCount int
	in := bufio.NewReader(rd)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCount++
		}
		if unicode.IsNumber(r) {
			numberCount++
		}
		if unicode.IsSpace(r) {
			spaceCount++
		}
		counts[r]++
		utflen[n]++
	}
	return unicodeCount{letterCount, numberCount, spaceCount}
}
