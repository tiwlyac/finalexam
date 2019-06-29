package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/model"
)

func PostByIDHandler(c *gin.Context) {
	db := database.Connect(c)
	defer db.Close()
	
	customer := model.Customer{}
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
	}

	row := db.QueryRow("INSERT INTO customers (name, status) VALUES ($1, $2) RETURNING id")
	err = row.Scan(&customer.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	c.JSON(http.StatusCreated, customer)
} 