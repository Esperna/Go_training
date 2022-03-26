package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	length := len(os.Args)
	if length != 2 {
		fmt.Println("Invalid Number of Argument")
		os.Exit(1)
	}
	b := []byte(os.Args[1])
	b = compressUnicodeSpaces(b)
	fmt.Println(string(b))
}

func compressUnicodeSpaces(b []byte) []byte {
	for i := 0; i < len(b); {
		r1, size1 := utf8.DecodeRune(b[i:])
		if r1 == utf8.RuneError && size1 == 1 {
			panic(fmt.Sprintf("Invalid byte s %q %d", b, size1))
		}
		r2, size2 := utf8.DecodeRune(b[i+size1:])
		if r2 == utf8.RuneError && size2 == 1 {
			panic(fmt.Sprintf("Invalid byte s %q %d", b, size2))
		}
		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			b = remove(b, i, size1+size2)
			b = append(b[:i+1], b[i:]...)
			b[i] = ' '
		} else {
			i += size1
		}
	}
	return b
}

func remove(s []byte, i, size int) []byte {
	copy(s[i:], s[i+size:])
	return s[:len(s)-size]
}
