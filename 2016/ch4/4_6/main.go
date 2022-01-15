package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "foo  foo\n"
	fmt.Println(s)
	fmt.Println(string(removeDuplication([]byte(s))))
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeDuplication(slice []byte) []byte {
	for i := 0; i < len(slice)-1; i++ {
		if unicode.IsSpace(rune(slice[i])) {
			if slice[i] == slice[i+1] {
				return removeDuplication(remove(slice, i))
			}
		}
	}
	return slice[:]
}
