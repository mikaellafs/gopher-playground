package user

import (
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrUserNotFound      = errors.New("User not found")
)

type Repository interface {
	CreateUser(ctx context.Context, username, password, name string) (*User, error)
	ReadUser(ctx context.Context, username string) (*User, error)
	DeleteUser(ctx context.Context, username string) error
}
