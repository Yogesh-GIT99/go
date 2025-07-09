package user

import (
	"fmt"
	"time"
)

type User struct { // defining a struct

	firstName   string
	lastName    string
	birthDate   string
	currentTime time.Time
}

func Newuser(firstName, lastName, birthDate string) *User { // with pointer

	return &User{

		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		currentTime: time.Now(),
	}

}

func (u *User) outputUserData() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *User) cleanuserData() {
	u.firstName = ""
	u.lastName = ""
}
