package lineParser

import "strings"

type Line struct {
	Key   string
	Flags []string
}

func (l *Line) GetEnrichedFlags() string {
	result := ""
	for _, v := range l.Flags {
		switch v {
		case "j":
			result += ":jira-new:"
		case "c":
			result += ":codereview:"
		case "w":
			result += ":white_check_mark:"
		case "b":
			result += ":building:"
		case "e":
			result += ":eyes:"
		case "r":
			result += ":brain:"

		}
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
