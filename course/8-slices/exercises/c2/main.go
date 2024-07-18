package main

import "unicode"

func isValidPassword(password string) bool {
	containsUpper := false
	containsDigit := false
	length := 0
	for _, char := range password {
		switch v := char; {
		case unicode.IsUpper(v):
			containsUpper = true
		case unicode.IsDigit(v):
			containsDigit = true
		}
		length++
	}
	if length < 5 || length > 12 {
		return false
	}
	if !containsDigit || !containsUpper {
		return false
	}
	return true
}
