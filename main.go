// @title Chinook API
// @version 1.0
// @description RESTful API for Chinook database
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"chinook-api/internal/config"
	"chinook-api/internal/logging"
	"chinook-api/internal/routes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "chinook-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("No .env file found or unable to load")

	}
	cfg := config.LoadConfig()
	db := config.SetupDB(cfg.DBPath)
	defer db.Close()

	logFile, err := logging.InitLogger("app.log")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer logFile.Close()

	r := gin.New()
	r.Use(logging.RequestContextMiddleware())
	r.Use(cors.Default())
	r.Use(logging.ZerologMiddleware(), gin.Recovery())
	routes.SetupRoutes(r, db)

	// Add Swagger docs route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Info().Msg("Server running at http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown: %v", err)
	}

	log.Info().Msg("Server exiting")
}
