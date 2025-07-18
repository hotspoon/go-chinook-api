package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type GenreRepository struct {
	DB *sql.DB
}

func (r *GenreRepository) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			GenreId, Name
		FROM Genre
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query genres")
		return nil, fmt.Errorf("error fetching genres: %w", err)
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.GenreId, &genre.Name); err != nil {
			log.Error().Err(err).Msg("failed to scan genre")
			return nil, fmt.Errorf("error scanning genre: %w", err)
		}
		genres = append(genres, genre)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over genres")
		return nil, fmt.Errorf("error iterating over genres: %w", err)
	}
	return genres, nil
}

func (r *GenreRepository) GetGenreByID(ctx context.Context, id int) (models.Genre, error) {
	var genre models.Genre
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			GenreId, Name
		FROM Genre
		WHERE GenreId = ?
	`, id).Scan(&genre.GenreId, &genre.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Genre{}, fmt.Errorf("genre with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query genre by ID")
		return models.Genre{}, fmt.Errorf("error fetching genre by ID: %w", err)
	}
	return genre, nil
}
