package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/service/customer"
)

func main() {
	database.CreateTable()
	r := gin.Default()
	r.GET("/customers", customer.GetCustomersHandler)
	r.GET("/customers/:id", customer.GetCustomersHandler)
	r.POST("/customers", customer.PostByIDHandler)
	r.Run(":2019")
}

