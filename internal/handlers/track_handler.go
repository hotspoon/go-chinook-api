package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrackHandler struct {
	Repo *repositories.TrackRepository
}

// @Summary Get all tracks
// @Description Returns a list of all tracks
// @Tags tracks
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Track
// @Router /api/v1/tracks [get]
func (h *TrackHandler) GetAll(c *gin.Context) {
	tracks, err := h.Repo.GetAllTracks(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tracks)
}

// @Summary Get track by ID
// @Description Returns a single track by ID
// @Tags tracks
// @Produce json
// @Security BearerAuth
// @Param id path int true "Track ID"
// @Success 200 {object} models.Track
// @Router /api/v1/tracks/{id} [get]
func (h *TrackHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	track, err := h.Repo.GetTrackByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, track)
}
