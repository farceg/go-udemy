package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	NoteName    string    `json:"note_name"`
	NoteContent string    `json:"note_content"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(noteName, noteContent string) (Note, error) {
	if noteName == "" || noteContent == "" {
		return Note{}, errors.New("no values passed")
	}
	return Note{
		NoteName:    noteName,
		NoteContent: noteContent,
		CreatedAt:   time.Now(),
	}, nil
}

func (note Note) PrintNote() {
	fmt.Printf("Note name: %v\nNote content: %v\n", note.NoteName, note.NoteContent)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.NoteName, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName+".json", json, 0644)
}
