package main

import "fmt"

type StatusError struct {
	msg string
}

func (sE StatusError) Error() string {
	return fmt.Sprint(sE.msg)
}

func validateStatus(status string) error {
	if len(status) <= 0 {
		return StatusError{msg: "status cannot be empty"}
	}
	if len(status) > 140 {
		return StatusError{msg: "status exceeds 140 characters"}
	}
	return nil
}
