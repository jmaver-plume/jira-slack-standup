// Package lineParser provides utility functions for line parsing
package lineParser

import "strings"

type Line struct {
	Key   string
	Flags []string
}

var EnrichFlagMap = map[string]string{
	"j": ":jira-new:",
	"c": ":codereview:",
	"w": ":white_check_mark:",
	"b": ":building:",
	"e": ":eyes:",
	"r": ":brain:",
	"m": ":merge:",
}

// GetEnrichedFlags returns a mapping of each flag
func (l *Line) GetEnrichedFlags() string {
	result := ""
	for _, v := range l.Flags {
		m, ok := EnrichFlagMap[v]
		if !ok {
			continue
		}
		result += m
	}

	return result
}

func ParseLine(line string) Line {
	split := strings.Split(line, " ")

	var flags []string
	// no flags were passed
	if len(split) == 2 {
		flags = strings.Split(split[1], "")
	}

	return Line{Key: split[0], Flags: flags}
}

func ParseLines(s []string) []Line {
	var lines []Line
	for _, s2 := range s {
		lines = append(lines, ParseLine(s2))
	}
	return lines
}
