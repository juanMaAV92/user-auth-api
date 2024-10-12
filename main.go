package main

import (
	"github.com/juanMaAV92/user-auth-api/utils/logger"
	"log"
	"os"

	"github.com/juanMaAV92/user-auth-api/cmd"
	"github.com/juanMaAV92/user-auth-api/cmd/routes"
	"github.com/juanMaAV92/user-auth-api/config"
)

const (
	environmentEnv = "ENV"

	exitCodeFailReadConfigs = 2
)

func main() {

	env := os.Getenv(environmentEnv)
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		os.Exit(exitCodeFailReadConfigs)
	}

	server := cmd.NewServer(cfg)
	routes.RegisterRoutes(server)

	errChannel := server.Start()

	if err := <-errChannel; err != nil {
		server.Logger.Info("server", "Error starting server", logger.Field("error", err))
		os.Exit(1)
	}
}
