package db

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func CreateDBConnection(logger *zap.Logger, config *Config) (*gorm.DB, error) {
	dbLogger := zapgorm2.New(logger)
	dbLogger.SlowThreshold = time.Second
	dbLogger.IgnoreRecordNotFoundError = true
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%d",
		config.User, config.Password, config.Database, config.Sslmode, config.Host, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
