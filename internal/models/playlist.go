package models

type Playlist struct {
	PlaylistId int     `json:"playlist_id"`
	Name       *string `json:"name,omitempty"`
}
