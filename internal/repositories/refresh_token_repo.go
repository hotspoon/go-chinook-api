package repositories

import (
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

func (r *RefreshTokenRepository) Save(token, username string, expiresAt time.Time) error {
	_, err := r.DB.Exec(
		"INSERT INTO Refresh_Tokens (token, username, expires_at) VALUES (?, ?, ?)",
		token, username, expiresAt,
	)
	return err
}

func (r *RefreshTokenRepository) Get(token string) (RefreshToken, error) {
	var rt RefreshToken
	err := r.DB.QueryRow(
		"SELECT token, username, expires_at FROM Refresh_Tokens WHERE token = ?",
		token,
	).Scan(&rt.Token, &rt.Username, &rt.ExpiresAt)
	return rt, err
}

func (r *RefreshTokenRepository) Delete(token string) error {
	_, err := r.DB.Exec("DELETE FROM Refresh_Tokens WHERE token = ?", token)
	return err
}
