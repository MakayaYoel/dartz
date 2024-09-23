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

		res.Scan(&row.ID, &row.Title, &row.Description, &row.Priority, &row.CreatedAt, &row.Completed)

		tasks = append(tasks, row)
	}

	return tasks, nil
}

// GetTaskByID returns a task based on its ID.
func GetTaskByID(taskID int) (models.Task, error) {
	db := config.GetDB()

	rows := db.QueryRow(queries.GetTaskByID, taskID)

	var task models.Task

	err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.CreatedAt, &task.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to fetch task by ID: %s", err.Error())
	}

	return task, nil
}

// AddTask creates a task and returns it.
func AddTask(d interface{}) (models.Task, error) {
	userInput, ok := d.(struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		CreatedAt   int    `json:"created_at"`
		Completed   bool   `json:"completed"`
	})

	if !ok {
		return models.Task{}, fmt.Errorf("could not process user input when trying to add task")
	}

	db := config.GetDB()

	stmt, err := db.Prepare(queries.AddTask)
	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to add a task: %s", err.Error())
	}

	res, err := stmt.Exec(userInput.Title, userInput.Description, userInput.Priority, userInput.CreatedAt, userInput.Completed)
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

// UpdateTask updates the specified task.
func UpdateTask(ID int, d interface{}) (models.Task, error) {
	userInput, ok := d.(struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
		CreatedAt   int    `json:"created_at"`
		Completed   bool   `json:"completed"`
	})

	if !ok {
		return models.Task{}, fmt.Errorf("could not process user input when trying to add task")
	}

	db := config.GetDB()

	stmt, err := db.Prepare(queries.UpdateTask)
	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to update a task: %s", err.Error())
	}

	_, err = stmt.Exec(ID, userInput.Title, userInput.Description, userInput.Priority, userInput.CreatedAt, userInput.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("ran into an error trying to update a task: %s", err.Error())
	}

	task, err := GetTaskByID(ID)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

// DeleteTask deletes a task.
func DeleteTask(task models.Task) error {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.DeleteTask)

	if err != nil {
		return fmt.Errorf("ran into an error trying to delete a task: %s", err.Error())
	}

	_, err = stmt.Exec(task.ID)

	if err != nil {
		return fmt.Errorf("ran into an error trying to delete a task: %s", err.Error())
	}

	return nil
}
