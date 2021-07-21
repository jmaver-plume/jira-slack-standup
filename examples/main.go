package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchQuery := searchCmd.String("query", "", "query by which to search issues")
	//issueFunc := issueCmd.String("")
	switch os.Args[1] {
	case "issue":
		switch os.Args[2] {
		case "enrich":
			fmt.Println("enrich")
		case "search":
			searchCmd.Parse(os.Args[3:])
			fmt.Println("searchQuery: ", *searchQuery)
		default:
			errors.New("invalid subcommand")
		}
	default:
		errors.New("invalid subcommand")
	}
}
