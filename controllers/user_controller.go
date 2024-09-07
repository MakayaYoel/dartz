package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/MakayaYoel/dartz/models"
	"github.com/MakayaYoel/dartz/repository"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var userInput models.User

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not process request"})
		return
	}

	// Validate username
	if err := isValidUsername(userInput.Username); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validate email
	var err error
	userInput.Email, err = isValidEmail(userInput.Email)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validate password
	if err := isValidPassword(userInput.Password); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := repository.CreateUser(userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user created successfully", "user": userInput})
}

// isValidUsername validates the given username. It returns an error if the username isn't valid.
func isValidUsername(username string) error {
	usernameMinLength, usernameMaxLength := 3, 24

	if len(username) < usernameMinLength {
		return fmt.Errorf("username length has to be at least %d characters", usernameMinLength)
	}

	if len(username) > usernameMaxLength {
		return fmt.Errorf("username length cannot be over %d characters", usernameMaxLength)
	}

	return nil
}

// isValidEmail validates the given email address. It returns an error if the email isn't valid.
func isValidEmail(email string) (string, error) {
	e, err := mail.ParseAddress(email)

	if err != nil {
		return "", errors.New("invalid email address given")
	}

	return e.Address, nil
}

// isValidPassword validates the given password. It returns an error if the password isn't valid.
func isValidPassword(password string) error {
	passwordMinLength, passwordMaxLength := 8, 48

	if len(password) < passwordMinLength {
		return fmt.Errorf("password length has to be at least %d characters", passwordMinLength)
	}

	if len(password) > passwordMaxLength {
		return fmt.Errorf("password length cannot be over %d characters", passwordMaxLength)
	}

	return nil
}
