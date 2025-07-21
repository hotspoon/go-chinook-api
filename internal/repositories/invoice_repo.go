package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type InvoiceRepository struct {
	DB *sql.DB
}

func (r *InvoiceRepository) GetAllInvoices(ctx context.Context) ([]models.Invoice, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			InvoiceId, CustomerId, InvoiceDate, BillingAddress, BillingCity,
			BillingState, BillingCountry, BillingPostalCode, Total
		FROM Invoice
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query invoices")
		return nil, fmt.Errorf("error fetching invoices: %w", err)
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		if err := rows.Scan(&invoice.InvoiceId, &invoice.CustomerId, &invoice.InvoiceDate,
			&invoice.BillingAddress, &invoice.BillingCity, &invoice.BillingState,
			&invoice.BillingCountry, &invoice.BillingPostalCode, &invoice.Total); err != nil {
			log.Error().Err(err).Msg("failed to scan invoice")
			return nil, fmt.Errorf("error scanning invoice: %w", err)
		}
		invoices = append(invoices, invoice)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over invoices")
		return nil, fmt.Errorf("error iterating over invoices: %w", err)
	}
	return invoices, nil
}

func (r *InvoiceRepository) GetInvoiceByID(ctx context.Context, id int) (models.Invoice, error) {
	var invoice models.Invoice
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			InvoiceId, CustomerId, InvoiceDate, BillingAddress, BillingCity,
			BillingState, BillingCountry, BillingPostalCode, Total
		FROM Invoice
		WHERE InvoiceId = ?
	`, id).Scan(&invoice.InvoiceId, &invoice.CustomerId, &invoice.InvoiceDate,
		&invoice.BillingAddress, &invoice.BillingCity, &invoice.BillingState,
		&invoice.BillingCountry, &invoice.BillingPostalCode, &invoice.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Invoice{}, fmt.Errorf("invoice with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to scan invoice by ID")
		return models.Invoice{}, fmt.Errorf("error scanning invoice by ID: %w", err)
	}
	return invoice, nil
}
