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
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	statement, err := sqliteDb.Prepare(createUsersTable)
	if err != nil {
		panic("Could not create users database.")
	}

	statement.Exec()


	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	statement, err = sqliteDb.Prepare(createEventsTable)
	if err != nil {
		panic("Could not create events database.")
	}

	statement.Exec()

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	statement, err = sqliteDb.Prepare(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations database.")
	}

	statement.Exec()
}
