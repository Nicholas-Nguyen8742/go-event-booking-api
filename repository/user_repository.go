package repository

import (
	"errors"
	"event-booking-api/storage"
	"event-booking-api/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := storage.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := storage.DB.QueryRow(query, user.Email)

	var retrievedPassword string

	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.ValidatePasswordHash(user.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
