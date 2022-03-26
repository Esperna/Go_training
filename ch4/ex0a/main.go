package main

import (
	"ch4/ex0a/treesort"
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := len(os.Args)
	if length > 1 {
		var slice []int
		for i := 1; i < length; i++ {
			value, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("Invalid String")
				return
			}
			slice = append(slice, value)
		}
		treesort.Sort(slice)
		for _, v := range slice {
			fmt.Println(v)
		}
	} else {
		fmt.Println("Invalid Number of Argument")
	}

}
