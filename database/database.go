package database

import (
	"os"
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" 
)

func Connect(c *gin.Context) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}
	return db
}

func CreateTable() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("fatal", err.Error())
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	)
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err.Error())
	} 
}

