package repositories

import (
	"context"
	"database/sql"
	"time"
)

type RefreshToken struct {
    Token     string
    Username  string
    ExpiresAt time.Time
}

type RefreshTokenRepository struct {
    DB *sql.DB
}

func (r *RefreshTokenRepository) Save(ctx context.Context, token, username string, expiresAt time.Time) error {
    _, err := r.DB.ExecContext(
        ctx,
        "INSERT INTO Refresh_Tokens (token, username, expires_at) VALUES (?, ?, ?)",
        token, username, expiresAt,
    )
    return err
}

func (r *RefreshTokenRepository) Get(ctx context.Context, token string) (RefreshToken, error) {
    var rt RefreshToken
    err := r.DB.QueryRowContext(
        ctx,
        "SELECT token, username, expires_at FROM Refresh_Tokens WHERE token = ?",
        token,
    ).Scan(&rt.Token, &rt.Username, &rt.ExpiresAt)
    return rt, err
}

func (r *RefreshTokenRepository) Delete(ctx context.Context, token string) error {
    _, err := r.DB.ExecContext(ctx, "DELETE FROM Refresh_Tokens WHERE token = ?", token)
    return err
}