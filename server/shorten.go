package server

import "strings"

const (
	maxChars = 1024
	maxLines = 6
)

func Shorten(s string) string {
	if len(s) > maxChars {
		s = s[:maxChars]
	}

	lines := strings.Split(s, "\n")
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return strings.Join(lines, "\n")
}
