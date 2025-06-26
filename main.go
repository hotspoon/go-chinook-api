package main

import (
	"chinook-api/internal/config"
	"chinook-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDB()
	defer db.Close()

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
