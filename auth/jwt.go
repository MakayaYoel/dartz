package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey []byte

// CreateJWTToken creates a JWT Authentication Token for the specified username. It returns an error if a token could not be generated.
func CreateJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(JWTSecretKey)

	if err != nil {
		return "", fmt.Errorf("ran into an error trying to create a jwt token for %s: %s", username, err.Error())
	}

	return tokenString, nil
}

// VerifyJWTToken verifies the given token. It returns an error if the token is invalid.
func VerifyJWTToken(token string) error {
	invalidToken := fmt.Errorf("invalid jwt authentication token")

	if token == "" {
		return invalidToken
	}

	token = token[len("Bearer "):]

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		return fmt.Errorf("failed to verify jwt authentication token: %s", err.Error())
	}

	if !t.Valid {
		return invalidToken
	}

	return nil
}
