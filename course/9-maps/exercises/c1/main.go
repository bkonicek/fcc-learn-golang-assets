package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	wordCount := make(map[string]struct{})
	for _, message := range messages {
		for _, word := range strings.Fields(strings.ToLower(message)) {
			wordCount[word] = struct{}{}
		}
	}
	return len(wordCount)
}
