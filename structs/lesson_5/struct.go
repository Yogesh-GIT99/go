package main

import (
	"errors"
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

func newuser(firstName, lastName, birthDate string) (*user, error) { // with pointer

	if firstName == "" || lastName == "" || birthDate == "" { // can also add a validation to a constructer
		return nil, errors.New("firstName, lastName and birthDate should not be empty")

	}

	return &user{

		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		currentTime: time.Now(),
	}, nil

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

	appUser, err := newuser(userfirstName, userlastName, userbirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData() // passing instanciated values to function using pointer
	appUser.cleanuserData()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
