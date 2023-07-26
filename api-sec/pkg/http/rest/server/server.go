package server

import (
	"log"
	"os"

	"gopher-playground/api-sec/pkg/auth/accesscontrol"
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	authrepo "gopher-playground/api-sec/pkg/auth/repository"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/config"
	"gopher-playground/api-sec/pkg/env"
	"gopher-playground/api-sec/pkg/http/rest/router"
	logrepo "gopher-playground/api-sec/pkg/log/repository"
)

func Start(cfg *config.Configuration) error {
	// Create repos
	userRepo := authrepo.NewMemoryUserRepository()
	logRepo := logrepo.NewMemoryAuditLogRepository()

	tokenStore := authrepo.NewInMemoryTokenStore(token.GenerateRandTokenString)

	// Get auth mode
	mode, err := authmode.Get(cfg.Auth, userRepo, tokenStore)
	if err != nil {
		return err
	}

	// Create access control
	ac := accesscontrol.NewFileCasbinRBAC(os.Getenv(env.CASBIN_CONFIG_PATH), os.Getenv(env.CASBIN_POLICY_PATH))
	err = ac.LoadPolicy()
	if err != nil {
		return err
	}

	// Router config
	rcfg := &router.Config{
		RateLimit:  cfg.Server.RateLimit,
		RetryAfter: cfg.Server.RetryAfter,

		UserRepo: userRepo,
		LogRepo:  logRepo,

		AuthMode:      mode,
		AccessControl: ac,
		TokenStore:    tokenStore,
	}

	r := router.Initialize(rcfg)

	// Start server
	addr := ":" + cfg.Server.Port

	if cfg.Server.Https.Enable {
		log.Println("Listening HTTPS on", addr)
		return r.RunTLS(addr, cfg.Server.Https.CertPath, cfg.Server.Https.KeyPath)
	}

	log.Println("Listening HTTP on", addr)
	return r.Run(addr)
}
