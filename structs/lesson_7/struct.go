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

	appUser, err := user.Newuser(userfirstName, userlastName, userbirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.Newadmin("yogesh@gmail.com", "1234")

	adminUser.OutputUserData()
	adminUser.CleanuserData()
	adminUser.OutputUserData()

	appUser.OutputUserData() // passing instanciated values to function using pointer
	appUser.CleanuserData()
	appUser.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
