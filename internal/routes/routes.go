package routes

import (
	"chinook-api/internal/handlers"
	"chinook-api/internal/repositories"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	artistRepo := &repositories.ArtistRepository{DB: db}
	artistHandler := &handlers.ArtistHandler{Repo: artistRepo}

	api := r.Group("/api")
	{
		artists := api.Group("/artists")
		{
			artists.GET("", artistHandler.GetAll)
			artists.GET("/:id", artistHandler.GetOne)
			artists.POST("", artistHandler.Create)
			artists.PUT(":id", artistHandler.Update)
			artists.DELETE(":id", artistHandler.Delete)
		}
	}
}
