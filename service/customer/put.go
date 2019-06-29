package customer

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/model"
)

func (s *CustomerService) PutByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
		return
	}

	customer := model.Customer{ID: id}
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	stmt, err := s.Database.Prepare("UPDATE customers SET name = $2, email = $3, status = $4 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}
	
	c.JSON(http.StatusOK, customer)
	
}