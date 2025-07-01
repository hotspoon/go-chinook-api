package models

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}
