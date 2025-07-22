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

// GetInvoiceLinesByInvoiceID fetches all invoice lines for a given invoice ID
func (r *InvoiceRepository) GetInvoiceLinesByInvoiceID(ctx context.Context, invoiceID int) ([]models.InvoiceLine, error) {
    rows, err := r.DB.QueryContext(ctx, `
        SELECT InvoiceLineId, InvoiceId, TrackId, UnitPrice, Quantity
        FROM InvoiceLine
        WHERE InvoiceId = ?
    `, invoiceID)
    if err != nil {
        log.Error().Err(err).Msg("failed to query invoice lines")
        return nil, fmt.Errorf("error fetching invoice lines: %w", err)
    }
    defer rows.Close()

    var lines []models.InvoiceLine
    for rows.Next() {
        var line models.InvoiceLine
        if err := rows.Scan(&line.InvoiceLineId, &line.InvoiceId, &line.TrackId, &line.UnitPrice, &line.Quantity); err != nil {
            log.Error().Err(err).Msg("failed to scan invoice line")
            return nil, fmt.Errorf("error scanning invoice line: %w", err)
        }
        lines = append(lines, line)
    }
    if err := rows.Err(); err != nil {
        log.Error().Err(err).Msg("error iterating over invoice lines")
        return nil, fmt.Errorf("error iterating over invoice lines: %w", err)
    }
    return lines, nil
}