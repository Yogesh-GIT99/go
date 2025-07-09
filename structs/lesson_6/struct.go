package main

import (
	"fmt"
	"structs/lesson_6/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	// var appUser user

	appUser, err := &user.Newuser(userfirstName, userlastName, userbirthdate)

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
