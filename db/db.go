package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb() {
	var error error
	db, error = sql.Open("mysql", "root:aA1243690.@tcp(localhost:3306)/go_rest?parseTime=true")
	if error != nil {
		panic(error)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	createTable()
}

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id int PRIMARY KEY auto_increment,
		name text NOT NULL,
		description text NOT NULL,
		location text NOT NULL,
		date_time DATETIME NOT NULL,
		user_id int
	)
	`

	_, err := db.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
