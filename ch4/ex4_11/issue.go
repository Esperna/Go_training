package main

import (
	"ch4/ex4_11/github"
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := len(os.Args)
	if length < 2 || length > 8 {
		fmt.Println("Invalid Number of Argument.")
		fmt.Println("One of followings is expected")
		fmt.Println("./issue -c GitHubID Token Title Body Labels")
		fmt.Println("./issue -r")
		fmt.Println("./issue -u IssueNo GitHubID Token Title Body Labels")
		fmt.Println("./issue -uc IssueNo GitHubID BodyOfComment")
		fmt.Println("./issue -d IssueNo GitHubID Token")
		os.Exit(1)
	}
	if os.Args[1] == "-c" {
		if length == 7 {
			json_str := "{" + strconv.Quote("title") + ":" + strconv.Quote(os.Args[4])
			json_str += "," + strconv.Quote("body") + ":" + strconv.Quote(os.Args[5])
			json_str += "," + strconv.Quote("labels") + ":" + "[" + strconv.Quote(os.Args[6]) + "]"
			json_str += "}"
			fmt.Printf("%s\n", json_str)
			github.CreateIssue(os.Args[2], os.Args[3], json_str)
		} else {
			fmt.Println("Invalid Number of Argument.")
			fmt.Println("./issue -c GitHubID Token Title Body Labels")
			os.Exit(1)
		}
	} else if os.Args[1] == "-r" {
		github.ReadIssues()
	} else if os.Args[1] == "-u" {
		if length == 8 {
			json_str := "{" + strconv.Quote("title") + ":" + strconv.Quote(os.Args[5])
			json_str += "," + strconv.Quote("body") + ":" + strconv.Quote(os.Args[6])
			json_str += "," + strconv.Quote("labels") + ":" + "[" + strconv.Quote(os.Args[7]) + "]"
			json_str += "}"
			fmt.Printf("%s\n", json_str)
			github.UpdateIssue(os.Args[2], os.Args[3], os.Args[4], json_str)
		} else {
			fmt.Println("Invalid Number of Argument.")
			fmt.Println("./issue -u IssueNo GitHubID Token Title Body Labels")
			os.Exit(1)
		}
	} else if os.Args[1] == "-uc" {
		if length == 6 {
			json_str := "{" + strconv.Quote("body") + ":" + strconv.Quote(os.Args[5])
			json_str += "}"
			fmt.Printf("%s\n", json_str)
			//github.UpdateIssue(os.Args[2], os.Args[3], os.Args[4], json_str)
		} else {
			fmt.Println("Invalid Number of Argument.")
			fmt.Println("./issue -uc IssueNo GitHubID BodyOfComment")
			os.Exit(1)
		}
	} else if os.Args[1] == "-d" {
		if length == 5 {
			github.CloseIssue(os.Args[2], os.Args[3], os.Args[4])
		} else {
			fmt.Println("Invalid Number of Argument.")
			fmt.Println("./issue -d IssueNo GitHubID Token")
			os.Exit(1)
		}
	} else {
		fmt.Println("Invalid Option.")
		fmt.Println("./issue -c GitHubID Token Title Body Labels")
		fmt.Println("./issue -r")
		fmt.Println("./issue -u IssueNo GitHubID Token Title Body Labels")
		fmt.Println("./issue -uc IssueNo GitHubID BodyOfComment")
		fmt.Println("./issue -d IssueNo GitHubID Token")
		os.Exit(1)
	}

}
