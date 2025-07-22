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

	// media types
	mediaTypesRepo := &repositories.MediaTypeRepository{DB: db}
	mediaTypeHandler := &handlers.MediaTypeHandler{Repo: mediaTypesRepo}

	// playlists
	playlistRepo := &repositories.PlaylistRepository{DB: db}
	playlistHandler := &handlers.PlaylistHandler{Repo: playlistRepo}

	// playlist tracks
	playlistTrackRepo := &repositories.PlaylistTrackRepository{DB: db}
	playlistTrackHandler := &handlers.PlaylistTrackHandler{Repo: playlistTrackRepo}

	// customers
	customerRepo := &repositories.CustomerRepository{DB: db}
	customerHandler := &handlers.CustomerHandler{Repo: customerRepo}

	// invoices
	invoiceRepo := &repositories.InvoiceRepository{DB: db}
	invoiceHandler := &handlers.InvoiceHandler{Repo: invoiceRepo}

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
			artists.GET("/search", artistHandler.SearchByName)
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

		mediaTypes := protected.Group("/media_types")
		{
			mediaTypes.GET("", mediaTypeHandler.GetAll)
			mediaTypes.GET("/:id", mediaTypeHandler.GetOne)
		}

		playlists := protected.Group("/playlists")
		{
			playlists.GET("", playlistHandler.GetAll)
			playlists.GET("/:id", playlistHandler.GetOne)
			playlists.GET("/:id/tracks", playlistTrackHandler.GetPlaylistTrack)
		}

		customers := protected.Group("/customers")
		{
			customers.GET("", customerHandler.GetAll)
			customers.GET("/:id", customerHandler.GetOne)
		}

		invoices := protected.Group("/invoices")
		{
			invoices.GET("", invoiceHandler.GetAll)
			invoices.GET("/:id", invoiceHandler.GetOne)
			invoices.GET("/:id/lines", invoiceHandler.GetInvoiceLines)
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
				"error": "internal server error " + c.Errors.Last().Error(),
			})
		}
	}
}
