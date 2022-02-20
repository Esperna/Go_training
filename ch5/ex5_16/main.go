package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Join())
	fmt.Println(Join("1", " "))
	fmt.Println(Join("1", "2", " "))
	fmt.Println(Join("abc", "cde", "fgh", " "))

}

func Join(elems ...string) string {
	length := len(elems)
	if length < 1 {
		fmt.Fprintf(os.Stderr, "Invalid number of argument")
		return ""
	}
	sep := elems[len(elems)-1]
	sum := ""
	for i, v := range elems {
		sum += v
		if i < len(elems)-1 {
			sum += sep
		}
	}
	return sum
}
