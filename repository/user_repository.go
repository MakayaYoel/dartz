package repository

import (
	"fmt"

	"github.com/MakayaYoel/dartz/config"
	"github.com/MakayaYoel/dartz/models"
	"github.com/MakayaYoel/dartz/queries"
)

// CreateUser inserts a new user into the database. It returns an error if the attempt rendered unsuccessful.
func CreateUser(u models.User) error {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.CreateUser)
	if err != nil {
		return fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	_, err = stmt.Exec(u.Username, u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	return nil
}
