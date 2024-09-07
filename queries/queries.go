package queries

const (
	CreateTables string = "CREATE TABLE IF NOT EXISTS users(id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL);"
	CreateUser   string = "REPLACE INTO users(username, email, password) VALUES(?, ?, ?)"
)
