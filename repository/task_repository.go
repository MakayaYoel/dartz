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
		return tasks, fmt.Errorf("SQL query failed on retrieving all tasks: %s.", err.Error())
	}

	for res.Next() {
		var row models.Task

		err = res.Scan(&row.ID, &row.Title, &row.Description, &row.Priority, &row.CreatedAt, &row.Completed)
		if err != nil {
			return tasks, fmt.Errorf("Failed to scan task row: %s.", err.Error())
		}

		tasks = append(tasks, row)
	}

	return tasks, nil
}

// GetTaskByID returns a task based on its ID.
func GetTaskByID(taskID int) (models.Task, error) {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.GetTaskByID)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not prepare SQL statement to fetch task by ID: %s.", err.Error())
	}

	var task models.Task

	err = stmt.QueryRow(taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.CreatedAt, &task.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("SQL query failed on retrieving task by ID: %s.", err.Error())
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
		return models.Task{}, fmt.Errorf("Could not process user input.")
	}

	db := config.GetDB()

	stmt, err := db.Prepare(queries.AddTask)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not prepare SQL statement to add a task: %s.", err.Error())
	}

	res, err := stmt.Exec(userInput.Title, userInput.Description, userInput.Priority, userInput.CreatedAt, userInput.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not execute statement to add a task: %s.", err.Error())
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not retrieve insert ID: %s.", err.Error())
	}

	task, err := GetTaskByID(int(insertID))
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not retrieve newly created task: %s.", err.Error())
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
		return models.Task{}, fmt.Errorf("Could not process user input.")
	}

	db := config.GetDB()

	stmt, err := db.Prepare(queries.UpdateTask)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not prepare SQL statement to update a task: %s.", err.Error())
	}

	_, err = stmt.Exec(ID, userInput.Title, userInput.Description, userInput.Priority, userInput.CreatedAt, userInput.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not execute statement to update a task: %s.", err.Error())
	}

	task, err := GetTaskByID(ID)
	if err != nil {
		return models.Task{}, fmt.Errorf("Could not retrieve updated task: %s.", err.Error())
	}

	return task, nil
}

// DeleteTask deletes a task.
func DeleteTask(task models.Task) error {
	db := config.GetDB()

	stmt, err := db.Prepare(queries.DeleteTask)
	if err != nil {
		return fmt.Errorf("Could not prepare SQL statement to delete a task: %s.", err.Error())
	}

	_, err = stmt.Exec(task.ID)
	if err != nil {
		return fmt.Errorf("Could not execute statement to delete a task: %s.", err.Error())
	}

	return nil
}
