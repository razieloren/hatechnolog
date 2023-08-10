package main

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints"
	"backend/modules/api/endpoints/auth"
	"backend/modules/api/endpoints/auth/hooks"
	"backend/x/db"
	"backend/x/entrypoint"
	"backend/x/identity"
	"fmt"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	_, logger := entrypoint.Entrypoint(&config.Globals)

	dbConn, err := db.CreateDBConnection(logger, &config.Globals.Database)
	if err != nil {
		logger.Fatal("Could not create connection to DB", zap.Error(err))
	}

	serverIdentity := identity.NewIdentity(config.Globals.Identity.Secret, config.Globals.Identity.Salt)

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(echozap.ZapLogger(logger))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.Globals.Server.AllowedOrigin},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.POST, echo.OPTIONS},
		AllowCredentials: true,
	}))

	// Registering Router endpoints.
	router := endpoints.NewRouter(dbConn, serverIdentity)
	e.GET("/sitemap.xml", router.Sitemap)
	rpcGroup := e.Group("/rpc")
	rpcGroup.POST("/server",
		router.PublicRPCServer,
		router.RPCCommunicationWrapper(config.Globals.Api.ServerToken))
	rpcGroup.POST("/client",
		router.PublicRPCServer,
		router.RPCCommunicationWrapper(config.Globals.Api.ClientToken))
	rpcAuthGroup := rpcGroup.Group("/private", router.SessionMiddleware)
	rpcAuthGroup.POST("/server",
		router.PrivateRPCServer,
		router.RPCCommunicationWrapper(config.Globals.Api.ServerToken))
	rpcAuthGroup.POST("/client",
		router.PrivateRPCClient,
		router.RPCCommunicationWrapper(config.Globals.Api.ClientToken))

	// Registering auth endpoints.
	authGroup := e.Group("/auth")
	auth.InstallLogoutHandler(authGroup, dbConn, serverIdentity)
	auth.InstallOauth2Provider(authGroup, dbConn, serverIdentity, hooks.DiscordHooks{})
	auth.InstallOauth2Provider(authGroup, dbConn, serverIdentity, hooks.GithubHooks{})

	if config.Globals.TLSConfigured() {
		logger.Info("Starting HTTPS server",
			zap.String("addr", config.Globals.Server.ListenAddr),
			zap.Int("port", config.Globals.Server.ListenPort))
		e.Logger.Fatal(
			e.StartTLS(
				fmt.Sprintf("%s:%d", config.Globals.Server.ListenAddr,
					config.Globals.Server.ListenPort),
				*config.Globals.Server.FullchainPath,
				*config.Globals.Server.KeyPath))
	}

	logger.Info("Starting HTTP server",
		zap.String("addr", config.Globals.Server.ListenAddr),
		zap.Int("port", config.Globals.Server.ListenPort))
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf("%s:%d", config.Globals.Server.ListenAddr,
				config.Globals.Server.ListenPort)))
}
