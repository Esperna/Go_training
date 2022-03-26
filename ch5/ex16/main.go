package main

import (
	"fmt"
)

func main() {
	fmt.Println(Join(" ", "1"))
	fmt.Println(Join(" ", "1", "2"))
	fmt.Println(Join(" ", "abc", "cde", "fgh"))

}

func Join(sep string, elems ...string) string {
	sum := ""
	for i, v := range elems {
		sum += v
		if i < len(elems)-1 {
			sum += sep
		}
	}
	return sum
}
