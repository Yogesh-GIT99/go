package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // defining a struct

	firstName   string
	lastName    string
	birthDate   string
	currentTime time.Time
}

type Admin struct { // struct embeding
	email       string
	personal_id string
	User        // declaring anonymous struct, can also give it name like:  user User
}

func Newadmin(email, personal_id string) Admin {

	return Admin{
		email:       email,
		personal_id: personal_id,
		User: User{
			firstName:   "Admin",
			lastName:    "Admin",
			birthDate:   "---",
			currentTime: time.Now(),
		},
	}

}

func Newuser(firstName, lastName, birthDate string) (*User, error) { // with pointer

	if firstName == "" || lastName == "" || birthDate == "" {

		return nil, errors.New("firstName, lastName and birthDate should not be empty")
	}

	return &User{

		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		currentTime: time.Now(),
	}, nil

}

func (u *User) OutputUserData() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *User) CleanuserData() {
	u.firstName = ""
	u.lastName = ""
}
