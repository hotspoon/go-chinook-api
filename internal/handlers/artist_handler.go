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

// @Summary Get all artists
// @Description Returns a list of all artists
// @Tags artists
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Artist
// @Router /api/artists [get]
func (h *ArtistHandler) GetAll(c *gin.Context) {
	artists, err := h.Repo.GetAllArtists()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}

// @Summary Get artist by ID
// @Description Returns a single artist by ID
// @Tags artists
// @Produce json
// @Security BearerAuth
// @Param id path int true "Artist ID"
// @Success 200 {object} models.Artist
// @Failure 404 {object} map[string]string
// @Router /api/artists/{id} [get]
func (h *ArtistHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	artist, err := h.Repo.GetArtistByID(utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// @Summary Create a new artist
// @Description Creates a new artist
// @Tags artists
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param artist body models.Artist true "Artist to create"
// @Success 201 {object} models.Artist
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/artists [post]
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

// @Summary Update an artist
// @Description Updates an existing artist by ID
// @Tags artists
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Artist ID"
// @Param artist body models.Artist true "Artist data to update"
// @Success 200 {object} models.Artist
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/artists/{id} [put]
func (h *ArtistHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
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

// @Summary Delete an artist
// @Description Deletes an artist by ID
// @Tags artists
// @Produce json
// @Security BearerAuth
// @Param id path int true "Artist ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/artists/{id} [delete]
func (h *ArtistHandler) Delete(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	if err := h.Repo.DeleteArtist(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
