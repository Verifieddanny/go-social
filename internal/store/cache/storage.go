package cache

import (
	"context"

	"github.com/Verifieddanny/go-social/internal/store"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Users interface {
		Get(context.Context, uuid.UUID) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage {
		Users: &UserStore{rdb: rbd},
	}
}
