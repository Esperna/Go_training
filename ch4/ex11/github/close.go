// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

// SearchIssues queries the GitHub issue tracker.
func CloseIssue(number, id, token string) {
	json_str := "{" + strconv.Quote("state") + ":" + strconv.Quote("closed")
	json_str += "}"
	fmt.Printf("%s\n", json_str)

	req, err := http.NewRequest("POST", IssuesURL+"/"+number, bytes.NewBuffer([]byte(json_str)))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/vnd.github.v3+json")
	req.SetBasicAuth(id, token)
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")

	client := new(http.Client)
	resp, err := client.Do(req)

	fmt.Printf("Status:%s\n", resp.Status)
	fmt.Printf("Body\n%v\n", resp.Body)
	resp.Body.Close()
}

//!-
