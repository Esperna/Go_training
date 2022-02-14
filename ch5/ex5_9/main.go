package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	length := len(os.Args)
	if length != 2 {
		fmt.Fprintf(os.Stderr, "expand failed: invalid num of args ")
		os.Exit(1)
	}
	f := Length
	fmt.Println(expand(os.Args[1], f))
}

func Length(s string) string {
	return strconv.Itoa(len(s))
}

func expand(s string, f func(string) string) string {
	slices := strings.Split(s, "$")
	var str string
	for _, item := range slices {
		str += f(item)
	}
	if HasPrefix(s, "$") {
		return str[1:]
	} else {
		return slices[0] + str[1:]
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
