package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length >= 2 {
		s := os.Args[1:]
		fmt.Println(s)
		s = deleteDup(s)
		fmt.Println(s)
	} else {
		fmt.Println("Invalid Number of Argument")
	}
}

func deleteDup(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i+1] == s[i] {
				s = remove(s, i)
				s = deleteDup(s)
			}
		}
	}
	return s
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
