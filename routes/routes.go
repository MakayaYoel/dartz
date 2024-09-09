package routes

import (
	"github.com/MakayaYoel/dartz/controllers"
	"github.com/MakayaYoel/dartz/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	mainRouteGroup := router.Group("/api")
	mainRouteGroup.Use(middleware.VerifyJWTAuthToken)

	mainRouteGroup.GET("/test", func(ctx *gin.Context) {})

	// Auth routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.AuthenticateUser)
}
