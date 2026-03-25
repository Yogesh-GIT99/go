package models

import (
	"booking/db"
	"booking/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	fmt.Println("userID Save func: ", id)
	user.ID = id
	return err

}

func (user *User) ValidateLogin() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievePassword string
	err := row.Scan(&user.ID, &retrievePassword)

	if err != nil {
		//fmt.Println(err)
		return errors.New("Credentials Invalid !")
	}

	validatePassword := utils.ComparePassword(retrievePassword, user.Password)

	if !validatePassword {
		//fmt.Println("password validation failed")
		return errors.New("Credentials Invalid !")
	}

	return nil
}
