package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateEmail    = errors.New("a user with that email already exist")
	ErrDuplicateUsername = errors.New("a user with that username already exist")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, uuid.UUID) (*Post, error)
		Delete(context.Context, uuid.UUID) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, uuid.UUID, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, uuid.UUID) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, uuid.UUID) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, uuid.UUID) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerID uuid.UUID, userID uuid.UUID) error
		Unfollow(ctx context.Context, followerID uuid.UUID, userID uuid.UUID) error
	}
	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
		Roles:     &RolesStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
