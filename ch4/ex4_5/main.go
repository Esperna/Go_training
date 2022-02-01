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
		b = deleteDup(b)
		fmt.Println(string(b))
	} else {
		fmt.Println("Invalid Number of Argument")
	}
}

func deleteDup(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		t := s[i]
		for j := i + 1; j < len(s); j++ {
			if s[j] == t {
				s = remove(s, j)
				s = deleteDup(s)
			}
		}
	}
	return s
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
