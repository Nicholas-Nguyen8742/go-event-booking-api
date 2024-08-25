package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables(DB)
}

func createTables(sqliteDb *sql.DB) {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	statement, err := sqliteDb.Prepare(createEventsTable)
	if err != nil {
		panic("Could not create events database.")
	}

	statement.Exec()
}
