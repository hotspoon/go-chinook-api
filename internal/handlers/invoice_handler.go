package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	Repo *repositories.InvoiceRepository
}

// @Summary Get all invoices
// @Description Returns a list of all invoices
// @Tags invoices
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Invoice
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/invoices [get]
func (h *InvoiceHandler) GetAll(c *gin.Context) {
	invoices, err := h.Repo.GetAllInvoices(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

// @Summary Get invoice by ID
// @Description Returns a single invoice by ID
// @Tags invoices
// @Produce json
// @Security BearerAuth
// @Param id path int true "Invoice ID"
// @Success 200 {object} models.Invoice
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/invoices/{id} [get]
func (h *InvoiceHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	invoice, err := h.Repo.GetInvoiceByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoice)
}
