package routes

import (
	"net/http"
	"strconv"

	"main/src/services"
	"main/src/routes/helper"

	"github.com/gin-gonic/gin"
)

// GetPerson gets a person from the database
func GetPerson(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if helper.JSONError(c, err, "Person ID must be an integer") { return }

	person, err := services.GetPerson(id)
	if helper.JSONError(c, err, "Invalid person ID") { return }

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
