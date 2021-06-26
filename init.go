package main

import (
	"os"
)

func pickFile() string {
	home, _ := os.UserHomeDir()
	_ = os.Mkdir(home+"/.boffish", 0755)

	room := pickRoom(home + "/.boffish/")
	return home + "/.boffish/" + room
}

// 20-6 Roomstrap function takes a prompt in; checks if room exists, makes room with that name if not after user being prompted.
func pickRoom(directory string) string {
	searching := true
	room := ""

	for searching {
		room = prompt("What room are you using?")

		file, _ := os.Open(directory + room)

		if file != nil {
			return room
		} else {
			response := prompt("This room does not exist. Do you want to create it [1] or try again [2]?")

			if response == "1" {
				return room
			}
		}
	}
	return room // 20-6 Software does not compilation without this queue
}

