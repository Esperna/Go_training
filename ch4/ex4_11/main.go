package main

import (
	"ch4/ex4_11/github"
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := len(os.Args)
	if length == 6 {
		json_str := "{" + strconv.Quote("title") + ":" + strconv.Quote(os.Args[3])
		json_str += "," + strconv.Quote("body") + ":" + strconv.Quote(os.Args[4])
		json_str += "," + strconv.Quote("labels") + ":" + "[" + strconv.Quote(os.Args[5]) + "]"
		json_str += "}"
		fmt.Printf("%s\n", json_str)
		github.CreateIssue(os.Args[1], os.Args[2], json_str)
		github.ReadIssues()
	} else {
		fmt.Println("Invalid Number of Argument. ./main ID Token Title Body Labels is expected")
	}
}
