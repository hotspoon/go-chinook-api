package models

type Artist struct {
	ID   int
	Name string
}

type PaginatedArtistsResponse struct {
    Data    []Artist `json:"data"`
    Total   int      `json:"total"`
    Limit   int      `json:"limit"`
    Offset  int      `json:"offset"`
    HasMore bool     `json:"hasMore"`
}