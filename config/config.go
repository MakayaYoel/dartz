package config

import (
	"database/sql"
	"log"

	"github.com/MakayaYoel/dartz/queries"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func LoadConfig() {
	// Load Database
	DB, err := sql.Open("mysql", "yoel@tcp(127.0.0.1:3310)/dartz")

	if err != nil {
		log.Fatalf("Could not start database: %s", err.Error())
	}

	// Initialize tables
	_, err = DB.Exec(queries.CreateTables)

	if err != nil {
		log.Fatalf("Could not initialize database tables: %s", err.Error())
	}
}
