package main

import (
	"fmt"
	"os"
	"strings"
)

// This command determines if a contribution is a command, and does not log it if so, unless it is a /plain command.
func commandProcessing(message string) (bool, string) {
	switch strings.Split(message, " ")[0] {
	case "/plain":
		return true, strings.TrimPrefix(message, "/plain ")
	case "/exit", "/quit":
		exit()
		return false, ""
	case "/rainbow":
		return true, rainbow(strings.TrimPrefix(message, "/rainbow "))
	default:
		return false, ""
	}
}

func exit() {
	var response string
	fmt.Println("Are you sure you want to exit? <Y/N> ")
	fmt.Scanln(&response)

	if response == "Y" {
		os.Exit(0)
	}
}

