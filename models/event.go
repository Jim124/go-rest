package models

import (
	"fmt"
	"go-rest/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event

func (e *Event) Save() error {
	insertSql := `insert into events(name,description,location,date_time,user_id)values(?,?,?,?,?)`
	stmt, error := db.DB.Prepare(insertSql)
	if error != nil {
		fmt.Println(error)
		return error
	}
	defer stmt.Close()
	result, error := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if error != nil {
		fmt.Println(error)
		return nil
	}
	id, error := result.LastInsertId()
	if error != nil {
		fmt.Println(error)
		return error
	}
	e.ID = id
	return nil
}

func GetEvents() ([]Event, error) {
	var events []Event
	querySql := "select * from events"
	rows, error := db.DB.Query(querySql)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	singleQuery := "select * from events where id=?"
	row := db.DB.QueryRow(singleQuery, id)
	var event Event
	error := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if error != nil {
		return nil, error
	}
	return &event, nil
}

func (e Event) UpdateEvent() error {
	updateSql := `update events set name=?,description=?,location=?,date_time = ? where id=?`
	stmt, error := db.DB.Prepare(updateSql)
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if error != nil {
		return error
	}
	return nil
}
