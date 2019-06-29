package main

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/service/customer"
)

const port = ":2019"

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("fatal", err.Error())
	}
	defer db.Close()
	r := setupRouter(db)
	r.Run(port)
}

func setupRouter(db *sql.DB) *gin.Engine {
	s := customer.CustomerService{Database: db}
	r := gin.Default()
	r.Use(checkAuthorization)
	r.GET("/customers", s.GetCustomersHandler)
	r.GET("/customers/:id", s.GetCustomersByIDHandler)
	r.POST("/customers", s.PostByIDHandler)
	r.PUT("/customers/:id", s.PutByIDHandler)
	r.DELETE("/customers/:id", s.DeleteHandler)
	return r
}

func checkAuthorization(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": http.StatusText(http.StatusUnauthorized) })
		c.Abort()
	} else {
		c.Next()
	}
}