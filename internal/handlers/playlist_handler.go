package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct {
	Repo *repositories.PlaylistRepository
}

// @Summary Get all playlists
// @Description Returns a list of all playlists
// @Tags playlists
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Playlist
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/playlists [get]
func (h *PlaylistHandler) GetAll(c *gin.Context) {
	playlists, err := h.Repo.GetAllPlaylists(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, playlists)
}

// @Summary Get playlist by ID
// @Description Returns a single playlist by ID
// @Tags playlists
// @Produce json
// @Security BearerAuth
// @Param id path int true "Playlist ID"
// @Success 200 {object} models.Playlist
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/playlists/{id} [get]
func (h *PlaylistHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	playlist, err := h.Repo.GetPlaylistByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, playlist)
}
