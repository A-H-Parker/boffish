package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	room := pickRoom()
	mainloop(room)
}

// The structure in which we will keep the message we are working on. It contains the contrib ID (hash of all other info), its contents, the time at which it was sent, and the prior contrib.
type contrib struct {
	id        string
	prev      string
	timestamp string
	contents  string
}

// This function returns the ID of the previous entry. if the ID variable exists. If so, it returns that variable. If not, it checks if the test.txt file exists. If so, it will read the most
// ecent contribution's ID. If not, then it falls back to returning "root", assuming this must be the initializing commit of a given ring.
func getPrev(id string, room string) string {
	if id != "" {
		return id
	} else {
		file, _ := os.Open(room)

		if file != nil {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			line := strings.TrimPrefix(text[len(text)-4], "ID ")

			file.Close()

			return line
		} else {
			file.Close()

			return "root"
		}
	}
}

// Gets the user's input using the bufio library and returns it as a string.
func getInput() string {
	inputReader := bufio.NewReader(os.Stdin)

	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

// Hashes the user input and the time it was inputted, and returns this as a string.
func getHash(strinput string, chrono string) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%s", strinput, chrono)))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Appends the contrib to the log, as well as prints it for debuggage purposes.
func commit(info contrib, room string) {
	pass := commandProcessing(info.contents)
	if pass {
		info.contents = strings.TrimPrefix(info.contents, "/plain ")
		file, _ := os.OpenFile(room, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // This works, for some reason. I have no idea what it's doing.

		file.Write([]byte(fmt.Sprintf("ID %s\nPrev %s\nTime %s\nContents %s\n", info.id, info.prev, info.timestamp, info.contents)))

		fmt.Printf("Message ID: %s\nPrevious Message: %s\nMessage Timestamp: %s\nMessage Contents: %s\n", info.id, info.prev, info.timestamp, info.contents)
	}
}

// The main loop of taking in user input, finding the time and previous hash, identifying the message, and displaying the message. It does this by calling other functions.
func mainloop(room string) {
	x := 0
	message := contrib{}

	for x < 1 {
		prev := getPrev(message.id, room)

		input := getInput()

		chrono := fmt.Sprintf("%d", time.Now().Unix())

		hash := getHash(input, chrono)

		message = contrib{
			id:        hash,
			timestamp: chrono,
			contents:  input,
			prev:      prev,
		}

		commit(message, room)
	}
}

