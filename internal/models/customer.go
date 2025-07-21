package models

type Customer struct {
	CustomerId   int     `json:"customer_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Company      *string `json:"company,omitempty"`
	Address      *string `json:"address,omitempty"`
	City         *string `json:"city,omitempty"`
	State        *string `json:"state,omitempty"`
	Country      *string `json:"country,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	Fax          *string `json:"fax,omitempty"`
	Email        string  `json:"email"`
	SupportRepId *int    `json:"support_rep_id,omitempty"`
}
