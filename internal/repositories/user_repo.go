package repositories

import (
	"chinook-api/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (int64, error) {
    result, err := r.DB.ExecContext(
        ctx,
        "INSERT INTO User (Username, Email, Password) VALUES (?, ?, ?)",
        user.Username, user.Email, user.Password,
    )
    if err != nil {
        log.Error().Err(err).Msg("Error creating user")
        return 0, fmt.Errorf("error creating user: %w", err)
    }
    return result.LastInsertId()
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
    var user models.User
    err := r.DB.QueryRowContext(
        ctx,
        "SELECT UserId, Username, Email, Password FROM User WHERE Username = ?",
        username,
    ).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    if err != nil {
        log.Error().Err(err).Msg("User not found")
        return user, fmt.Errorf("user not found")
    }
    return user, nil
}