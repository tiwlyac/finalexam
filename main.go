package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/model"
	"github.com/tiwly/finalexam/service"
	"strconv"
)

func main() {
	database.CreateTable()
	r := gin.Default()
	r.GET("/customers", service.GetCustomersHandler)
	r.GET("/customers/:id", service.GetCustomersHandler)
	r.Run(":2019")
}

