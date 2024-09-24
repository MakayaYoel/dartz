package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MakayaYoel/dartz/auth"
	"github.com/MakayaYoel/dartz/queries"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var dbUsername, dbPassword, dbHost, dbDatabase, dbPort string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("There was an error trying to load the environment file: %s", err.Error())
	}
}

func LoadConfig() {
	if err := loadEnvVariables(); err != nil {
		log.Fatalf("There was an error trying to load the environment variables: %s", err.Error())
	}

	// Load Database
	database, err := sql.Open("mysql", fmt.Sprintf("%s%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbDatabase))

	if err != nil {
		log.Fatalf("Could not start database: %s", err.Error())
	}

	// Initialize tables
	_, err = database.Exec(queries.CreateUsersTable)

	if err != nil {
		log.Fatalf("Could not initialize database tables: %s", err.Error())
	}

	_, err = database.Exec(queries.CreateTasksTable)

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

	return db
}

func loadEnvVariables() error {
	var ok bool

	dbUsername, ok = os.LookupEnv("DB_USERNAME")
	if !ok || len(dbUsername) == 0 {
		return fmt.Errorf("could not load DB_USERNAME environment variable")
	}

	dbPassword, ok = os.LookupEnv("DB_PASSWORD")
	if !ok {
		return fmt.Errorf("could not load DB_PASSWORD environment variable")
	}

	if len(dbPassword) > 0 {
		dbPassword = ":" + dbPassword
	}

	dbHost, ok = os.LookupEnv("DB_HOST")
	if !ok || len(dbHost) == 0 {
		return fmt.Errorf("could not load DB_HOST environment variable")
	}

	dbDatabase, ok = os.LookupEnv("DB_DATABASE")
	if !ok || len(dbDatabase) == 0 {
		return fmt.Errorf("could not load DB_DATABASE environment variable")
	}

	dbPort, ok = os.LookupEnv("DB_PORT")
	if !ok || len(dbPort) == 0 {
		return fmt.Errorf("could not load DB_PORT environment variable")
	}

	jwtSecretKey, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok || len(jwtSecretKey) == 0 {
		return fmt.Errorf("could not load JWT_SECRET_KEY environment variable")
	}

	auth.JWTSecretKey = []byte(jwtSecretKey)

	return nil
}
