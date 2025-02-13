package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	database2 "github.com/juanMaAV92/user-auth-api/platform/database"
	"github.com/juanMaAV92/user-auth-api/utils/path"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	connectionRetries = 5
)

var Logger = gormLogger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	gormLogger.Config{
		SlowThreshold:             200 * time.Millisecond, //nolint
		LogLevel:                  gormLogger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	},
)

var singleton *database2.Database

func New(cfg DBConfig) (*database2.Database, error) {
	if singleton != nil {
		singleton = &database2.Database{}
	}

	gdb, err := connect(cfg)
	if err != nil {
		return nil, err
	}
	singleton = &database2.Database{DB: gdb}
	return singleton, nil
}

func connect(cfg DBConfig) (*gorm.DB, error) {
	dsn := _generateDSN(cfg)
	instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: Logger.LogMode(gormLogger.Error),
	})

	for connectionTries := 0; err != nil && connectionTries <= connectionRetries; connectionTries++ {
		instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: Logger.LogMode(gormLogger.Error),
		})
		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("an error occurred while open gorm: %w", err)
	}
	db, err := instance.DB()
	if err != nil {
		return nil, fmt.Errorf("an error occurred getting db instance: %w", err)
	}
	db.SetMaxIdleConns(cfg.MaxPoolSize)
	db.SetMaxOpenConns(cfg.MaxPoolSize)
	db.SetConnMaxLifetime(cfg.MaxLifeTime)

	if err = applyMigrations(cfg); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("an error occurred applying migration: %w", err)
	}

	return instance, nil
}

func applyMigrations(cfg DBConfig) error {
	currentPath := fmt.Sprintf("file:///%s/migration", path.GetMainPath())
	url := generatePgURL(cfg)

	m, err := migrate.New(currentPath, url)
	if err != nil {
		return err
	}

	return m.Up()
}

func generatePgURL(cfg DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}

func _generateDSN(cfg DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
}
