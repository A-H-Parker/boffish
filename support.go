package main

import "fmt"

// 20-6 Created event for getting user input to make process simple
func prompt(question string) string {
	var answer string
	fmt.Println(question)
	fmt.Scan(&answer)
	return answer
}

