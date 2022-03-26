package main

import "fmt"

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(ptr *[6]int) {
	for i := 0; i < 5-i; i++ {
		tmp := ptr[i]
		ptr[i] = ptr[5-i]
		ptr[5-i] = tmp
	}
}
