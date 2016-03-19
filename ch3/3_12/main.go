package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_len := len(os.Args)
	if _len < 3 {
		fmt.Println("less Arguments")
	} else if _len == 3 {
		if isAnagram(os.Args[1], os.Args[2]) {
			fmt.Println("Anagram")
		} else {
			fmt.Println("Not Anagram")
		}
	} else {
		fmt.Println("too much Arguments")
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func isAnagram(s1 string, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)

	if len1 == len2 {
		for i := 0; i < len1; i++ {
			if !strings.Contains(s1, string(s2[i])) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

//!-
