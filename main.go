package main

import (
	"chinook-api/internal/config"
	"chinook-api/internal/logging"
	"chinook-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or unable to load")
	}
	db := config.SetupDB()
	defer db.Close()

	logFile, err := logging.InitLogger("app.log")
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	r := gin.New()
	r.Use(logging.JSONLogger(), gin.Recovery())
	routes.SetupRoutes(r, db)
	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
