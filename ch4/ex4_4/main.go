package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x []int
	length := len(os.Args)
	if length > 4 {
		for i := 1; i < length; i++ {
			if os.Args[i] != "-r" {
				j, err := strconv.Atoi(os.Args[i])
				if err != nil {
					fmt.Println("Invalid Argument.")
					return
				}
				x = append(x, j)
			} else {
				if i+1 < length {
					r, err := strconv.Atoi(os.Args[i+1])
					if err != nil {
						fmt.Println("Invalid Argument.")
						return
					}
					x = rotate(x, r)
					fmt.Println(x)
				} else {
					fmt.Println("Invalid Argument.")
				}
			}
		}
	} else {
		fmt.Println("Invalid Number of Argument. More than 3(value1 value2 -r rotateValue) is expected.")
	}
}

func rotate(s []int, i int) []int {
	length := len(s)
	if i < length {
		return append(s[i:length], s[0:i]...)
	} else {
		return rotate(s, i-length)
	}
}
