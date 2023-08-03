package token

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenStore interface {
	Create(token Token) (string, error)
	Read(id string) (*Token, error)
	Delete(id string) error
}
