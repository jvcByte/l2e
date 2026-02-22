package main

import (
	"regexp"
	"strings"
)

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func normalizeModifiers(text string) string {
	re := regexp.MustCompile(`\((up|low|cap|hex|bin),\s*(\d+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		// remove spaces inside parentheses
		inner := regexp.MustCompile(`\s+`).ReplaceAllString(match, "")
		return inner
	})
}