package repositories

import (
	"chinook-api/internal/models"
	"database/sql"
	"fmt"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func (r *EmployeeRepository) CreateEmployee(emp models.Employee) (int64, error) {
	result, err := r.DB.Exec(
		"INSERT INTO Employee (FirstName, LastName, Title, Email) VALUES (?, ?, ?, ?)",
		emp.FirstName, emp.LastName, emp.Title, emp.Email,
	)
	if err != nil {
		return 0, fmt.Errorf("error creating employee: %w", err)
	}
	return result.LastInsertId()
}
