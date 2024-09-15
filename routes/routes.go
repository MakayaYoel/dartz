package routes

import (
	"github.com/MakayaYoel/dartz/controllers"
	"github.com/MakayaYoel/dartz/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	mainRouteGroup := router.Group("/api")
	mainRouteGroup.Use(middleware.VerifyJWTAuthToken)

	mainRouteGroup.GET("/tasks", controllers.GetAllTasks)
	mainRouteGroup.GET("/tasks/:id", controllers.GetTask)
	mainRouteGroup.POST("/tasks", controllers.CreateTask)

	// Auth routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.AuthenticateUser)
}
