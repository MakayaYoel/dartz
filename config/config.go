package config

import (
	"database/sql"
	"log"

	"github.com/MakayaYoel/dartz/queries"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func LoadConfig() {
	// Load Database
	database, err := sql.Open("mysql", "yoel@tcp(127.0.0.1:3310)/dartz")

	if err != nil {
		log.Fatalf("Could not start database: %s", err.Error())
	}

	// Initialize tables
	_, err = database.Exec(queries.CreateTables)

	if err != nil {
		log.Fatalf("Could not initialize database tables: %s", err.Error())
	}

	db = database
}

// GetDB pings the database to verify if a connection is still alive then it returns it. It throws a fatal error if a connection could not be found.
func GetDB() *sql.DB {
	if err := db.Ping(); err != nil {
		log.Fatalf("Lost connection with database: %s", err.Error())
	}

	log.Println("Still have connection with db", db)

	return db
}
