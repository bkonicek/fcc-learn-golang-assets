package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	wordCount := make(map[string]int)
	for _, message := range messages {
		words := strings.Fields(message)
		for _, w := range words {
			word := strings.ToLower(w)
			wordCount[word]++
		}
	}
	return len(wordCount)
}
