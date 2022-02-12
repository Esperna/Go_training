package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	length := len(os.Args)
	if length != 2 {
		fmt.Println("Invalid Number of Argument")
		os.Exit(1)
	}
	r := []rune(os.Args[1])
	r = compressUnicodeSpaces(r)
	fmt.Println(string(r))
}

func compressUnicodeSpaces(s []rune) []rune {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if unicode.IsSpace(s[i+1]) && unicode.IsSpace(s[i]) {
				s = remove(s, i)
				s[i] = 0x20
				s = compressUnicodeSpaces(s)
			}
		}
	}
	return s
}

func remove(slice []rune, i int) []rune {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
