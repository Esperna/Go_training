// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchIssues queries the GitHub issue tracker.
func ReadIssues() {
	resp, err := http.Get(IssuesURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var results []IssuesReadResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		fmt.Printf("%v\n", err)
		resp.Body.Close()
		return
	}
	for _, result := range results {
		fmt.Printf("#%d\t%s\t%s\t%s\t%s\n", result.Number, result.Title,
			result.User.Login, result.State, result.Body)
	}
	resp.Body.Close()
}

//!-
