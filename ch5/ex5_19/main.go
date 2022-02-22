package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() (greeting string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			err = fmt.Errorf("within expected")
			greeting = "Hello."
		default:
			panic(p)
		}
	}()
	panic(nil)
}
