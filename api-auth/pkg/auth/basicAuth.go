package auth

import (
	"context"
	"encoding/base64"
	"strings"

	"gopher-playground/api-auth/pkg/user"
)

type BasicAuth struct {
}

var _ AuthMode = (*BasicAuth)(nil)

func (a *BasicAuth) Authenticate(authHeader string, userRepo user.Repository) (string, error) {
	if !strings.HasPrefix(authHeader, "Basic ") {
		return "", ErrInvalidAuthHeader
	}

	// Decode credentials
	token := strings.Replace(authHeader, "Basic ", "", 1)
	credentials, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", ErrInvalidAuthHeader
	}

	credArr := strings.Split(string(credentials), ":")
	if len(credArr) != 2 {
		return "", ErrInvalidAuthHeader
	}

	// Check credentials
	username, password := credArr[0], credArr[1]
	user, err := userRepo.ReadUser(context.TODO(), username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !CompareEqual(user.Password, password) {
		return "", ErrInvalidCredentials
	}

	return username, nil
}
