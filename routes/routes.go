package routes

import (
	"github.com/MakayaYoel/dartz/controllers"
	"github.com/MakayaYoel/dartz/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	mainRouteGroup := router.Group("/api")
	mainRouteGroup.Use(middleware.VerifyJWTAuthToken)

	// API Routes
	mainRouteGroup.GET("/tasks", controllers.GetAllTasks)
	mainRouteGroup.POST("/tasks", controllers.CreateTask)

	mainRouteGroup.GET("/tasks/:id", controllers.GetTask)
	mainRouteGroup.PUT("/tasks/:id", controllers.UpdateTask)
	mainRouteGroup.DELETE("/tasks/:id")

	// Auth Routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.AuthenticateUser)
}
