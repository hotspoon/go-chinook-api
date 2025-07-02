package repositories

import (
	"database/sql"
	"fmt"

	"chinook-api/internal/models"

	"github.com/rs/zerolog/log"
)

type ArtistRepository struct {
	DB *sql.DB
}

func (r *ArtistRepository) GetAllArtists() ([]models.Artist, error) {
	log.Debug().Msg("Fetching all artists")
	rows, err := r.DB.Query("SELECT ArtistId, Name FROM Artist")
	if err != nil {
		log.Error().Err(err).Msg("Error fetching artists")
		return nil, fmt.Errorf("error fetching artists: %w", err)
	}
	defer rows.Close()

	var artists []models.Artist
	for rows.Next() {
		var artist models.Artist
		if err := rows.Scan(&artist.ID, &artist.Name); err != nil {
			log.Error().Err(err).Msg("Error scanning artist row")
			return nil, fmt.Errorf("error scanning artist row: %w", err)
		}
		artists = append(artists, artist)
	}
	log.Debug().Int("count", len(artists)).Msg("Fetched artists")
	return artists, nil
}

func (r *ArtistRepository) GetArtistByID(id int) (models.Artist, error) {
	log.Debug().Int("id", id).Msg("Fetching artist by ID")
	var artist models.Artist
	err := r.DB.QueryRow("SELECT ArtistId, Name FROM Artist WHERE ArtistId = ?", id).
		Scan(&artist.ID, &artist.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Int("id", id).Msg("Artist not found")
			return artist, fmt.Errorf("artist not found")
		}
		log.Error().Err(err).Int("id", id).Msg("Database error fetching artist")
		return artist, fmt.Errorf("database error: %w", err)
	}

	log.Debug().Int("id", artist.ID).Msg("Fetched artist by ID")
	return artist, nil
}

func (r *ArtistRepository) CreateArtist(artist models.Artist) (int64, error) {
	log.Debug().Str("name", artist.Name).Msg("Creating artist")
	result, err := r.DB.Exec("INSERT INTO Artist (Name) VALUES (?)", artist.Name)
	if err != nil {
		log.Error().Err(err).Str("name", artist.Name).Msg("Error creating artist")
		return 0, fmt.Errorf("error creating artist: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("Error getting last insert id")
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}

	log.Info().Int64("id", id).Str("name", artist.Name).Msg("Artist created")
	return id, nil
}

func (r *ArtistRepository) UpdateArtist(artist models.Artist) error {
	log.Debug().Int("id", artist.ID).Str("name", artist.Name).Msg("Updating artist")
	_, err := r.DB.Exec("UPDATE Artist SET Name = ? WHERE ArtistId = ?", artist.Name, artist.ID)
	if err != nil {
		log.Error().Err(err).Int("id", artist.ID).Msg("Error updating artist")
		return fmt.Errorf("error updating artist: %w", err)
	}
	log.Info().Int("id", artist.ID).Msg("Artist updated")
	return nil
}

func (r *ArtistRepository) DeleteArtist(id int) error {
	log.Debug().Int("id", id).Msg("Deleting artist")
	_, err := r.DB.Exec("DELETE FROM Artist WHERE ArtistId = ?", id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("Error deleting artist")
		return fmt.Errorf("error deleting artist: %w", err)
	}
	log.Info().Int("id", id).Msg("Artist deleted")
	return nil
}
