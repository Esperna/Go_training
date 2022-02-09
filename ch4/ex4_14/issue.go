// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ch4/ex4_14/github"
)

//!+
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%d issues:\n", result.TotalCount)

	fmt.Fprintln(w, "\n<Bug List>")
	milestones := make(map[string]string)
	userNames := make(map[string]string)
	for _, item := range result.Items {
		fmt.Fprintf(w, "#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		if item.Milestone != nil {
			milestones[item.Milestone.Title] = item.Milestone.Description
		}
		if item.User != nil {
			userNames[item.User.Login] = item.User.HTMLURL
		}
	}
	fmt.Fprintln(w, "\n<Milestone List>")
	for k, v := range milestones {
		fmt.Fprintf(w, "%s\t%s\n", k, v)
	}
	fmt.Fprintln(w, "\n<User List>")
	for k, v := range userNames {
		fmt.Fprintf(w, "%s\t%s\n", k, v)
	}

}
