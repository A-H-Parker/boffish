package main

import (
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

var colorRed = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
var colorOrange = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF7F00"))
var colorYellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
var colorGreen = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
var colorBlue = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
var colorPurple = lipgloss.NewStyle().Foreground(lipgloss.Color("#4B0082"))

func rainbow(input string) string {
	var output string

	for x := 0; x < utf8.RuneCountInString(input); x++ {
		switch x % 6 {
		case 0:
			output = output + red(string([]rune(input)[x]))
		case 1:
			output = output + orange(string([]rune(input)[x]))
		case 2:
			output = output + yellow(string([]rune(input)[x]))
		case 3:
			output = output + green(string([]rune(input)[x]))
		case 4:
			output = output + blue(string([]rune(input)[x]))
		case 5:
			output = output + purple(string([]rune(input)[x]))
		}
	}
	return output
}
func red(input string) string {
	return colorRed.Render(input)
}

func orange(input string) string {
	return colorOrange.Render(input)
}

func yellow(input string) string {
	return colorYellow.Render(input)
}

func green(input string) string {
	return colorGreen.Render(input)
}

func blue(input string) string {
	return colorBlue.Render(input)
}

func purple(input string) string {
	return colorPurple.Render(input)
}

