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
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, uuid.UUID) (*Post, error)
		Delete(context.Context, uuid.UUID) error
		Update(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, uuid.UUID) (*User, error)
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, uuid.UUID) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerID uuid.UUID, userID uuid.UUID) error
		Unfollow(ctx context.Context, followerID uuid.UUID, userID uuid.UUID) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
