package config

import (
	"github.com/juanMaAV92/user-auth-api/utils/database"
	"time"
)

type Config struct {
	MicroserviceName string
	Env              string
	HTTP             HTTPConfig
	Database         database.DBConfig
}

type HTTPConfig struct {
	Port          string
	GracefullTime time.Duration
}
