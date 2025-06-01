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

// func newuser(firstName, lastName, birthDate string) user {

// 	return user{

// 		firstName:   firstName,
// 		lastName:    lastName,
// 		birthDate:   birthDate,
// 		currentTime: time.Now(),
// 	}
// }

func newuser(firstName, lastName, birthDate string) *user { // with pointer

	return &user{

		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		currentTime: time.Now(),
	}

}

func (u *user) outputUserData() { // defining parameters using struct using pointer

	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate) // calling out struct values using pointer the real way
	fmt.Println(u.firstName, u.lastName, u.birthDate) // this is not correct though but go supports this syntax for the ease of writing code.

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

	appUser := newuser(userfirstName, userlastName, userbirthdate)

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
