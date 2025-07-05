package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:postgres@pg-container:5432/livreria?sslmode=disable"

	log.Printf("!!!!!!!!!!!!!!!!!connecting to %v", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
		return nil, err
	}
	return db, nil
}
