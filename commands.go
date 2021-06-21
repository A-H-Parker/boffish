package main

import (
	"os"
	"strings"
)

// This command determines if a contribution is a command, and does not log it if so, unless it is a /plain command.
func commandProcessing(message string) bool {
	if strings.HasPrefix(message, "/") {
		if strings.HasPrefix(message, "/plain ") {
			return true
		} else {
			if message == "/exit" {
				os.Exit(0)
				return false
			} else {
				return false
			}
		}
	} else {
		return true
	}
}

