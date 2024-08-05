package main

import "fmt"

func processMessages(messages []string) []string {
	ch := make(chan string, len(messages))
	go func() {
		for _, message := range messages {
			ch <- message
		}
		close(ch)
	}()

	var processed []string
	for i := 0; i < len(messages); i++ {
		processed = append(processed, fmt.Sprintf("%s-processed", <-ch))
	}
	return processed
}
