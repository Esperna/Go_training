package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() (greeting string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			//no panic
		case bailout{}:
			err = fmt.Errorf("within expected")
			greeting = "Hello."
		default:
			panic(p)
		}
	}()
	panic(bailout{})
}
