package handlers

import (
	"net/http"

	"chinook-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ArtistHandler struct {
	Repo *repositories.ArtistRepository
}

func (h *ArtistHandler) GetAll(c *gin.Context) {
	artists, err := h.Repo.GetAllArtists()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}

func (h *ArtistHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	artist, err := h.Repo.GetArtistByID(ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artist)
}
