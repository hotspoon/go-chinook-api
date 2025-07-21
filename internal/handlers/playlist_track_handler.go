package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaylistTrackHandler struct {
	Repo *repositories.PlaylistTrackRepository
}

// @Summary Get all tracks in a playlist
// @Description Returns a list of all tracks in a specific playlist
// @Tags playlist_tracks
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "Playlist ID"
// @Success 200 {array} models.Track
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/playlists/{playlistId}/tracks [get]
func (h *PlaylistTrackHandler) GetPlaylistTrack(c *gin.Context) {
	playlistId := utils.ParseInt(c.Param("id"))
	// print playlistId for debugging
	fmt.Println("Playlist ID:", playlistId)
	if playlistId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	tracks, err := h.Repo.GetTracksByPlaylistID(c.Request.Context(), playlistId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(tracks) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No tracks found for this playlist"})
		return
	}

	c.JSON(http.StatusOK, tracks)

}
