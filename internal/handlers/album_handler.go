package handlers

import (
	"net/http"

	"chinook-api/internal/models"
	_ "chinook-api/internal/models"
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	_ "chinook-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	Repo *repositories.AlbumRepository
}

// @Summary Get all albums
// @Description Returns a list of all albums
// @Tags albums
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Album
// @Router /api/v1/albums [get]
func (h *AlbumHandler) GetAll(c *gin.Context) {
	albums, err := h.Repo.GetAllAlbums()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, albums)
}

// @Summary Get album by ID
// @Description Returns a single album by ID
// @Tags albums
// @Produce json
// @Security BearerAuth
// @Param id path int true "Album ID"
// @Success 200 {object} models.Album
// @Failure 404 {object} map[string]string
// @Router /api/v1/albums/{id} [get]
func (h *AlbumHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	album, err := h.Repo.GetAlbumByID(utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) Create(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Repo.CreateAlbum(album)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	album.ID = int(id)

	c.JSON(http.StatusCreated, album)
}

func (h *AlbumHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	album.ID = id
	if err := h.Repo.UpdateAlbum(album); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteAlbum(utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
