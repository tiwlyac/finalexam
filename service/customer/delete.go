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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H { "error" : http.StatusText(http.StatusBadRequest) })
	}

	db := database.Connect(c)
	defer db.Close()

	query := "DELETE FROM WHERE id=$1;"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"error" : http.StatusText(http.StatusInternalServerError)})
	}

	stmt.Exec(id)
	c.JSON(http.StatusOK, gin.H{ "status": "success" })
}