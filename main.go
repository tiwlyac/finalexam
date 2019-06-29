package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/service/customer"
)

func main() {
	database.CreateTable()
	r := setupRouter()
	r.Run(":2019")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/customers", customer.GetCustomersHandler)
	r.GET("/customers/:id", customer.GetCustomersHandler)
	r.POST("/customers", customer.PostByIDHandler)
	r.PUT("/customers/:id", customer.PutByIDHandler)
	r.DELETE("/customers/:id", customer.DeleteHandler)
	return r
}

