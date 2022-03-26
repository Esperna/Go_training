package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Invalid number of Argument. Expected is like 'reverseByteSlice string'")
	}
	b := []byte(os.Args[1])
	reverse(b)
	fmt.Println(string(b))
}

func reverse(b []byte) {
	r, size := utf8.DecodeRune(b)
	if r == utf8.RuneError && size == 1 {
		panic(fmt.Sprintf("Invalid rune %q size: %d", r, size))
	} else if r == utf8.RuneError && size == 0 {
		return
	}
	rotate(b, size)
	reverse(b[:len(b)-size])
}

func rotate(b []byte, n int) {
	reverseByteSlice(b[:n])
	reverseByteSlice(b[n:])
	reverseByteSlice(b)
}

func reverseByteSlice(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
