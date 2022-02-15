package main

import (
	"ch4/ex4_11a/github"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	length := len(os.Args)
	if length != 3 {
		fmt.Fprint(os.Stderr, fmt.Sprintf("invalid number of args. issue issueNo editor is expected \n"))
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[2], "tmp.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("failed run command. %s\n", err.Error()))
	}

	content, err := ioutil.ReadFile("tmp.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("failed read content. %s\n", err.Error()))
		os.Exit(1)
	}

	jsonStr := string(content)
	fmt.Printf("%s\n", jsonStr)

	err = godotenv.Load("github.env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		os.Exit(1)
	}

	issueNo := os.Args[1]
	GitHubID := os.Getenv("GITHUB_ID")
	Token := os.Getenv("GITHUB_TOKEN")
	github.UpdateIssue(issueNo, GitHubID, Token, jsonStr)

}
