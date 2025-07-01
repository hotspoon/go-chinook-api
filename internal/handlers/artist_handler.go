package handlers

import (
	"net/http"

	"chinook-api/internal/models"
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"

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
	artist, err := h.Repo.GetArtistByID(utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (h *ArtistHandler) Create(c *gin.Context) {
	var artist models.Artist

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	id, err := h.Repo.CreateArtist(artist)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	artist.ID = int(id)

	c.JSON(http.StatusCreated, artist)
}

func (h *ArtistHandler) Update(c *gin.Context) {
	id := ParseInt(c.Param("id"))
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	artist.ID = id
	if err := h.Repo.UpdateArtist(artist); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (h *ArtistHandler) Delete(c *gin.Context) {
	id := ParseInt(c.Param("id"))
	if err := h.Repo.DeleteArtist(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
