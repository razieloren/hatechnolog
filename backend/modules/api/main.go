package main

import (
	"backend/db"
	"backend/entrypoint"
	"backend/modules/api/endpoints"
	"fmt"
	"net"
	"net/http"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Config struct {
	Main struct {
		Server struct {
			ListenAddr       *string `yaml:"listenAddr"`
			ListenPort       *int    `yaml:"listenPort"`
			ListenUnixSocket *string `yaml:"listenUnixSocket"`
		} `yaml:"server"`
		Database db.Config `yaml:"database"`
	} `yaml:"main"`
	Endpoints endpoints.Config `yaml:"endpoints"`
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

	if config.Main.Server.ListenUnixSocket != nil {
		logger.Info("Starting echo server using a UNIX socket", zap.String("path", *config.Main.Server.ListenUnixSocket))
		listener, err := net.Listen("unix", *config.Main.Server.ListenUnixSocket)
		if err != nil {
			logger.Fatal("Error listening on UNIX socket", zap.String("path", *config.Main.Server.ListenUnixSocket), zap.Error(err))
		}
		e.Listener = listener
		server := new(http.Server)
		e.Logger.Fatal(e.StartServer(server))
	} else if config.Main.Server.ListenAddr != nil && config.Main.Server.ListenPort != nil {
		logger.Info("Starting echo server using normal socket", zap.String("addr", *config.Main.Server.ListenAddr), zap.Int("port", *config.Main.Server.ListenPort))
		e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", *config.Main.Server.ListenAddr, *config.Main.Server.ListenPort)))
	}
	logger.Fatal("Either a UNIX socket or a normal socket must be defined in the config.yaml file")
}
