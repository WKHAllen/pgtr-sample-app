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
	"main/src/db"
	"main/src/routes"
	"main/src/services"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize the database
	dbURL := os.Getenv("DATABASE_URL")
	dbm, err := db.NewDBManager(dbURL)
	if err != nil {
		panic(err)
	}
	defer dbm.Close()

	app.InitDB(dbm)

	// Force Gin console color
	gin.ForceConsoleColor()

	// Set up routing and services
	router := gin.Default()
	routes.LoadRoutes(router, "/api")
	services.SetDBManager(dbm)

	// Set up server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
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
		panic(err)
	}
	log.Println("Exiting successfully")
}
