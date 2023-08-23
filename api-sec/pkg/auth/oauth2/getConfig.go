package oauth2

import (
	"os"

	"gopher-playground/api-sec/pkg/env"

	"golang.org/x/oauth2"
)

func GetConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     os.Getenv(env.GOOGLE_CLIENT_ID),
		ClientSecret: os.Getenv(env.GOOGLE_CLIENT_SECRET),
		RedirectURL:  os.Getenv(env.OAUTH2_REDIRECT_URL),
		Scopes:       []string{"profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv(env.OAUTH2_AUTH_URL),
			TokenURL: os.Getenv(env.OAUTH2_TOKEN_URL),
		},
	}
}
