package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, emp models.Employee) (int64, error) {
	result, err := r.DB.ExecContext(
		ctx,
		"INSERT INTO Employee (FirstName, LastName, Title, Email) VALUES (?, ?, ?, ?)",
		emp.FirstName, emp.LastName, emp.Title, emp.Email,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating employee")
		return 0, fmt.Errorf("error creating employee: %w", err)
	}
	return result.LastInsertId()
}

func (r *EmployeeRepository) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			EmployeeId, LastName, FirstName, Title, ReportsTo, BirthDate, HireDate,
			Address, City, State, Country, PostalCode, Phone, Fax, Email
		FROM Employee
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query employees")
		return nil, fmt.Errorf("error fetching employees: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(
			&employee.EmployeeId,
			&employee.LastName,
			&employee.FirstName,
			&employee.Title,
			&employee.ReportsTo,
			&employee.BirthDate,
			&employee.HireDate,
			&employee.Address,
			&employee.City,
			&employee.State,
			&employee.Country,
			&employee.PostalCode,
			&employee.Phone,
			&employee.Fax,
			&employee.Email,
		); err != nil {
			log.Error().Err(err).Msg("failed to scan employee")
			return nil, fmt.Errorf("error scanning employee: %w", err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("rows error")
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return employees, nil
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (models.Employee, error) {
	var employee models.Employee
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			EmployeeId, LastName, FirstName, Title, ReportsTo, BirthDate, HireDate,
			Address, City, State, Country, PostalCode, Phone, Fax, Email
		FROM Employee
		WHERE EmployeeId = ?
	`).Scan(
		&employee.EmployeeId,
		&employee.LastName,
		&employee.FirstName,
		&employee.Title,
		&employee.ReportsTo,
		&employee.BirthDate,
		&employee.HireDate,
		&employee.Address,
		&employee.City,
		&employee.State,
		&employee.Country,
		&employee.PostalCode,
		&employee.Phone,
		&employee.Fax,
		&employee.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Int("id", id).Msg("Employee not found")
			return employee, fmt.Errorf("employee not found")
		}
		log.Error().Err(err).Int("id", id).Msg("Database error fetching employee")
		return employee, fmt.Errorf("database error: %w", err)
	}
	log.Debug().Int("id", employee.EmployeeId).Msg("Fetched employee by ID")
	return employee, nil
}
