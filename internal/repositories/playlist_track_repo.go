package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type PlaylistTrackRepository struct {
	DB *sql.DB
}

func (r *PlaylistTrackRepository) GetTracksByPlaylistID(ctx context.Context, playlistId int) ([]models.Track, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT Track.* FROM Track
		JOIN PlaylistTrack ON Track.TrackId = PlaylistTrack.TrackId
		WHERE PlaylistTrack.PlaylistId = ?
	`, playlistId)
	if err != nil {
		log.Error().Err(err).Int("playlist_id", playlistId).Msg("failed to query tracks for playlist")
		return nil, fmt.Errorf("error fetching tracks for playlist: %w", err)
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
