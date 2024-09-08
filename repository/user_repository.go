package repository

import (
	"fmt"

	"github.com/MakayaYoel/dartz/config"
	"github.com/MakayaYoel/dartz/models"
	"github.com/MakayaYoel/dartz/queries"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserts a new user into the database. It returns an error if the attempt rendered unsuccessful.
func CreateUser(u models.User) (models.User, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.CreateUser)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	u.Password = string(password)

	_, err = stmt.Exec(u.Username, u.Email, u.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	// Get user model struct (with ID field and hashed password)
	user, err := GetUserByUsername(u.Username)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to create a user: %s", err.Error())
	}

	return user, nil
}

// Authenticate attempts to authenticate a user using the given credentials. It returns an error if the attempt rendered unsuccessful.
func Authenticate(d interface{}) (models.User, error) {
	userInput, ok := d.(struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	})

	if !ok {
		return models.User{}, fmt.Errorf("cannot process user input")
	}

	auth := func(u models.User) (models.User, error) {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userInput.Password))

		if err != nil {
			return models.User{}, fmt.Errorf("invalid credentials")
		}

		return u, nil
	}

	if user, err := GetUserByEmail(userInput.UsernameOrEmail); err == nil {
		return auth(user)
	} else if user, err := GetUserByUsername(userInput.UsernameOrEmail); err == nil {
		return auth(user)
	}

	return models.User{}, fmt.Errorf("user does not exist")
}

// GetUserByUsername returns the specified user's model struct. It returns an error if the attempt rendered unsuccessful.
func GetUserByUsername(username string) (models.User, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.GetUserByUsername)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to fetch user by username: %s", err.Error())
	}

	var user models.User

	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to fetch user by username: %s", err.Error())
	}

	return user, nil
}

// GetUserByEmail returns the specified user's model struct. It returns an error if the attempt rendered unsuccessful.
func GetUserByEmail(email string) (models.User, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.GetUserByEmail)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to fetch user by email: %s", err.Error())
	}

	var user models.User

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("ran into an error trying to fetch user by username: %s", err.Error())
	}

	return user, nil
}

// CheckUsernameExists returns whether the specified username already exists in the DB. It returns an error if the attempt rendered unsuccessful.
func CheckUsernameExists(username string) (bool, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.CheckExistingUsername)
	if err != nil {
		return false, fmt.Errorf("ran into an error trying to verify if username exists: %s", err.Error())
	}

	var count int

	err = stmt.QueryRow(username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("ran into an error trying to verify if username exists: %s", err.Error())
	}

	return count != 0, nil
}

// CheckEmailExists returns whether the specified email already exists in the DB. It returns an error if the attempt rendered unsuccessful.
func CheckEmailExists(email string) (bool, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.CheckExistingEmail)
	if err != nil {
		return false, fmt.Errorf("ran into an error trying to verify if email exists: %s", err.Error())
	}

	var count int

	err = stmt.QueryRow(email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("ran into an error trying to verify if email exists: %s", err.Error())
	}

	return count != 0, nil
}
