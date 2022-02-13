package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length == 2 {
		b := []byte(os.Args[1])
		b = reverse(b)
		fmt.Println(string(b))
	} else {
		fmt.Println("Invalid Number of Argument")
	}
}

func reverse(b []byte) []byte {
	var isASCII, is2Byte, is3Byte bool
	for i := 0; i < len(b); i++ {
		if b[i]>>7 == 0 {
			isASCII = true
		}
		if b[i]>>5 == 0b110 {
			is2Byte = true
		} else if b[i]>>5 == 0b111 {
			is3Byte = true
		}
	}
	if isASCII && !is2Byte && !is3Byte {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	} else if !isASCII && is2Byte && !is3Byte {
		for i, j := 0, len(b)-1; i < j; i, j = i+2, j-2 {
			b[i], b[j-1] = b[j-1], b[i]
			b[i+1], b[j] = b[j], b[i+1]
		}
	} else if !isASCII && !is2Byte && is3Byte {
		for i, j := 0, len(b)-1; i < j; i, j = i+3, j-3 {
			b[i], b[j-2] = b[j-2], b[i]
			b[i+1], b[j-1] = b[j-1], b[i+1]
			b[i+2], b[j] = b[j], b[i+2]
		}
	} else {
		//gave up byte slice handling
		r := []rune(string(b))
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return []byte(string(r))
	}
	return b
}
