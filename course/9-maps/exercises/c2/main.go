package main

import "slices"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggested := []string{}
	for _, friend := range friendships[username] {
		for _, v := range friendships[friend] {
			if v != username && v != friend {
				if !slices.Contains(suggested, v) {
					suggested = append(suggested, v)
				}
			}
		}
	}
	if len(suggested) > 0 {
		return suggested
	}
	return nil
}
