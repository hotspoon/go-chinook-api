package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaTypeHandler struct {
	Repo *repositories.MediaTypeRepository
}

// @Summary Get all media types
// @Description Returns a list of all media types
// @Tags media_types
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MediaType
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/media_types [get]
func (h *MediaTypeHandler) GetAll(c *gin.Context) {
	mediaTypes, err := h.Repo.GetAllMediaTypes(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediaTypes)
}

// @Summary Get media type by ID
// @Description Returns a single media type by ID
// @Tags media_types
// @Produce json
// @Security BearerAuth
// @Param id path int true "Media Type ID"
// @Success 200 {object} models.MediaType
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/media_types/{id} [get]
func (h *MediaTypeHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	mediaType, err := h.Repo.GetMediaTypeByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediaType)
}
