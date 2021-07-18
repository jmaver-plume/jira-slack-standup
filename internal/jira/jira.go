// Package jira is a wrapper around Jira REST API v2.
package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Issue struct {
	Fields Fields `json:"fields"`
	Key    string `json:"key"`
}

type Fields struct {
	Summary string `json:"summary"`
}

type Jira struct {
	Username string
	Password string
	BaseUrl  string
}

// GetIssue returns issue by Jira key as specified in https://docs.atlassian.com/software/jira/docs/api/REST/7.6.1/#api/2/issue-getIssue
func (j *Jira) GetIssue(key string) Issue {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/api/2/issue/%s", j.BaseUrl, key), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(j.Username, j.Password)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Issue
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

// GetIssues calls GetIssue for issue item in keys
func (j *Jira) GetIssues(keys []string) []Issue {
	var issues []Issue
	for _, key := range keys {
		issues = append(issues, j.GetIssue(key))
	}
	return issues
}
