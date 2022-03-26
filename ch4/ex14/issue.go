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

	"ch4/ex14/github"
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
`))

var milestoneList = template.Must(template.New("milestonelist").Parse(`
<h1>Milestones</h1>
<table>
<tr>
	<th>Title</th>
	<th>Description</th>
</tr>
{{range .Items}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Description}}</td>
</tr>
{{end}}
</table>
`))

var userList = template.Must(template.New("userlist").Parse(`
<h1>Users</h1>
<table>
<tr>
	<th>Login</th>
	<th>HTMLURL</th>
</tr>
{{range .Items}}
<tr>
	<td>{{.Login}}</td>
	<td>{{.HTMLURL}}</td>
</tr>
{{end}}
</table>
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
	issueResult, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(w, issueResult); err != nil {
		log.Fatal(err)
	}

	milestones := make(map[string]string)
	users := make(map[string]string)
	for _, item := range issueResult.Items {
		if item.Milestone != nil {
			milestones[item.Milestone.Title] = item.Milestone.Description
		}
		if item.User != nil {
			users[item.User.Login] = item.User.HTMLURL
		}
	}
	var milestoneResult github.MilestonesSearchResult
	milestones1 := make([]*github.Milestone, 0)
	for k, v := range milestones {
		var milestone1 github.Milestone
		if k != "" && v != "" {
			milestone1.Title = k
			milestone1.Description = v
			milestones1 = append(milestones1, &milestone1)
		}
	}
	milestoneResult.Items = milestones1

	var userResult github.UsersSearchResult
	users1 := make([]*github.User, 0)
	for k, v := range users {
		var user1 github.User
		if k != "" && v != "" {
			user1.Login = k
			user1.HTMLURL = v
			users1 = append(users1, &user1)
		}
	}
	userResult.Items = users1

	if err := milestoneList.Execute(w, milestoneResult); err != nil {
		log.Fatal(err)
	}
	if err := userList.Execute(w, userResult); err != nil {
		log.Fatal(err)
	}
}
