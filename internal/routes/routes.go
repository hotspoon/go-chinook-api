package routes

import (
	"chinook-api/internal/handlers"
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	artistRepo := &repositories.ArtistRepository{DB: db}
	artistHandler := &handlers.ArtistHandler{Repo: artistRepo}
	authHandler := &handlers.AuthHandler{}

	r.NoRoute(notFoundHandler)
	r.Use(internalServerErrorMiddleware())

	r.POST("/login", authHandler.Login)
	r.POST("/signup", authHandler.Signup)

	api := r.Group("/api", utils.AuthMiddlewareJWT())
	{
		artists := api.Group("/artists")
		{
			artists.GET("", artistHandler.GetAll)
			artists.GET("/:id", artistHandler.GetOne)
			artists.POST("", artistHandler.Create)
			artists.PUT("/:id", artistHandler.Update)
			artists.DELETE("/:id", artistHandler.Delete)
		}
	}
}

// Custom 404 handler
func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "resource not found"})
}

// Custom 500 middleware
func internalServerErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{
				"error":   "internal server error",
				"details": c.Errors.Last().Error(),
			})
		}
	}
}
