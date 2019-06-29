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
		c.JSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
		return
	}

	stmt, err := db.Prepare("INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	row := stmt.QueryRow(&customer.Name, &customer.Email, &customer.Status)
	err = row.Scan(&customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusCreated, customer)
} 