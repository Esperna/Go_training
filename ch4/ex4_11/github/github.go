// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "encoding/json"

const IssuesURL = "https://api.github.com/repos/Esper0328/Go_training/issues"

type IssuesReadResult struct {
	Url                   string
	RepositoryUrl         string `json:"repository_url"`
	LabelsUrl             string `json:"labels_url"`
	CommentsUrl           string `json:"comments_url"`
	EventsUrl             string `json:"events_url"`
	HtmlUrl               string `json:"html_url"`
	Id                    int
	NodeId                string `json:"node_id"`
	Number                int
	Title                 string
	User                  User
	Labels                []Label
	State                 string
	Locked                bool
	Assignee              string
	Assignees             json.RawMessage
	Milestone             string
	Comments              int
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	ClosedAt              string `json:"closed_at"`
	AuthorAssociation     string `json:"author_association"`
	ActiveLockReason      string `json:"active_lock_reason"`
	Body                  string
	Reactions             Reaction
	TimelineUrl           string `json:"timeline_url"`
	PerformedViaGithubApp string `json:"performed_via_github_app"`
}

type User struct {
	Login             string
	Id                int
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"Following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	Repos_url         string `json:"Repos_url"`
	Eventsurl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string
	SiteAdmin         bool `json:"site_admin"`
}

type Label struct {
	Id          int
	Node_id     string
	Url         string
	Name        string
	Color       string
	Default     bool
	Description string
}

type Reaction struct {
	Url         string
	Total_count int
	PosiValue   int
	NegaValue   int
	Laugh       int
	Hooray      int
	Confused    int
	Heart       int
	Rocket      int
	Eyes        int
}

//!-
