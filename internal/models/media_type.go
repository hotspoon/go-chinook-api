package models

type MediaType struct {
	MediaTypeId int     `json:"media_type_id"`
	Name        *string `json:"name,omitempty"`
}
