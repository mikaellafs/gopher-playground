package mode

import (
	"errors"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/config"
	"strings"
)

const (
	modeBasic          string = "basic"
	modeSessionCookies string = "session_cookies"
	modeJwt            string = "jwt"
)

var (
	ErrInvalidAuthMode = errors.New("invalid auth mode")
)

func Get(config *config.Auth, urepo user.Repository, tstore token.TokenStore) (AuthMode, error) {
	mode := strings.ToLower(config.Mode)

	switch mode {
	case modeBasic:
		return NewBasicAuth(urepo), nil
	case modeSessionCookies:
		return NewSessionCookiesAuth(urepo, tstore), nil
	case modeJwt:
		return NewJwtAuth(urepo, config.SigningAlgorithm, config.DurationMinutes, config.RefreshDurationMinutes, config.MaxRefreshMinutes), nil
	}

	return nil, ErrInvalidAuthMode
}
