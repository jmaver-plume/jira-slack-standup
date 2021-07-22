package lineParser

import (
	"fmt"
)

func ExampleLine_GetEnrichedFlags() {
	l := Line{
		Key:   "",
		Flags: []string{"j", "c", "w", "b", "e", "r"},
	}
	fmt.Println(l.GetEnrichedFlags())
	// Output: :jira-new::codereview::white_check_mark::building::eyes::brain:
}

func ExampleParseLine() {
	fmt.Println(ParseLine("JIRA-1 jb"))
	// Output: {JIRA-1 [j b]}
}
