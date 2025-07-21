package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Repo *repositories.CustomerRepository
}

// @Summary Get all customers
// @Description Returns a list of all customers
// @Tags customers
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Customer
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/customers [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.Repo.GetAllCustomers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// @Summary Get customer by ID
// @Description Returns a single customer by ID
// @Tags customers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Customer ID"
// @Success 200 {object} models.Customer
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/customers/{id} [get]
func (h *CustomerHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.Repo.GetCustomerByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}
