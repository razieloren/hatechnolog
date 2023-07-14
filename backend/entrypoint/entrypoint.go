package entrypoint

import (
	"backend/cnf"
	"backend/log"
	"flag"

	"go.uber.org/zap"
)

func Entrypoint(config interface{}) (string, *zap.Logger) {
	configPath := flag.String("config", "config.yaml", "Path to config file, default: './config.yaml'")
	flag.Parse()

	logger := log.CreateLogger()
	if err := cnf.LoadConfig(*configPath, config); err != nil {
		logger.Fatal("Could not load config", zap.Error(err))
	}
	logger.Info("Welcome")
	return *configPath, logger
}
