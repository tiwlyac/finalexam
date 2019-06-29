package main

import (
	"net/http"
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
	r.Use(checkAuthorization)
	r.GET("/customers", customer.GetCustomersHandler)
	r.GET("/customers/:id", customer.GetCustomersByIDHandler)
	r.POST("/customers", customer.PostByIDHandler)
	r.PUT("/customers/:id", customer.PutByIDHandler)
	r.DELETE("/customers/:id", customer.DeleteHandler)
	return r
}

func checkAuthorization(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": http.StatusText(http.StatusUnauthorized) })
		c.Abort()
	}
	c.Next()
}