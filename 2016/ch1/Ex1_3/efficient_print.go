package main

import(
	"fmt"
	"os"
	"strings"
)

func EfficientPrintln() {
	for i := 1; i < 500; i++ {
		os.Args[i] = i
	}	
	fmt.Println(strings.Join(os.Args[1:], " "))
}
