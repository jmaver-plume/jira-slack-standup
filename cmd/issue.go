package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"jiraEnrich/internal/jira"
	"jiraEnrich/internal/lineParser"
	"log"
	"strings"
)

var Username string
var Password string
var BaseUrl string
var Alfred bool

type AlfredItemsOutput struct {
	Items []AlfredItem `json:"items"`
}

type AlfredItem struct {
	Uid          string         `json:"uid"`
	Type         string         `json:"type"`
	Title        string         `json:"title"`
	Subtitle     string         `json:"subtitle"`
	Arg          string         `json:"arg"`
	Autocomplete string         `json:"autocomplete"`
	Icon         AlfredItemIcon `json:"icon"`
}

type AlfredItemIcon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "", "Username with which to make API requests to Jira")
	err := rootCmd.MarkPersistentFlagRequired("username")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(&Password, "password", "p", "", "Password which belongs to user")
	err = rootCmd.MarkPersistentFlagRequired("password")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(&BaseUrl, "url", "b", "", "Base url to call")
	err = rootCmd.MarkPersistentFlagRequired("url")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().BoolVarP(&Alfred, "alfred", "a", true, "Output valid for Alfred workflows")

	rootCmd.AddCommand(issueCmd)
	issueCmd.AddCommand(issueGetCmd)
	issueCmd.AddCommand(issueSearchCmd)
	issueCmd.AddCommand(issueEnrichCmd)
}

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Commands for Jira issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var issueGetCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get Jira issue by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// jira.Issue.Get
		fmt.Printf("Jira - %+v", args)
	},
}

var issueSearchCmd = &cobra.Command{
	Use:   "search [pattern]",
	Short: "Search Jira issues by pattern",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := jira.NewClient(Username, Password, BaseUrl)
		query := args[0]
		result := client.SearchIssue(query)
		items := result[0].Items

		var alfredItems []AlfredItem
		for _, item := range items {
			alfredItems = append(alfredItems, AlfredItem{
				Uid:          item.Subtitle,
				Type:         "",
				Title:        item.Title,
				Subtitle:     item.Subtitle,
				Arg:          item.Subtitle,
				Autocomplete: item.Title,
				Icon: AlfredItemIcon{
					Type: item.AvatarUrl,
					Path: "",
				},
			})
		}

		output := AlfredItemsOutput{Items: alfredItems}
		j, err := json.Marshal(output)
		if err != nil {
			cmd.Println("Error")
			return
		}
		fmt.Println(string(j))
	},
}

var issueEnrichCmd = &cobra.Command{
	Use:   "enrich [keys]",
	Short: "Enrich a list of keys",
	Long:  "Enrich a list of keys, delimited by newline character, by calling GetIssue for each key and returning <key>: <summary> format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := jira.NewClient(Username, Password, BaseUrl)
		lines := lineParser.ParseLines(strings.Split(args[0], "\n"))
		var r []string
		for _, line := range lines {
			issue := client.GetIssue(line.Key)
			s := fmt.Sprintf("%s: %s", line.Key, issue.Fields.Summary)
			ef := line.GetEnrichedFlags()
			if ef != "" {
				s = fmt.Sprintf("%s %s", ef, s)
			}
			r = append(r, s)
		}
		fmt.Println(strings.Join(r, "\n"))
	},
}
