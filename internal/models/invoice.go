package models

import "time"

type Invoice struct {
	InvoiceId         int       `json:"invoice_id"`
	CustomerId        int       `json:"customer_id"`
	InvoiceDate       time.Time `json:"invoice_date"`
	BillingAddress    *string   `json:"billing_address,omitempty"`
	BillingCity       *string   `json:"billing_city,omitempty"`
	BillingState      *string   `json:"billing_state,omitempty"`
	BillingCountry    *string   `json:"billing_country,omitempty"`
	BillingPostalCode *string   `json:"billing_postal_code,omitempty"`
	Total             float64   `json:"total"`
}
