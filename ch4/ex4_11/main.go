package main

import (
	"ch4/ex4_11/github"
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length == 3 {
		/*	issue1 := new(github.Issue)
			issue1.Title = "Additional Issue"
			issue1.Body = "This shall be solved"

			issue1_json, _ := json.Marshal(issue1)
			fmt.Printf("%s\n", string(issue1_json))
			github.CreateIssue(bytes.NewBuffer(issue1_json))
		*/
		params := `{"title": "test2"}`
		github.CreateIssue(os.Args[1], os.Args[2], params)
		github.ReadIssues()
	} else {
		fmt.Println("Invalid Number of Argument. ./main ID Token is expected")
	}
}
