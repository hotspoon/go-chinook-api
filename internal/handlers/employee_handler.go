package handlers

import (
	"chinook-api/internal/repositories"
	"chinook-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	Repo *repositories.EmployeeRepository
}

// @Summary Get all employees
// @Description Returns a list of all employees
// @Tags employees
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Employee
// @Router /api/v1/employees [get]
func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.Repo.GetAllEmployees(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// @Summary Get employee by ID
// @Description Returns a single employee by ID
// @Tags employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Router /api/v1/employees/{id} [get]
func (h *EmployeeHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	employee, err := h.Repo.GetEmployeeByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}
