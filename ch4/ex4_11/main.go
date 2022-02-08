package main

import (
	"ch4/ex4_11/github"
)

func main() {
	/*	issue1 := new(github.Issue)
		issue1.Title = "Additional Issue"
		issue1.Body = "This shall be solved"

		issue1_json, _ := json.Marshal(issue1)
		fmt.Printf("%s\n", string(issue1_json))
		github.CreateIssue(bytes.NewBuffer(issue1_json))*/
	github.ReadIssues()
}
