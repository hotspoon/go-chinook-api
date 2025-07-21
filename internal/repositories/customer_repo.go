package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type CustomerRepository struct {
	DB *sql.DB
}

func (r *CustomerRepository) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			CustomerId, FirstName, LastName, Company, Address, City, State, Country,
			PostalCode, Phone, Fax, Email, SupportRepId
		FROM Customer
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query customers")
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.CustomerId, &customer.FirstName, &customer.LastName,
			&customer.Company, &customer.Address, &customer.City, &customer.State,
			&customer.Country, &customer.PostalCode, &customer.Phone,
			&customer.Fax, &customer.Email, &customer.SupportRepId); err != nil {
			log.Error().Err(err).Msg("failed to scan customer")
			return nil, fmt.Errorf("error scanning customer: %w", err)
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over customers")
		return nil, fmt.Errorf("error iterating over customers: %w", err)
	}
	return customers, nil
}

func (r *CustomerRepository) GetCustomerByID(ctx context.Context, id int) (models.Customer, error) {
	var customer models.Customer
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			CustomerId, FirstName, LastName, Company, Address, City, State,
			Country, PostalCode, Phone, Fax, Email, SupportRepId
		FROM Customer
		WHERE CustomerId = ?
	`, id).Scan(&customer.CustomerId, &customer.FirstName, &customer.LastName,
		&customer.Company, &customer.Address, &customer.City, &customer.State,
		&customer.Country, &customer.PostalCode, &customer.Phone,
		&customer.Fax, &customer.Email, &customer.SupportRepId)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Customer{}, fmt.Errorf("customer with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query customer by ID")
		return models.Customer{}, fmt.Errorf("error fetching customer by ID: %w", err)
	}
	return customer, nil
}
