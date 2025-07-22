package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"strconv"

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
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Track
// @Router /api/v1/tracks [get]
func (h *TrackHandler) GetAll(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}
	tracks, err := h.Repo.GetTracksPaginated(c.Request.Context(), limit, offset)
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
