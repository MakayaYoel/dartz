package queries

const (
	CreateTables          = "CREATE TABLE IF NOT EXISTS users(id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL);"
	CreateUser            = "INSERT INTO users(username, email, password) VALUES(?, ?, ?)"
	GetUserByUsername     = "SELECT * FROM users WHERE username = LOWER(?)"
	GetUserByID           = "SELECT * FROM users WHERE id = ?"
	GetUserByEmail        = "SELECT * FROM users WHERE email = LOWER(?)"
	CheckExistingUsername = "SELECT COUNT(*) FROM users WHERE username = LOWER(?)"
	CheckExistingEmail    = "SELECT COUNT(*) FROM users WHERE email = LOWER(?)"
)
