package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	note "project_1/note"
	todo "project_1/todo"
	"strings"
)

type WriteFloatToFiler interface {
	WriteFloatToFile()
}

func main() {

	title, body, err := checkInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := userInput("Enter Todo Text: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	newtodo, err := todo.NewNote(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNoteCreate, err := note.NewNote(title, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	newtodo.OutputNote()
	writeData(newtodo)

	newNoteCreate.OutputNote()
	writeData(newNoteCreate)
}

func writeData(data WriteFloatToFiler) {

	data.WriteFloatToFile()
}

func userInput(promptText string) (string, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptText)
	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return "", errors.New("invalid value")
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value, nil

}

func checkInput() (string, string, error) {

	title, err := userInput("Note Title: ")

	if err != nil {
		fmt.Println(err)
		return "", "", errors.New("invalid title")
	}

	body, err := userInput("Note Body: ")

	if err != nil {
		fmt.Println(err)
		return "", "", errors.New("invalid body")
	}

	return title, body, nil
}
