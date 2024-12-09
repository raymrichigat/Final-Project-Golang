package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// PostgreSQL connection string format:
	// "user=USERNAME password=PASSWORD dbname=DBNAME sslmode=disable"
	connStr := "user=postgres password=database dbname=fp_pbkk sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
}
