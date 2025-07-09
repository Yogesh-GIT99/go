package main

import (
	"fmt"
	"time"
)

type user struct { // defining a struct

	firstName   string
	lastName    string
	birthDate   string
	currentTime time.Time
}

func (u *user) outputUserData() { // defining parameters using struct using pointer

	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *user) cleanuserData() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	// var appUser user

	appUser := user{ // declaring and instanciating a struct

		firstName:   userfirstName,
		lastName:    userlastName,
		birthDate:   userbirthdate,
		currentTime: time.Now(),
	}

	appUser.outputUserData() // passing instanciated values to function using pointer
	appUser.cleanuserData()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
