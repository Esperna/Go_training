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
	if !(length == 3 || length == 4) {
		fmt.Fprint(os.Stderr, fmt.Sprintf("invalid number of args. \n"))
		displayHowToUse()
		os.Exit(1)
	}
	if !(os.Args[1] == "-c" || os.Args[1] == "-r" || os.Args[1] == "-u" || os.Args[1] == "-d" || os.Args[1] == "-uc") {
		fmt.Fprint(os.Stderr, fmt.Sprintf("invalid option. \n"))
		displayHowToUse()
		os.Exit(1)
	}

	err := godotenv.Load("github.env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		os.Exit(1)
	}

	GitHubID := os.Getenv("GITHUB_ID")
	Token := os.Getenv("GITHUB_TOKEN")

	if os.Args[1] == "-c" {

	} else if os.Args[1] == "-r" {

	} else if os.Args[1] == "-u" {
		jsonStr, err := inputFromEditor(os.Args[3])
		if err != nil {
			fmt.Fprint(os.Stderr, "input from editor failed")
			os.Exit(1)
		}
		issueNo := os.Args[2]
		github.UpdateIssue(issueNo, GitHubID, Token, jsonStr)

	} else if os.Args[1] == "-d" {

	}
}

func inputFromEditor(editor string) (string, error) {
	cmd := exec.Command(editor, "tmp.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("failed run command. %s\n", err.Error()))
		return "", err
	}
	content, err := ioutil.ReadFile("tmp.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("failed read content. %s\n", err.Error()))
		return string(content), err
	}
	return string(content), nil
}

func displayHowToUse() {
	fmt.Println("./issue -c editor")
	fmt.Println("./issue -r")
	fmt.Println("./issue -u IssueNo editor")
	fmt.Println("./issue -d IssueNo")
}
