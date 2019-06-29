package database

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/lib/pq" 
)

func Connect() (*sql.DB, error)  {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) 
	return db, err
}

func CreateTable(db *sql.DB) {
	createTb := `
	CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	)
	`
	if _, err := db.Exec(createTb); err != nil {
		log.Fatal("Can't create table", err.Error())
	} 
}






