package routes

import (
	"chinook-api/internal/handlers"
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// auth
	userRepo := &repositories.UserRepository{DB: db}
	refreshTokenRepo := &repositories.RefreshTokenRepository{DB: db}
	authHandler := &handlers.AuthHandler{
		UserRepo:         userRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
	// artists
	artistRepo := &repositories.ArtistRepository{DB: db}
	artistHandler := &handlers.ArtistHandler{Repo: artistRepo}

	// albums
	albumRepo := &repositories.AlbumRepository{DB: db}
	albumHandler := &handlers.AlbumHandler{Repo: albumRepo}

	// employees
	employeeRepo := &repositories.EmployeeRepository{DB: db}
	employeeHandler := &handlers.EmployeeHandler{Repo: employeeRepo}

	// tracks
	trackRepo := &repositories.TrackRepository{DB: db}
	trackHandler := &handlers.TrackHandler{Repo: trackRepo}

	// genres
	genreRepo := &repositories.GenreRepository{DB: db}
	genreHandler := &handlers.GenreHandler{Repo: genreRepo}

	r.NoRoute(notFoundHandler)
	r.Use(internalServerErrorMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	version := os.Getenv("API_VERSION")
	api := r.Group("/api/" + version)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes (no JWT required for login/signup/refresh)
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/refresh", authHandler.Refresh)
	}

	// Protected routes
	var protected *gin.RouterGroup
	if os.Getenv("GO_ENV") == "production" {
		protected = api.Group("", utils.AuthMiddlewareJWT())
	} else {
		protected = api.Group("")
	}
	{
		protected.GET("/auth/me", authHandler.Me)
		protected.POST("/auth/logout", authHandler.Logout)
		artists := protected.Group("/artists")
		{
			artists.GET("", artistHandler.GetAll)
			artists.GET("/:id", artistHandler.GetOne)
			artists.POST("", artistHandler.Create)
			artists.PUT("/:id", artistHandler.Update)
			artists.DELETE("/:id", artistHandler.Delete)
		}

		albums := protected.Group("/albums")
		{
			albums.GET("", albumHandler.GetAll)
			albums.GET("/:id", albumHandler.GetOne)
			albums.POST("", albumHandler.Create)
			albums.PUT("/:id", albumHandler.Update)
			albums.DELETE("/:id", albumHandler.Delete)
		}

		employees := protected.Group("/employees")
		{
			employees.GET("", employeeHandler.GetAll)
			employees.GET("/:id", employeeHandler.GetOne)
		}

		tracks := protected.Group("/tracks")
		{
			tracks.GET("", trackHandler.GetAll)
			tracks.GET("/:id", trackHandler.GetOne)
		}

		genres := protected.Group("/genres")
		{
			genres.GET("", genreHandler.GetAll)
			genres.GET("/:id", genreHandler.GetOne)
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
		if len(c.Errors) > 0 && !c.Writer.Written() {
			c.JSON(500, gin.H{
				"error":   "internal server error",
				"details": c.Errors.Last().Error(),
			})
		}
	}
}
