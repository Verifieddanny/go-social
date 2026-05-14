package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return nil
}

func (m *MockUserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, exp time.Duration, userID uuid.UUID) error {
	return nil
}

func (m *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) update(ctx context.Context, tx *sql.Tx, user *User) error {
	return nil
}

func (m *MockUserStore) deleteUserInvitations(ctx context.Context, tx *sql.Tx, userID uuid.UUID) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (m *MockUserStore) delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	return nil
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return &User{}, nil
}
