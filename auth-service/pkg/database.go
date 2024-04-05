package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewDBConnection() *sql.DB {
	return connectDatabase()
}

func connectDatabase() *sql.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	pingDatabase(db)
	return db
}

func pingDatabase(db *sql.DB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
