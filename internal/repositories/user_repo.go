package repositories

import (
	"chinook-api/internal/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	result, err := r.DB.Exec(
		"INSERT INTO User (Username, Email, Password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password,
	)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return result.LastInsertId()
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT UserId, Username, Email, Password FROM User WHERE Username = ?",
		username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}
