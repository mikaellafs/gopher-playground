package server

import (
	"log"

	authmode "gopher-playground/api-auth/pkg/auth/mode"
	"gopher-playground/api-auth/pkg/auth/repository"
	"gopher-playground/api-auth/pkg/config"
	"gopher-playground/api-auth/pkg/http/rest/router"
)

func Start(cfg *config.Configuration) error {
	// Create repos
	userRepo := repository.NewMemoryUserRepository()

	// Get auth mode
	mode, err := authmode.Get(cfg.Auth.Mode)
	if err != nil {
		return err
	}

	// Router config
	rcfg := &router.Config{
		RateLimit:  cfg.Server.RateLimit,
		RetryAfter: cfg.Server.RetryAfter,

		UserRepo: userRepo,

		AuthMode: mode,
	}

	r := router.Initialize(rcfg)

	// Start server
	log.Println("Listening on port", cfg.Server.Port)
	return r.Run(":" + cfg.Server.Port)
}
