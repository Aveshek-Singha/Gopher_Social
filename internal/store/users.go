package store

import (
	"context"
	"database/sql"
	"time"
)

type UsersStore struct {
	db *sql.DB
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Do not return password in JSON responses
	CreatedAt time.Time `json:"created_at"`
}

func (u *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
        INSERT INTO users (username, email, password, created_at)
        VALUES ($1, $2, $3, NOW())
        RETURNING id, created_at;
    `

	err := u.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
