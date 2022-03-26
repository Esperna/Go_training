// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type UsersSearchResult struct {
	Items []*User
}

type MilestonesSearchResult struct {
	Items []*Milestone
}

type Issue struct {
	HTMLURL   string `json:"html_url"`
	Number    int
	Title     string
	User      *User
	State     string
	Body      string // in Markdown format
	Milestone *Milestone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title       string
	Description string
}

//!-
