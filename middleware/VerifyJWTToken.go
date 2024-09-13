package middleware

import (
	"net/http"

	"github.com/MakayaYoel/dartz/auth"
	"github.com/gin-gonic/gin"
)

// VerifyJWTAuthToken verifies the given token.
func VerifyJWTAuthToken(c *gin.Context) {
	err := auth.VerifyJWTToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}
