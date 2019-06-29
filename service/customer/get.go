package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/model"
	"strconv"
)

func (s *CustomerService) GetCustomersHandler(c *gin.Context) {
	customers := []model.Customer{}
	stmt, err := s.Database.Prepare("SELECT id, name, email, status FROM customers;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	for rows.Next() {
		customer := model.Customer{}
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		} 
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

func (s *CustomerService) GetCustomersByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error" : http.StatusText(http.StatusBadRequest)})
		return
	}

	customer := model.Customer{}
	stmt, err := s.Database.Prepare("SELECT id, name, email, status FROM customers WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	row := stmt.QueryRow(id)
	if err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, customer)
}
