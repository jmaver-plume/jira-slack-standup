package lineParser

import (
	"fmt"
)

func ExampleLine_GetEnrichedFlags() {
	l := Line{
		Key:   "",
		Flags: []string{"j", "c", "w", "b", "e", "r", "m"},
	}
	fmt.Println(l.GetEnrichedFlags())
	// Output: :jira-new::codereview::white_check_mark::building::eyes::brain::merge:
}

func ExampleParseLine() {
	fmt.Println(ParseLine("JIRA-1 jb"))
	// Output: {JIRA-1 [j b]}
}
