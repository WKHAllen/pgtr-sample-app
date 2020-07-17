package routes

import (
	"net/http"
	"strconv"

	"main/src/services"

	"github.com/gin-gonic/gin"
)

// GetPerson gets a person from the database
func GetPerson(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Person ID must be an integer",
		})
		return
	}

	person, err := services.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid person ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
