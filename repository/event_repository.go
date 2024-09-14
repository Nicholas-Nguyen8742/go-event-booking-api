package repository

import (
	"event-booking-api/storage"
	"time"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}

var events = []Event{}

func (event Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := storage.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := storage.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close();

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM events
		WHERE id = ?
	`

	row, err := storage.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	var event Event
	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func  (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	stmt, err := storage.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`

	stmt, err := storage.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (event Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := storage.DB.Prepare(query)
	if (err != nil) {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}

func (event Event) Cancel(userId int64) error {
	query := `
		DELETE FROM registrations
		WHERE event_id = ? AND user_id = ?
	`

	stmt, err := storage.DB.Prepare(query)
	if (err != nil) {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}
