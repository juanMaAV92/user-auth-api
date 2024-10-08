package config

import (
	"errors"
	"os"
)

var _config = map[string]Config{
	"prod": {
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("ENV"),
		AppName: os.Getenv("APP_NAME"),
	},
	"local": {
		Port:    "8080",
		Env:     "local",
		AppName: "user-auth-api",
	},
}

func LoadConfig(env string) (Config, error) {
	cfg, found := _config[env]
	if !found {
		return Config{}, errors.New("config not found")
	}
	return cfg, nil
}
