package models

type PlaylistTrack struct {
	PlaylistId int `json:"playlist_id"`
	TrackId    int `json:"track_id"`
}
