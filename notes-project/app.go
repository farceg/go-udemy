package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-project/note"
)

func main() {

	noteName, noteContent := getNoteData()
	newNote, err := note.New(noteName, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.PrintNote()
	err = newNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

}
func getNoteData() (string, string) {
	noteName := getUserInput("Please enter the note name: ")
	noteContent := getUserInput("Please enter the note content: ")

	return noteName, noteContent
}

func getUserInput(promptText string) string {
	fmt.Print(promptText)

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(value)
}
