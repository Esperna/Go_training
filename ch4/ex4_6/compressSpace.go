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
		r, size := utf8.DecodeRune(b[i:])
		if size < 1 || size > 4 {
			return nil
		}
		if i+size-1 < len(b) {
			if unicode.IsSpace(r) {
				for j := i + size - 1; j > i; j-- {
					b = remove(b, j)
				}
				b[i] = ' '
				i++
			} else {
				i += size
			}
		}
	}
	b = deleteDupSpace(b)
	return b
}
func deleteDupSpace(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			r1 := []rune(string(s[i+1]))
			r0 := []rune(string(s[i]))
			if unicode.IsSpace(r1[0]) && unicode.IsSpace(r0[0]) {
				s = remove(s, i)
				s = deleteDupSpace(s)
			}
		}
	}
	return s
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
