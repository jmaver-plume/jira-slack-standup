package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Fields FieldsObjs `json:"fields"`
	Key    string     `json:key`
}

type FieldsObjs struct {
	Summary string `json:"summary"`
}

func getIssue(key string) Response {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/api/2/issue/%s", os.Getenv("BASE_URL"), strings.Split(key, " ")[0]), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

func getIssues(keys []string) []Response {
	issues := make([]Response, len(keys))
	for i, key := range keys {
		issues[i] = getIssue(key)
		//issues[i] = fmt.Sprintf("%s: %s", issue.Key, issue.Fields.Summary)
	}
	return issues
	//return strings.Join(issues, "\n")
}

type flags struct {
	Username string
	Password string
	BaseUrl  string
	Keys     string
}

func getFlags() flags {
	usernamePtr := flag.String("username", "username", "Jira user username")
	passwordPtr := flag.String("password", "password", "Jira user password")
	baseUrlPtr := flag.String("baseUrl", "https://example.jira.com", "Jira base url")
	keyString := flag.String("keys", "JIRA-100", "List of jira tasks separated by newline character")

	flag.Parse()
	return flags{*usernamePtr, *passwordPtr, *baseUrlPtr, *keyString}
}

func setEnvVariables(f flags) {
	os.Setenv("USERNAME", f.Username)
	os.Setenv("PASSWORD", f.Password)
	os.Setenv("BASE_URL", f.BaseUrl)
}

func handleKeyFlags(f string) string {
	result := ""
	split := strings.Split(f, "")
	for _, v := range split {
		if v == "j" {
			result += ":jira-new:"
		}
		if v == "c" {
			result += ":codereview:"
		}
		if v == "w" {
			result += ":white_check_mark:"
		}
		if v == "b" {
			result += ":building:"
		}
		if v == "e" {
			result += ":eyes:"
		}
		if v == "r" {
			result += ":brain:"
		}
	}

	return result
}

func getResult(issues []Response, rawKeys []string) string {
	result := make([]string, len(issues))
	for i, issue := range issues {
		f := handleKeyFlags(strings.Split(rawKeys[i], " ")[1])
		result[i] = fmt.Sprintf("%s %s: %s", f, issue.Key, issue.Fields.Summary)
	}
	return strings.Join(result, "\n")
}

func main() {
	f := getFlags()
	setEnvVariables(f)
	keys := strings.Split(f.Keys, "\n")
	issues := getIssues(keys)
	result := getResult(issues, keys)
	err := clipboard.WriteAll(result)
	if err != nil {
		log.Fatal(err)
	}
}
