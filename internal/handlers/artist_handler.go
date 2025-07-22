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
// @Router /api/v1/artists [get]
func (h *ArtistHandler) GetAll(c *gin.Context) {
	artists, err := h.Repo.GetAllArtists(c.Request.Context())
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
// @Router /api/v1/artists/{id} [get]
func (h *ArtistHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	artist, err := h.Repo.GetArtistByID(c.Request.Context(), utils.ParseInt(id))
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
// @Router /api/v1/artists [post]
func (h *ArtistHandler) Create(c *gin.Context) {
	var artist models.Artist

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	id, err := h.Repo.CreateArtist(c.Request.Context(), artist)
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
// @Router /api/v1/artists/{id} [put]
func (h *ArtistHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	artist.ID = id
	if err := h.Repo.UpdateArtist(c.Request.Context(), artist); err != nil {
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
// @Router /api/v1/artists/{id} [delete]
func (h *ArtistHandler) Delete(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	// Check if artist exists before deleting
	_, err := h.Repo.GetArtistByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}
	if err := h.Repo.DeleteArtist(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

// @Summary Search artists by name
// @Description Returns artists whose names match the search term
// @Tags artists
// @Produce json
// @Security BearerAuth
// @Param name query string true "Artist name to search for"
// @Success 200 {array} models.Artist
// @Failure 400 {object} models.ErrorResponse
// @Router /api/v1/artists/search [get]
func (h *ArtistHandler) SearchByName(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: "name query parameter is required"})
        return
    }
    artists, err := h.Repo.SearchArtistsByName(c.Request.Context(), name)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    if len(artists) == 0 {
        c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{Error: "no artists found"})
        return
    }
    c.JSON(http.StatusOK, artists)
}