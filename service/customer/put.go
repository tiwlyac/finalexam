package customer

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
	"github.com/tiwly/finalexam/model"
)

func PutByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
	}

	db := database.Connect(c)
	defer db.Close()

	customer := model.Customer{ID: id}
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	stmt, err := db.Prepare("UPDATE todos SET title = $1, status = $2 WHERE id=$3;")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	c.JSON(http.StatusOK, customer)
	
}