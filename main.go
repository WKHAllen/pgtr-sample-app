package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	message := c.Query("message")
	// get something...
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func postSomething(c *gin.Context) {
	// post something...
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()

	router.GET("/api/hello", getMessage)
	router.POST("/api/hello", postSomething)

	router.Run(":" + port)
}
