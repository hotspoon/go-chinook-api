package models

type Track struct {
	TrackId      int     `json:"track_id"`
	Name         string  `json:"name"`
	AlbumId      *int    `json:"album_id,omitempty"`
	MediaTypeId  int     `json:"media_type_id"`
	GenreId      *int    `json:"genre_id,omitempty"`
	Composer     *string `json:"composer,omitempty"`
	Milliseconds int     `json:"milliseconds"`
	Bytes        *int    `json:"bytes,omitempty"`
	UnitPrice    float64 `json:"unit_price"`
}
