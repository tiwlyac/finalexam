package customer

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (s *CustomerService) DeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
		return
	}

	query := "DELETE FROM customers WHERE id=$1;"
	stmt, err := s.Database.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{ "message": "customer deleted" })
}