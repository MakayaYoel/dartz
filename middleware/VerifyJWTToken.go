package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey = []byte("IHaveNoIdeaWhatIHaveToPutHere")

// VerifyJWTAuthToken verifies the given token. It returns an error if the token is invalid.
func VerifyJWTAuthToken(c *gin.Context) {
	invalidToken := fmt.Errorf("invalid jwt authentication token")

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": invalidToken.Error()})
		return
	}

	token = token[len("Bearer "):]

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("failed to verify jwt authentication token: %s", err.Error())})
		return
	}

	if !t.Valid {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": invalidToken.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "authorized"})
}
