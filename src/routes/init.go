package routes

import (
	"github.com/gin-gonic/gin"
)

// LoadRoutes creates all available routes
func LoadRoutes(router *gin.Engine, path string) {
	api := router.Group(path)

	api.GET("/person/:id", GetPerson)
	api.GET("/people",     GetPeople)
	api.GET("/quote/:id",  GetQuote)
	api.GET("/quotes",     GetQuotes)
}
