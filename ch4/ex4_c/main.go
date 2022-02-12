package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}
