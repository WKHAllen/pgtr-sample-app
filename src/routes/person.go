package routes

import (
	"net/http"

	"main/src/services"
	"main/src/routes/helper"

	"github.com/gin-gonic/gin"
)

// GetPerson gets a person from the database
func GetPerson(c *gin.Context) {
	id, err := helper.ParamInt(c, "id")
	if helper.JSONError(c, err, "Person ID must be an integer") { return }

	person, err := services.GetPerson(id)
	if helper.JSONError(c, err, "Invalid person ID") { return }

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
