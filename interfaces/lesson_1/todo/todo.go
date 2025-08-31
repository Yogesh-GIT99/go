package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func NewNote(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("content cannot be empty")
	}

	return Todo{
		Text: content,
	}, nil

}

func (todo Todo) OutputNote() {

	fmt.Println(todo.Text)

}

func (n Todo) WriteFloatToFile() {

	content := Todo{

		Text: n.Text,
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	}

	todofile := "todo.json"

	os.WriteFile(todofile, jsonData, 0644)

	fmt.Println("Note is written Successfuly")

}
