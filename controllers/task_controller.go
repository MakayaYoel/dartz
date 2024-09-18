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

// GetTask retrieves the specified task.
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

// CreateTask creates a new task.
func CreateTask(c *gin.Context) {
	var userInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		DueDate     int    `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not process request."})
		return
	}

	task, err := repository.AddTask(userInput)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "created task successfully.", "task": task})
}

// UpdateTask updates the specified task.
func UpdateTask(c *gin.Context) {
	var userInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		DueDate     int    `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not process request."})
		return
	}

	rawID := c.Param("id")
	intID, err := strconv.Atoi(rawID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not process request."})
		return
	}

	task, err := repository.UpdateTask(intID, userInput)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully updated task.", "task": task})
}

// DeleteTask deletes the specified task.
func DeleteTask(c *gin.Context) {
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

	err = repository.DeleteTask(task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully deleted task."})
}
