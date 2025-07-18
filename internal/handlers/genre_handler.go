package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	Repo *repositories.GenreRepository
}

// @Summary Get all genres
// @Description Returns a list of all genres
// @Tags genres
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Genre
// @Router /api/v1/genres [get]
func (h *GenreHandler) GetAll(c *gin.Context) {
	genres, err := h.Repo.GetAllGenres(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genres)
}

// @Summary Get genre by ID
// @Description Returns a single genre by ID
// @Tags genres
// @Produce json
// @Security BearerAuth
// @Param id path int true "Genre ID"
// @Success 200 {object} models.Genre
// @Router /api/v1/genres/{id} [get]
func (h *GenreHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	genre, err := h.Repo.GetGenreByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genre)
}
