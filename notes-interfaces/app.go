package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-project/note"
	"example.com/notes-project/todo"
)

type saver interface {
	Save() error // to implement this interface, the value must have it already implemented
}

// Sería posible llamar una interfaz dentro de otra
type printer interface {
	// saver
	Display()
}

func main() {

	todoText := getUserInput("Todo: ")
	newTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Todo struct have the Save() adn Display() methods implemented
	displayData(newTodo)
	err = saveData(newTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	noteName := getUserInput("Note name: ")
	noteContent := getUserInput("Note content: ")
	newNote, err := note.New(noteName, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Note struct have the Save() adn Display() methods implemented
	displayData(newNote)
	err = saveData(newNote)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

func displayData(data printer) {
	data.Display()
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
