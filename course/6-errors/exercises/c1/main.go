package main

import "fmt"

type StatusError struct {
	msg string
}

func (sE StatusError) Error() string {
	switch l := len(sE.msg); {
	case l <= 0:
		return fmt.Sprint("status cannot be empty")
	case l > 140:
		return fmt.Sprint("status exceeds 140 characters")
	default:
		return ""
	}
}

func validateStatus(status string) error {
	return StatusError{msg: status}
}
