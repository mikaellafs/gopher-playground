package mode

import (
	"errors"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
	"strings"
)

const (
	basic          string = "basic"
	sessionCookies string = "session_cookies"
)

var (
	ErrInvalidAuthMode = errors.New("invalid auth mode")
)

func Get(mode string, urepo user.Repository, tstore token.TokenStore) (AuthMode, error) {
	mode = strings.ToLower(mode)

	switch mode {
	case basic:
		return NewBasicAuth(urepo), nil
	case sessionCookies:
		return NewSessionCookiesAuth(urepo, tstore), nil
	}

	return nil, ErrInvalidAuthMode
}
