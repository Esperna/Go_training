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

func rotate(s []int, k int) []int {
	n := len(s)
	l := gcd(k, n)
	m := n / l
	for i := 0; i < l; i++ {
		var buf1, buf2 int
		for j := 0; j < m; j++ {
			if j == 0 {
				buf1 = s[(i+k*(j+1))%n]
				s[(i+k*(j+1))%n] = s[(i+k*j)%n]
			} else {
				buf2 = s[(i+k*(j+1))%n]
				s[(i+k*(j+1))%n] = buf1
				buf1 = buf2
			}
		}
	}
	return s
}

func gcd(a, b int) int {
	if a > b {
		temp := a
		a = b
		b = temp
	}
	for {
		c := b % a
		if c == 0 {
			return a
		}
		if a > c {
			b = a
			a = c
		} else {
			b = c
		}
	}
}
