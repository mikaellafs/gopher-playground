package mode

import (
	"context"
	"encoding/base64"
	"strings"

	"gopher-playground/api-sec/pkg/auth"
	"gopher-playground/api-sec/pkg/auth/user"
)

type BasicAuth struct {
}

var _ AuthMode = (*BasicAuth)(nil)

func (a *BasicAuth) Authenticate(authHeader string, userRepo user.Repository) (*user.User, error) {
	if !strings.HasPrefix(authHeader, "Basic ") {
		return nil, ErrInvalidAuthHeader
	}

	// Decode credentials
	token := strings.Replace(authHeader, "Basic ", "", 1)
	credentials, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrInvalidAuthHeader
	}

	credArr := strings.Split(string(credentials), ":")
	if len(credArr) != 2 {
		return nil, ErrInvalidAuthHeader
	}

	// Check credentials
	username, password := credArr[0], credArr[1]
	user, err := userRepo.ReadUser(context.TODO(), username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !auth.CompareEqual(user.Password, password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}