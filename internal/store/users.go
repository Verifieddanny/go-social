package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt string    `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (first_name, last_name, email, Password, username) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Username,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetByID(ctx context.Context, postId uuid.UUID) (*User, error) {
	query := `
	SELECT id, email, Password, username, created_at 
	FROM users
	WHERE ID = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		postId,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.CreatedAt,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return user, nil
}
