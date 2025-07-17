package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type TrackRepository struct {
	DB *sql.DB
}

func (r *TrackRepository) GetAllTracks(ctx context.Context) ([]models.Track, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			TrackId, Name, AlbumId, MediaTypeId, GenreId, Composer,
			Milliseconds, Bytes, UnitPrice
		FROM Track
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query tracks")
		return nil, fmt.Errorf("error fetching tracks: %w", err)
	}
	defer rows.Close()

	var tracks []models.Track
	for rows.Next() {
		var track models.Track
		if err := rows.Scan(
			&track.TrackId,
			&track.Name,
			&track.AlbumId,
			&track.MediaTypeId,
			&track.GenreId,
			&track.Composer,
			&track.Milliseconds,
			&track.Bytes,
			&track.UnitPrice,
		); err != nil {
			log.Error().Err(err).Msg("failed to scan track")
			return nil, fmt.Errorf("error scanning track: %w", err)
		}
		tracks = append(tracks, track)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over tracks")
		return nil, fmt.Errorf("error iterating over tracks: %w", err)
	}
	return tracks, nil
}

func (r *TrackRepository) GetTrackByID(ctx context.Context, id int) (models.Track, error) {
	var track models.Track
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			TrackId, Name, AlbumId, MediaTypeId, GenreId, Composer,
			Milliseconds, Bytes, UnitPrice
		FROM Track
		WHERE TrackId = ?
	`, id).Scan(
		&track.TrackId,
		&track.Name,
		&track.AlbumId,
		&track.MediaTypeId,
		&track.GenreId,
		&track.Composer,
		&track.Milliseconds,
		&track.Bytes,
		&track.UnitPrice,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Int("id", id).Msg("Track not found")
			return track, fmt.Errorf("track not found")
		}
		log.Error().Err(err).Int("id", id).Msg("Database error fetching track")
		return track, fmt.Errorf("database error: %w", err)
	}
	log.Debug().Int("id", track.TrackId).Msg("Fetched track by ID")
	return track, nil
}
