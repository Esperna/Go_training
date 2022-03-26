package main

import (
	"fmt"
	"os"
	"regexp"
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
	//return strings.ToUpper(s)
}

func expand(s string, f func(string) string) string {
	rep := regexp.MustCompile(`\$[A-Za-z]*`)
	stripDollar := func(s string) string {
		s = strings.Replace(s, "$", "", 1)
		return f(s)
	}
	s = rep.ReplaceAllStringFunc(s, stripDollar)
	return s
}
