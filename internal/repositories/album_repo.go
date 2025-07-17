package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"chinook-api/internal/models"

	"github.com/rs/zerolog/log"
)

type AlbumRepository struct {
    DB *sql.DB
}

func (r *AlbumRepository) GetAllAlbums(ctx context.Context) ([]models.Album, error) {
    rows, err := r.DB.QueryContext(ctx, "SELECT AlbumId, Title, ArtistId FROM Album")
    if err != nil {
        log.Error().Err(err).Msg("failed to query albums")
        return nil, fmt.Errorf("error fetching albums: %w", err)
    }
    defer rows.Close()

    var albums []models.Album
    for rows.Next() {
        var album models.Album
        if err := rows.Scan(&album.ID, &album.Title, &album.ArtistID); err != nil {
            log.Error().Err(err).Msg("failed to scan album")
            return nil, fmt.Errorf("error scanning album: %w", err)
        }
        albums = append(albums, album)
    }
    return albums, nil
}

func (r *AlbumRepository) GetAlbumByID(ctx context.Context, id int) (models.Album, error) {
    var album models.Album
    err := r.DB.QueryRowContext(ctx, "SELECT AlbumId, Title, ArtistId FROM Album WHERE AlbumId = ?", id).Scan(&album.ID, &album.Title, &album.ArtistID)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Warn().Int("id", id).Msg("Album not found")
            return album, fmt.Errorf("album not found")
        }
        log.Error().Err(err).Int("id", id).Msg("Database error fetching album")
        return album, fmt.Errorf("database error: %w", err)
    }
    log.Debug().Int("id", album.ID).Msg("Fetched album by ID")
    return album, nil
}

func (r *AlbumRepository) CreateAlbum(ctx context.Context, album models.Album) (int64, error) {
    log.Debug().Msg("Creating album")
    result, err := r.DB.ExecContext(ctx, "INSERT INTO Album (Title, ArtistId) VALUES (?, ?)", album.Title, album.ArtistID)
    if err != nil {
        log.Error().Err(err).Msg("failed to create album")
        return 0, fmt.Errorf("error creating album: %w", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        log.Error().Err(err).Msg("failed to get last inserted ID")
        return 0, fmt.Errorf("error getting last inserted ID: %w", err)
    }
    log.Debug().Int64("id", id).Msg("Created album")
    return id, nil
}

func (r *AlbumRepository) UpdateAlbum(ctx context.Context, album models.Album) error {
    _, err := r.DB.ExecContext(ctx, "UPDATE Album SET Title = ?, ArtistId = ? WHERE AlbumId = ?", album.Title, album.ArtistID, album.ID)
    if err != nil {
        log.Error().Err(err).Msg("failed to update album")
        return fmt.Errorf("error updating album: %w", err)
    }
    log.Debug().Int("id", album.ID).Msg("Updated album")
    return nil
}

func (r *AlbumRepository) DeleteAlbum(ctx context.Context, id int) error {
    log.Debug().Int("id", id).Msg("Deleting album")
    result, err := r.DB.ExecContext(ctx, "DELETE FROM Album WHERE AlbumId = ?", id)
    if err != nil {
        log.Error().Err(err).Msg("failed to delete album")
        return fmt.Errorf("error deleting album: %w", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Error().Err(err).Msg("failed to get rows affected")
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    if rowsAffected == 0 {
        log.Warn().Int("id", id).Msg("Album not found")
        return fmt.Errorf("album not found")
    }
    log.Debug().Int("id", id).Msg("Deleted album")
    return nil
}