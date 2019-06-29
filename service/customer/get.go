package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/model"
	"strconv"
)

func GetCustomersHandler(c *gin.Context) {
	db := database.Connect(c)
	defer db.Close()

	customers := []model.Customer{}
	stmt, err := db.Prepare("SELECT id, name, status FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	for rows.Next() {
		customer := model.Customer{}
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		} 
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

func GetCustomersByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error" : http.StatusText(http.StatusBadRequest)})
	}

	db := database.Connect(c)
	defer db.Close()

	customer := model.Customer{}
	stmt, err := db.Prepare("SELECT id, name, status FROM customers WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	err = stmt.QueryRow(id).Scan(&customer.ID, &customer.Name, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	c.JSON(http.StatusOK, customer)
}
