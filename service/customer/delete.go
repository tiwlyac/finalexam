package customer

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/tiwly/finalexam/database"
)

func DeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
		return
	}

	db := database.Connect(c)
	defer db.Close()

	query := "DELETE FROM customers WHERE id=$1;"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
		return
	}

	stmt.Exec(id)
	c.JSON(http.StatusOK, gin.H{ "message": "customer deleted" })
}