package controllers

import (
	"net/http"

	"github.com/MakayaYoel/dartz/repository"
	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, err := repository.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}
