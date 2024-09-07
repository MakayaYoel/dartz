package routes

import (
	"github.com/MakayaYoel/dartz/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Auth routes
	router.POST("/register", controllers.RegisterUser)
}
