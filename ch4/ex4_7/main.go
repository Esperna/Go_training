package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length == 2 {
		s := os.Args[1]
		b := []byte(s)
		fmt.Println(s)
		reverse(b)
		fmt.Println(string(b))
	} else {
		fmt.Println("Invalid Number of Argument")
	}
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
