// Package jira is a wrapper around Jira REST API v2.
package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Issue struct {
	Fields Fields `json:"fields"`
	Key    string `json:"key"`
}

type SearchIssueResponse struct {
	Id           string                    `json:"id"`
	Name         string                    `json:"name"`
	ViewAllTitle string                    `json:"viewAllTitle"`
	Items        []SearchIssueItemResponse `json:"items"`
}

type SearchIssueItemResponse struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	AvatarUrl string `json:"avatarUrl"`
	Url       string `json:"url"`
	Favorite  bool   `json:"favorite"`
}

type Fields struct {
	Summary string `json:"summary"`
}

type Jira struct {
	Username string
	Password string
	BaseUrl  string
}

type IssueInterface interface {
	GetIssue(string) Issue
	GetIssues([]string) []Issue
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

func (j Jira) SearchIssue(query string) SearchIssueResponse {
	escapedQuery := url.QueryEscape(query)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/quicksearch/1.0/productsearch/search?q=%s", j.BaseUrl, escapedQuery), nil)
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

	var responseObject SearchIssueResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
