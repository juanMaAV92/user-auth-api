package database

import (
	"gorm.io/gorm/logger"
	"time"
)

type DBConfig struct {
	Host        string
	LogLevel    logger.LogLevel
	MaxLifeTime time.Duration
	MaxPoolSize int
	Name        string
	Password    string
	Port        string
	User        string
}
