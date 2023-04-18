package repository

import (
	"context"
	"sync"

	"gopher-playground/api-auth/pkg/user"
)

type memoryUser struct {
	username string
	pwHash   string
}

type MemoryUserRepository struct {
	users map[string]*memoryUser
	lock  *sync.RWMutex
}

var _ user.Repository = (*MemoryUserRepository)(nil)

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: map[string]*memoryUser{},
		lock:  &sync.RWMutex{},
	}
}

func (r *MemoryUserRepository) CreateUser(ctx context.Context, username, password string) (*user.User, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	// Save to repo
	if r.users[username] != nil {
		return nil, user.ErrUserAlreadyExists
	}

	r.users[username] = &memoryUser{
		username: username,
		pwHash:   password,
	}

	return &user.User{
		Username: username,
		Password: password,
	}, nil
}

func (r *MemoryUserRepository) ReadUser(ctx context.Context, username string) (*user.User, error) {
	if r.users[username] == nil {
		return nil, user.ErrUserNotFound
	}

	return &user.User{
		Username: username,
		Password: r.users[username].pwHash,
	}, nil
}

func (r *MemoryUserRepository) DeleteUser(ctx context.Context, username string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.users[username] == nil {
		return user.ErrUserNotFound
	}

	r.users[username] = nil

	return nil
}
