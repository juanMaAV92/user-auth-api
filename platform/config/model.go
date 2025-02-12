package config

import "time"

type Config struct {
	MicroserviceName string
	Env              string
	HTTP             HTTPConfig
}

type HTTPConfig struct {
	Port          string
	GracefullTime time.Duration
}
