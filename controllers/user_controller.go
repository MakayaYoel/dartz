package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/MakayaYoel/dartz/auth"
	"github.com/MakayaYoel/dartz/models"
	"github.com/MakayaYoel/dartz/repository"
	"github.com/gin-gonic/gin"
)

// RegisterUser creates a user.
func RegisterUser(c *gin.Context) {
	var userInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not process request."})
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

	// Insert user into database
	user, err := repository.CreateUser(models.User{Username: userInput.Username, Email: userInput.Email, Password: userInput.Password})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Return message with user struct (with ID field and hashed password)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

// AuthenticateUser attempts to authenticate the user.
func AuthenticateUser(c *gin.Context) {
	var userInput struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not process request."})
		return
	}

	user, err := repository.Authenticate(userInput)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwtToken, err := auth.CreateJWTToken(user.Username)

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.Request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully authenticated user.", "user": user})
}

// isValidUsername validates the given username. It returns an error if the username isn't valid.
func isValidUsername(username string) error {
	usernameMinLength, usernameMaxLength := 3, 24

	if len(username) < usernameMinLength {
		return fmt.Errorf("Username length has to be at least %d characters.", usernameMinLength)
	}

	if len(username) > usernameMaxLength {
		return fmt.Errorf("Username length cannot be over %d characters.", usernameMaxLength)
	}

	exists, err := repository.CheckUsernameExists(username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("That username is already being used.")
	}

	return nil
}

// isValidEmail validates the given email address. It returns an error if the email isn't valid.
func isValidEmail(email string) (string, error) {
	e, err := mail.ParseAddress(email)

	if err != nil {
		return "", errors.New("Invalid email address given.")
	}

	exists, err := repository.CheckEmailExists(e.Address)
	if err != nil {
		return "", err
	}

	if exists {
		return "", errors.New("That email address is already being used.")
	}

	return e.Address, nil
}

// isValidPassword validates the given password. It returns an error if the password isn't valid.
func isValidPassword(password string) error {
	passwordMinLength, passwordMaxLength := 8, 36

	if len(password) < passwordMinLength {
		return fmt.Errorf("Password length has to be at least %d characters.", passwordMinLength)
	}

	if len(password) > passwordMaxLength {
		return fmt.Errorf("Password length cannot be over %d characters.", passwordMaxLength)
	}

	return nil
}
