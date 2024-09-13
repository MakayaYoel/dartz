package controllers

import (
	"net/http"
	"strconv"

	"github.com/MakayaYoel/dartz/repository"
	"github.com/gin-gonic/gin"
)

// GetAllTasks retrieves all tasks.
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

func GetTask(c *gin.Context) {
	rawID := c.Param("id")
	intID, err := strconv.Atoi(rawID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not process request."})
		return
	}

	task, err := repository.GetTaskByID(intID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}
