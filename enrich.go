package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"jiraEnrich/internal/jira"
	"jiraEnrich/internal/lineParser"
	"log"
	"strings"
	"time"
)

type flags struct {
	Username string
	Password string
	BaseUrl  string
	Lines    string
}

func getFlags() flags {
	usernamePtr := flag.String("username", "username", "Jira user username")
	passwordPtr := flag.String("password", "password", "Jira user password")
	baseUrlPtr := flag.String("baseUrl", "https://example.jira.com", "Jira base url")
	keyString := flag.String("keys", "JIRA-100", "List of jira tasks separated by newline character")

	flag.Parse()
	return flags{*usernamePtr, *passwordPtr, *baseUrlPtr, *keyString}
}

func getResult(issues []jira.Issue, enrichedFlags []string) string {
	var result []string
	result = append(result, getDay())
	for i, issue := range issues {
		result = append(result, fmt.Sprintf("%s %s: %s", enrichedFlags[i], issue.Key, issue.Fields.Summary))
	}
	return strings.Join(result, "\n")
}

func getKeys(lines []lineParser.Line) []string {
	var keys []string
	for _, line := range lines {
		keys = append(keys, line.Key)
	}

	return keys
}

func getEnrichedFlags(lines []lineParser.Line) []string {
	var keys []string
	for _, line := range lines {
		keys = append(keys, line.GetEnrichedFlags())
	}

	return keys
}

func getDay() string {
	return fmt.Sprintf("`%s`", time.Now().Weekday().String())
}

func main() {
	f := getFlags()
	j := jira.Jira{BaseUrl: f.BaseUrl, Username: f.Username, Password: f.Password}

	lines := strings.Split(f.Lines, "\n")
	parsedLines := lineParser.ParseLines(lines)

	keys := getKeys(parsedLines)
	enrichedFlags := getEnrichedFlags(parsedLines)

	issues := j.GetIssues(keys)
	result := getResult(issues, enrichedFlags)
	err := clipboard.WriteAll(result)
	if err != nil {
		log.Fatal(err)
	}
}
