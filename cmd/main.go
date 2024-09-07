package main

import (
	"log"

	"github.com/MakayaYoel/dartz/config"
	"github.com/MakayaYoel/dartz/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Create routes
	routes.RegisterRoutes(router)

	// Load configurations (Load DB)
	config.LoadConfig()

	if err := router.Run(); err != nil {
		log.Fatalf("There was an error trying to start the web server: %s", err.Error())
	}
}
