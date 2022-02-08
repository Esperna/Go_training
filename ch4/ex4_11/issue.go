package main

import (
	"ch4/ex4_11/github"
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := len(os.Args)
	if length < 1 || length > 7 {
		fmt.Println("Invalid Number of Argument. ./issue -option(-c,-r,-u, or -d) GitHubID Token Title Body Labels is expected")
		os.Exit(1)
	} else {
		if os.Args[1] == "-c" {
			if length == 7 {
				json_str := "{" + strconv.Quote("title") + ":" + strconv.Quote(os.Args[4])
				json_str += "," + strconv.Quote("body") + ":" + strconv.Quote(os.Args[5])
				json_str += "," + strconv.Quote("labels") + ":" + "[" + strconv.Quote(os.Args[6]) + "]"
				json_str += "}"
				fmt.Printf("%s\n", json_str)
				github.CreateIssue(os.Args[2], os.Args[3], json_str)
				github.ReadIssues()
			}
		} else if os.Args[1] == "-r" {
			github.ReadIssues()
		} else if os.Args[1] == "-u" {
			fmt.Printf("TBD\n")
		} else if os.Args[1] == "-d" {
			fmt.Printf("TBD\n")
		} else {
			fmt.Println("Invalid Option. ./issue -option(-c,-r,-u, or -d) GitHubID Token Title Body Labels is expected")
			os.Exit(1)
		}
	}
}
