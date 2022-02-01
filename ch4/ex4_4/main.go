package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5}
	x = rotate(x, 2)
	fmt.Println(x)
}

func rotate(s []int, i int) []int {
	return append(s[i:len(s)], s[0:i]...)
}
