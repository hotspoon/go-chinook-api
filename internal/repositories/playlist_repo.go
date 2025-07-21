package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type PlaylistRepository struct {
	DB *sql.DB
}

func (r *PlaylistRepository) GetAllPlaylists(ctx context.Context) ([]models.Playlist, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			PlaylistId, Name
		FROM Playlist
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query playlists")
		return nil, fmt.Errorf("error fetching playlists: %w", err)
	}
	defer rows.Close()

	var playlists []models.Playlist
	for rows.Next() {
		var playlist models.Playlist
		if err := rows.Scan(&playlist.PlaylistId, &playlist.Name); err != nil {
			log.Error().Err(err).Msg("failed to scan playlist")
			return nil, fmt.Errorf("error scanning playlist: %w", err)
		}
		playlists = append(playlists, playlist)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over playlists")
		return nil, fmt.Errorf("error iterating over playlists: %w", err)
	}
	return playlists, nil
}

func (r *PlaylistRepository) GetPlaylistByID(ctx context.Context, id int) (models.Playlist, error) {
	var playlist models.Playlist
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			PlaylistId, Name
		FROM Playlist
		WHERE PlaylistId = ?
	`, id).Scan(&playlist.PlaylistId, &playlist.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Playlist{}, fmt.Errorf("playlist with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query playlist by ID")
		return models.Playlist{}, fmt.Errorf("error fetching playlist by ID: %w", err)
	}
	return playlist, nil
}
