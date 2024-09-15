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

func AddTask(userInput interface{}) (models.Task, error) {
	uInput, ok := userInput.(struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		DueDate     int    `json:"due_date"`
	})

	if !ok {
		return models.Task{}, fmt.Errorf("could not process user input when trying to add task")
	}

	db := config.GetDB()

	stmt, err := db.Prepare(queries.AddTask)

	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to add a task: %s", err.Error())
	}

	res, err := stmt.Exec(uInput.Title, uInput.Description, uInput.Priority, uInput.DueDate)

	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to add a task: %s", err.Error())
	}

	insertID, err := res.LastInsertId()

	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to add a task: %s", err.Error())
	}

	task, err := GetTaskByID(int(insertID))

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
