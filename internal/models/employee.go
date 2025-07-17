package models

import (
	"chinook-api/internal/utils"
)

type Employee struct {
	EmployeeId int             `json:"employee_id"`
	LastName   string          `json:"last_name"`
	FirstName  string          `json:"first_name"`
	Title      *string         `json:"title,omitempty"`
	ReportsTo  *int            `json:"reports_to,omitempty"`
	BirthDate  *utils.DateOnly `json:"BirthDate,omitempty"`
	HireDate   *utils.DateOnly `json:"HireDate,omitempty"`
	Address    *string         `json:"address,omitempty"`
	City       *string         `json:"city,omitempty"`
	State      *string         `json:"state,omitempty"`
	Country    *string         `json:"country,omitempty"`
	PostalCode *string         `json:"postal_code,omitempty"`
	Phone      *string         `json:"phone,omitempty"`
	Fax        *string         `json:"fax,omitempty"`
	Email      *string         `json:"email,omitempty"`
}
