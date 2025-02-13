package config

import (
	"errors"
	"github.com/juanMaAV92/user-auth-api/utils/database"
	"github.com/juanMaAV92/user-auth-api/utils/enviroment"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

const (
	MicroserviceName = "user-auth-api"
)

var (
	httpConfig = HTTPConfig{
		Port:          os.Getenv("PORT"),
		GracefullTime: time.Duration(enviroment.GetEnvAsIntWithDefault("GRACEFUL_TIME", 60)) * time.Second,
	}

	databaseConfig = database.DBConfig{
		Host:        os.Getenv("DB_HOST_POSTGRES"),
		Password:    os.Getenv("DB_PASSWORD_POSTGRES"),
		User:        os.Getenv("DB_USER_POSTGRES"),
		Port:        os.Getenv("DB_PORT_POSTGRES"),
		Name:        os.Getenv("DB_NAME_POSTGRES"),
		LogLevel:    logger.Silent,
		MaxPoolSize: 2,
		MaxLifeTime: 5 * time.Minute,
	}
)

func _deployedConfig() Config {
	return Config{
		HTTP:     httpConfig,
		Database: databaseConfig,
	}
}

var _config = map[string]Config{
	"local": {
		HTTP: HTTPConfig{
			Port:          "8080",
			GracefullTime: 5 * time.Second,
		},
		Database: database.DBConfig{
			Host:        "localhost",
			Password:    "postgres",
			User:        "postgres",
			Port:        "5432",
			Name:        "user-auth-db",
			LogLevel:    logger.Info,
			MaxPoolSize: 15,
			MaxLifeTime: 5 * time.Minute,
		},
	},
	"stg": _deployedConfig(),
}

func Load(env string) (Config, error) {
	cfg, found := _config[env]
	if !found {
		return Config{}, errors.New("config not found")
	}
	cfg.MicroserviceName = MicroserviceName
	cfg.Env = env
	return cfg, nil
}
