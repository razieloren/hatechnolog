package main

import (
	"backend/modules/api/endpoints"
	"backend/x/db"
	"backend/x/entrypoint"
	"fmt"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Config struct {
	Main struct {
		Server struct {
			ListenAddr string `yaml:"listenAddr"`
			ListenPort int    `yaml:"listenPort"`
			TLS        struct {
				KeyPath       *string `yaml:"key_path"`
				FullchainPath *string `yaml:"fullchain_path"`
			} `yaml:"tls"`
		} `yaml:"server"`
		Database db.Config `yaml:"database"`
	} `yaml:"main"`
	Endpoints endpoints.Config `yaml:"endpoints"`
}

func (config *Config) TLSConfigured() bool {
	return config.Main.Server.TLS.FullchainPath != nil && config.Main.Server.TLS.KeyPath != nil
}

func main() {
	var config Config
	_, logger := entrypoint.Entrypoint(&config)

	dbConn, err := db.CreateDBConnection(logger, &config.Main.Database)
	if err != nil {
		logger.Fatal("Could not create connection to DB", zap.Error(err))
	}

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(echozap.ZapLogger(logger))

	router := endpoints.NewRouter(logger, dbConn, &config.Endpoints)
	e.GET("/ws", router.HandleWS)

	serverConf := &config.Main.Server
	if config.TLSConfigured() {
		logger.Info("Starting HTTPS server",
			zap.String("addr", serverConf.ListenAddr),
			zap.Int("port", serverConf.ListenPort))
		e.Logger.Fatal(
			e.StartTLS(
				fmt.Sprintf("%s:%d", serverConf.ListenAddr, serverConf.ListenPort),
				*serverConf.TLS.FullchainPath,
				*serverConf.TLS.KeyPath))
	}

	logger.Info("Starting HTTP server",
		zap.String("addr", serverConf.ListenAddr),
		zap.Int("port", serverConf.ListenPort))
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf("%s:%d", serverConf.ListenAddr, serverConf.ListenPort)))
}
