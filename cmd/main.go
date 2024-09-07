package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("There was an error trying to start the web server: %s", err.Error())
	}
}
