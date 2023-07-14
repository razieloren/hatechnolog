package services

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AutoMigrateFunc func(dbConn *gorm.DB) error
type WorkerFunc func(logger *zap.Logger, dbConn *gorm.DB, config interface{})

type Service struct {
	FriendlyName    string
	Config          interface{}
	WorkerFunc      WorkerFunc
	AutoMigrateFunc AutoMigrateFunc
}

func NewService(friendlyName string, config interface{}, workerFn WorkerFunc, autoMigrateFn AutoMigrateFunc) *Service {
	return &Service{
		FriendlyName:    friendlyName,
		Config:          config,
		WorkerFunc:      workerFn,
		AutoMigrateFunc: autoMigrateFn,
	}
}

func ExecuteWorkers(logger *zap.Logger, dbConn *gorm.DB, services []*Service) {
	var wg sync.WaitGroup
	start := time.Now()
	logger.Info("Executing workers",
		zap.Time("start", start), zap.Int("amount", len(services)))
	for i := range services {
		wg.Add(1)
		go func(service *Service) {
			defer wg.Done()
			logger.Debug("Service working",
				zap.String("friendly_name", service.FriendlyName))
			serviceStart := time.Now()
			service.WorkerFunc(logger, dbConn, service.Config)
			serviceElapsed := time.Since(serviceStart)
			logger.Debug("Service finished",
				zap.String("friendly_name", service.FriendlyName),
				zap.Duration("elapsed", serviceElapsed))
		}(services[i])
	}
	wg.Wait()
	elapsed := time.Since(start)
	logger.Info("Workers finished",
		zap.Duration("elapsed", elapsed))
}

func AutoMigrate(dbConn *gorm.DB, services []*Service) error {
	for _, service := range services {
		if err := service.AutoMigrateFunc(dbConn); err != nil {
			return fmt.Errorf("%s auto_migrate_func: %w", service.FriendlyName, err)
		}
	}
	return nil
}
