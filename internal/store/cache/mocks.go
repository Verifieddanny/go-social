package cache

import (
	"context"

	"github.com/Verifieddanny/go-social/internal/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) Get(ctx context.Context, id uuid.UUID) (*store.User, error) {
	args := m.Called(id)
	return nil, args.Error(1)
}

func (m *MockUserStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
