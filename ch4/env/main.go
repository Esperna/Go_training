package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("github.env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		os.Exit(1)
	}
	length := len(os.Args)
	if length != 3 {
		fmt.Fprintf(os.Stderr, "invalid number of args")
		os.Exit(1)
	}

	//os.Setenv does not change .env file itself
	err = os.Setenv("GITHUB_ID", os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github id error %s", err.Error())
	}
	err = os.Setenv("GITHUB_TOKEN", os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github token error %s", err.Error())
	}

	fmt.Printf("%s\n%s\n", os.Getenv("GITHUB_ID"), os.Getenv("GITHUB_TOKEN"))
}
