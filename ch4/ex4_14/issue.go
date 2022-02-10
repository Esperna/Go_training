// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ch4/ex4_14/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h1>Milestones</h1>
<h1>Users</h1>
`))

//!+
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/issue", issue)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func issue(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var q []string
	for _, v := range r.Form {
		q = v
	}
	if len(q) == 0 {
		return
	}
	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Fprintf(w, "%d issues:\n", result.TotalCount)
	//	fmt.Fprintln(w, "\n<Bug List>")
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
	milestones := make(map[string]string)
	userNames := make(map[string]string)
	for _, item := range result.Items {
		//		fmt.Fprintf(w, "#%-5d %9.9s %.55s\n",
		//			item.Number, item.User.Login, item.Title)
		if item.Milestone != nil {
			milestones[item.Milestone.Title] = item.Milestone.Description
		}
		if item.User != nil {
			userNames[item.User.Login] = item.User.HTMLURL
		}
	}
	/*
		fmt.Fprintln(w, "\n<Milestone List>")
		for k, v := range milestones {
					fmt.Fprintf(w, "%s\t%s\n", k, v)
		}
		fmt.Fprintln(w, "\n<User List>")
		for k, v := range userNames {
					fmt.Fprintf(w, "%s\t%s\n", k, v)
		}
	*/
}
