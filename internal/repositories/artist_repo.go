package repositories

import (
	"database/sql"
	"fmt"

	"chinook-api/internal/models"
)

type ArtistRepository struct {
	DB *sql.DB
}

func (r *ArtistRepository) GetAllArtists() ([]models.Artist, error) {
	rows, err := r.DB.Query("SELECT ArtistId, Name FROM Artist")
	if err != nil {
		return nil, fmt.Errorf("error fetching artists: %w", err)
	}
	defer rows.Close()

	var artists []models.Artist
	for rows.Next() {
		var artist models.Artist
		if err := rows.Scan(&artist.ID, &artist.Name); err != nil {
			return nil, fmt.Errorf("error scanning artist row: %w", err)
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) GetArtistByID(id int) (models.Artist, error) {
	var artist models.Artist
	err := r.DB.QueryRow("SELECT ArtistId, Name FROM Artist WHERE ArtistId = ?", id).
		Scan(&artist.ID, &artist.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return artist, fmt.Errorf("artist not found")
		}
		return artist, fmt.Errorf("database error: %w", err)
	}

	return artist, nil
}

func (r *ArtistRepository) CreateArtist(artist models.Artist) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO Artist (Name) VALUES (?)", artist.Name)
	if err != nil {
		return 0, fmt.Errorf("error creating artist: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}

	return id, nil
}

func (r *ArtistRepository) UpdateArtist(artist models.Artist) error {
	_, err := r.DB.Exec("UPDATE Artist SET Name = ? WHERE ArtistId = ?", artist.Name, artist.ID)
	if err != nil {
		return fmt.Errorf("error updating artist: %w", err)
	}
	return nil
}

func (r *ArtistRepository) DeleteArtist(id int) error {
	_, err := r.DB.Exec("DELETE FROM Artist WHERE ArtistId = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting artist: %w", err)
	}
	return nil
}
