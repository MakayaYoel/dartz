package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MakayaYoel/dartz/models"
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

	var cleanTasks []map[string]interface{}

	for _, t := range tasks {
		cleanTasks = append(cleanTasks, cleanTaskStruct(t))
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": cleanTasks})
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

	c.IndentedJSON(http.StatusOK, gin.H{"task": cleanTaskStruct(task)})
}

// CreateTask creates a new task.
func CreateTask(c *gin.Context) {
	var userInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		CreatedAt   int    `json:"created_at"`
		Completed   bool   `json:"completed"`
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

	c.IndentedJSON(http.StatusOK, gin.H{"message": "created task successfully.", "task": cleanTaskStruct(task)})
}

// UpdateTask updates the specified task.
func UpdateTask(c *gin.Context) {
	var userInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		CreatedAt   int    `json:"created_at"`
		Completed   bool   `json:"completed"`
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

	c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully updated task.", "task": cleanTaskStruct(task)})
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

func cleanTaskStruct(t models.Task) map[string]interface{} {
	tm := time.Unix(int64(t.CreatedAt), 0)

	var priority string
	switch t.Priority {
	case 0:
		priority = "Low"
	case 1:
		priority = "Medium"
	case 2:
		priority = "Urgent"
	default:
		priority = "N/A"
	}

	return map[string]interface{}{
		"id":          t.ID,
		"title":       t.Title,
		"description": t.Description,
		"priority":    priority,
		"created_at":  tm.Format("Monday, January 02, 2006"),
		"completed":   t.Completed,
	}
}
