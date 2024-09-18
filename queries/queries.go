package queries

const (
	CreateUsersTable      = "CREATE TABLE IF NOT EXISTS users(id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL)"
	CreateTasksTable      = "CREATE TABLE IF NOT EXISTS tasks(id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, description MEDIUMTEXT NOT NULL, priority TINYINT NOT NULL, due_date BIGINT NULL DEFAULT NULL)"
	CreateUser            = "INSERT INTO users(username, email, password) VALUES(?, ?, ?)"
	GetUserByUsername     = "SELECT * FROM users WHERE username = LOWER(?)"
	GetUserByID           = "SELECT * FROM users WHERE id = ?"
	GetUserByEmail        = "SELECT * FROM users WHERE email = LOWER(?)"
	CheckExistingUsername = "SELECT COUNT(*) FROM users WHERE username = LOWER(?)"
	CheckExistingEmail    = "SELECT COUNT(*) FROM users WHERE email = LOWER(?)"
	GetAllTasks           = "SELECT * FROM tasks"
	GetTaskByID           = "SELECT * FROM tasks WHERE id = ?"
	AddTask               = "INSERT INTO tasks(title, description, priority, due_date) VALUES(?, ?, ?, ?)"
	UpdateTask            = "REPLACE INTO tasks(id, title, description, priority, due_date) VALUES(?, ?, ?, ?, ?)"
	DeleteTask            = "DELETE * FROM tasks WHERE id = ?"
)
