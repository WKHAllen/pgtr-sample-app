package routes

import (
	"github.com/gin-gonic/gin"
)

// LoadRoutes creates all available routes
func LoadRoutes(router *gin.Engine, path string) {
	api := router.Group(path)

	api.GET("/person/id/:id", GetPerson)
	api.GET("/person/random", GetRandomPerson)
	api.GET("/people",        GetPeople)

	api.GET("/quote/id/:id",  GetQuote)
	api.GET("/quote/random",  GetRandomQuote)
	api.GET("/quotes",        GetQuotes)
}
