package main

import (
	"backend/modules/api/config"
	"backend/modules/api/initdb/models"
	"backend/x/db"
	"backend/x/entrypoint"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AutoMigrateFunc func(*gorm.DB) error
type CreateDataFunc func(*gorm.DB) error

type InitDBEntry struct {
	Name        string
	AutoMigrate AutoMigrateFunc
	CreateData  CreateDataFunc
}

func main() {
	_, logger := entrypoint.Entrypoint(&config.Globals)

	dbConn, err := db.CreateDBConnection(logger, &config.Globals.Database)
	if err != nil {
		logger.Fatal("Could not create connection to DB", zap.Error(err))
	}

	entries := []InitDBEntry{
		// {
		// 	Name:        "auth",
		// 	AutoMigrate: models.AutoMigrateAuth,
		// 	CreateData:  models.CreateDataAuth,
		// },
		// {
		// 	Name:        "courses",
		// 	AutoMigrate: models.AutoMigrateCourses,
		// 	CreateData:  models.CreateDataCourses,
		// },
		{
			Name:        "content",
			AutoMigrate: models.AutoMigrateContent,
			CreateData:  models.CreateDataContent,
		},
	}

	for _, entry := range entries {
		logger.Info("Auto migrating models", zap.String("name", entry.Name))
		if err := entry.AutoMigrate(dbConn); err != nil {
			logger.Fatal("Error auto-migrating", zap.String("name", entry.Name), zap.Error(err))
		}
		logger.Info("Populating tables with data", zap.String("name", entry.Name))
		if err := entry.CreateData(dbConn); err != nil {
			logger.Fatal("Error populating tables with data", zap.String("name", entry.Name), zap.Error(err))
		}
	}
}
