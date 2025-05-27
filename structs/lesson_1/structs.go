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

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser user

	appUser = user{ // instanciating a struct

		firstName:   userfirstName,
		lastName:    userlastName,
		birthDate:   userbirthdate,
		currentTime: time.Now(),
	}

	outputUserData(appUser) // passing instanciated values to function
}

func outputUserData(u user) { // defining parameters using struct

	fmt.Println(u.firstName, u.lastName, u.birthDate) // calling out struct values

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
