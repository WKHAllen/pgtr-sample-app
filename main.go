package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "main/src"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
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
	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize the database
	dbURL := os.Getenv("DATABASE_URL")
	dbm, err := app.NewDBManager(dbURL)
	if err != nil {
		panic(err)
	}
	defer dbm.Close()

	app.InitDB(dbm)

	// Force Gin console color
	gin.ForceConsoleColor()

	// Set up routing
	router := gin.Default()

	router.GET("/api/hello", getMessage)
	router.POST("/api/hello", postSomething)

	// Set up server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen error:", err)
		}
	}()

	// Wait for interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Forced shutdown:", err)
	}
	log.Println("Server exiting")
}
