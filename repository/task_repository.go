package repository

import (
	"fmt"

	"github.com/MakayaYoel/dartz/config"
	"github.com/MakayaYoel/dartz/models"
	"github.com/MakayaYoel/dartz/queries"
)

// GetTasks returns a slice containing all tasks. It returns an error if the attempt rendered unsuccessful.
func GetTasks() ([]models.Task, error) {
	db := config.GetDB()

	var tasks []models.Task

	res, err := db.Query(queries.GetAllTasks)

	if err != nil {
		return tasks, fmt.Errorf("ran into an error trying to retrieve all tasks: %s", err.Error())
	}

	for res.Next() {
		var row models.Task

		res.Scan(&row.ID, &row.Title, &row.Description, &row.Priority, &row.DueDate)

		tasks = append(tasks, row)
	}

	return tasks, nil
}

// GetTaskByID returns a task based on its ID.
func GetTaskByID(taskID int) (models.Task, error) {
	db := config.GetDB()

	rows := db.QueryRow(queries.GetTaskByID, taskID)

	var task models.Task

	err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueDate)

	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to fetch task by ID: %s", err.Error())
	}

	return task, nil
}
