package main

import (
	"backend/modules/query/services"
	"backend/modules/query/services/discord"
	"backend/modules/query/services/github"
	"backend/modules/query/services/youtube"
	"backend/x/cnf"
	"backend/x/db"
	"backend/x/entrypoint"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Main struct {
		QueryIntervalSec int       `yaml:"queryIntervalSec"`
		Database         db.Config `yaml:"database"`
	} `yaml:"main"`
	Discord discord.Config `yaml:"discord"`
	Youtube youtube.Config `yaml:"youtube"`
	Github  github.Config  `yaml:"github"`
}

func main() {
	var config Config
	configPath, logger := entrypoint.Entrypoint(&config)
	servicesDef := []*services.Service{
		services.NewService("Discord", &config.Discord, discord.Work, discord.AutoMigrate),
		services.NewService("Youtube", &config.Youtube, youtube.Work, youtube.AutoMigrate),
		services.NewService("Github", &config.Github, github.Work, github.AutoMigrate),
	}

	dbConn, err := db.CreateDBConnection(logger, &config.Main.Database)
	if err != nil {
		logger.Fatal("Could not create connection to DB", zap.Error(err))
	}
	if err := services.AutoMigrate(dbConn, servicesDef); err != nil {
		logger.Fatal("Error auto-migrating DB", zap.Error(err))
	}

	logger.Info("Starting workers loop")
	mainLoopStop := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Each iteration the config is re-loaded, looking for changes.
			if err := cnf.LoadConfig(configPath, &config); err != nil {
				logger.Fatal("Could not reload config", zap.Error(err))
			}
			services.ExecuteWorkers(logger, dbConn, servicesDef)
			select {
			case <-mainLoopStop:
				logger.Info("Stopping gracefully")
				return
			case <-time.After(time.Second * time.Duration(config.Main.QueryIntervalSec)):
				continue
			}

		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	mainLoopStop <- true
	wg.Wait()
	logger.Info("Bye!")
}
