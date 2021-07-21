// Package jira is a wrapper around Client REST API v2.
package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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

type Client struct {
	Username string
	Password string
	BaseUrl  string
}

type issueResponse struct {
	Fields issueFieldsResponse `json:"fields"`
	Key    string              `json:"key"`
}

type issueFieldsResponse struct {
	Summary string `json:"summary"`
}

type Issue interface {
	Get(string) issueResponse
	List([]string) []issueResponse
	Search(string) []issueResponse
}

func NewClient(username, password, baseUrl string) *Client {
	return &Client{
		Username: username,
		Password: password,
		BaseUrl:  baseUrl,
	}
}

// GetIssue returns issue by Client key as specified in https://docs.atlassian.com/software/jira/docs/api/REST/7.6.1/#api/2/issue-getIssue
func (j *Client) GetIssue(key string) issueResponse {
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

	var responseObject issueResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

// GetIssues calls GetIssue for issue item in keys
func (j *Client) GetIssues(keys []string) []issueResponse {
	var issues []issueResponse
	for _, key := range keys {
		issues = append(issues, j.GetIssue(key))
	}
	return issues
}

// SearchIssue returns a list of issues as specified in https://confluence.atlassian.com/jirakb/how-to-parse-access-log-in-jira-for-audit-purposes-1004934149.html
func (j Client) SearchIssue(query string) []SearchIssueResponse {
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

	var responseObject []SearchIssueResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
