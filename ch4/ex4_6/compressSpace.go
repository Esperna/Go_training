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
	b := []byte(os.Args[1])
	b = compressUnicodeSpaces(b)
	fmt.Println(string(b))
}

func compressUnicodeSpaces(b []byte) []byte {
	for i := 0; i < len(b); {
		isASCII := (b[i] >> 7) == 0
		is2Byte := (b[i] >> 5) == 0b110
		is3Byte := (b[i] >> 5) == 0b111
		if isASCII {
			if unicode.IsSpace(rune(b[i])) {
				b[i] = ' '
			}
			i++
		} else if is2Byte {
			if i+1 < len(b) {
				r := []rune(string(b[i : i+2]))
				if unicode.IsSpace(r[0]) {
					b = remove(b, i+1)
					b[i] = ' '
					i++
				} else {
					i += 2
				}
			}
		} else if is3Byte {
			if i+2 < len(b) {
				r := []rune(string(b[i : i+3]))
				if unicode.IsSpace(r[0]) {
					b = remove(b, i+2)
					b = remove(b, i+1)
					b[i] = ' '
					i++
				} else {
					i += 3
				}
			}
		} else {
			i++
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
