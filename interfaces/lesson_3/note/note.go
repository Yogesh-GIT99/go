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
	NoteTitle   string    `json:"title"`
	NoteBody    string    `json:"content"`
	CurrentTime time.Time `json:"currenttime"`
}

func NewNote(noteTitle, noteBody string) (Note, error) {

	if noteTitle == "" || noteBody == "" {
		return Note{}, errors.New("notetile or notebody cannot be empty")
	}

	return Note{
		NoteTitle:   noteTitle,
		NoteBody:    noteBody,
		CurrentTime: time.Now(),
	}, nil

}

func (n Note) OutputNote() {

	fmt.Printf("\nThe note you have writte in %v\n\n %v\n", n.NoteTitle, n.NoteBody)
	fmt.Println("------------")

}

func (n Note) WriteFloatToFile() {

	content := Note{

		NoteTitle:   n.NoteTitle,
		NoteBody:    n.NoteBody,
		CurrentTime: time.Now(),
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	}

	name := strings.Replace(n.NoteTitle, " ", "_", -1)
	filename := strings.ToLower(name) + ".json"

	os.WriteFile(filename, jsonData, 0644)

	fmt.Println("Note is written Successfuly")

}
