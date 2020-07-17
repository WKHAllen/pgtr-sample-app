package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONError sends an error message in JSON format if an error has occurred
func JSONError(c *gin.Context, err error, errorMsg string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": errorMsg,
		})
	}
	return err != nil
}
