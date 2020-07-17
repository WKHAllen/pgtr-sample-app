package routes

import (
	"net/http"
	"strconv"

	"main/src/services"

	"github.com/gin-gonic/gin"
)

// GetQuote gets a quote from the database
func GetQuote(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Quote ID must be an integer",
		})
		return
	}

	quote, err := services.GetQuote(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid quote ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})
}
