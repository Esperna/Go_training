package main

import "fmt"

func main() {

	s := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 4, 4}
	fmt.Println(removeDuplication(s))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeDuplication(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			return removeDuplication(remove(slice, i))
		}
	}
	return slice[:]
}
