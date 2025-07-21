package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type MediaTypeRepository struct {
	DB *sql.DB
}

func (r *MediaTypeRepository) GetAllMediaTypes(ctx context.Context) ([]models.MediaType, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			MediaTypeId, Name
		FROM MediaType
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query media types")
		return nil, fmt.Errorf("error fetching media types: %w", err)
	}
	defer rows.Close()

	var mediaTypes []models.MediaType
	for rows.Next() {
		var mediaType models.MediaType
		if err := rows.Scan(&mediaType.MediaTypeId, &mediaType.Name); err != nil {
			log.Error().Err(err).Msg("failed to scan media type")
			return nil, fmt.Errorf("error scanning media type: %w", err)
		}
		mediaTypes = append(mediaTypes, mediaType)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over media types")
		return nil, fmt.Errorf("error iterating over media types: %w", err)
	}
	return mediaTypes, nil
}

func (r *MediaTypeRepository) GetMediaTypeByID(ctx context.Context, id int) (models.MediaType, error) {
	var mediaType models.MediaType
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			MediaTypeId, Name
		FROM MediaType
		WHERE MediaTypeId = ?
	`, id).Scan(&mediaType.MediaTypeId, &mediaType.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.MediaType{}, fmt.Errorf("media type with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query media type by ID")
		return models.MediaType{}, fmt.Errorf("error fetching media type by ID: %w", err)
	}
	return mediaType, nil
}
