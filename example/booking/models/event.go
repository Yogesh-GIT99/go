package models

import (
	"booking/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} // intializing with an empty slice of events

func (e Event) Save() error {

	query := `INSERT INTO events(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() // connec closing managed by go itself

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() []Event {
	return events
}
