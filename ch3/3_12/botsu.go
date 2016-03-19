package main

import (
	"fmt"
	"os"
	"sort"
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
		str1 := make([]string, 30)
		str2 := make([]string, 30)

		for i := 0; i < len1; i++ {
			str1[i] = string(s1[i])
		}
		sort.Strings(str1)
		fmt.Println(str1[0])
		for i := 0; i < len2; i++ {
			str2[i] = string(s2[i])
		}
		sort.Strings(str2)

		for i := 0; i < len1; i++ {
			if str1[i] != str2[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

//!-
