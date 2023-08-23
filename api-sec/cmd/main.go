package main

import (
	"log"
	"os"

	"gopher-playground/api-sec/pkg/config"
	"gopher-playground/api-sec/pkg/env"
	"gopher-playground/api-sec/pkg/http/rest/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	godotenv.Load()

	// Load config
	cfgPath := os.Getenv(env.SERVER_CONFIG_PATH)
	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatal("Failed to load server config: ", err.Error())
	}

	// Start server
	if err = server.Start(cfg); err != nil {
		log.Fatal("Failed to start server: ", err.Error())
	}
}
