package main

import (
	"backend/modules/api/endpoints"
	"backend/modules/api/endpoints/auth"
	"backend/modules/api/endpoints/auth/hooks"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/db"
	"backend/x/entrypoint"
	"backend/x/identity"
	"backend/x/web"
	"flag"
	"fmt"
	"net/http"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/ravener/discord-oauth2"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Config struct {
	Server struct {
		ListenAddr    string  `yaml:"listenAddr"`
		ListenPort    int     `yaml:"listenPort"`
		KeyPath       *string `yaml:"key_path"`
		FullchainPath *string `yaml:"fullchain_path"`
	} `yaml:"server"`
	Identity struct {
		Salt   string `yaml:"salt"`
		Secret string `yaml:"secret"`
	} `yaml:"identity"`
	ApiToken struct {
		Client string `yaml:"client"`
		Server string `yaml:"server"`
	} `yaml:"api_token"`
	Auth struct {
		FinalRedirect string           `yaml:"final_redirect"`
		SessionCookie web.CookieConfig `yaml:"session_cookie"`
		Oauth2        struct {
			StateCookie web.CookieConfig `yaml:"state_cookie"`
			Config      struct {
				Discord auth.Oauth2ProviderConf `yaml:"discord"`
				Github  auth.Oauth2ProviderConf `yaml:"github"`
			} `yaml:"config"`
			Hooks struct {
				Discord hooks.DiscordHookConfig `yaml:"discord"`
			} `yaml:"hooks"`
		} `yaml:"oauth2"`
	} `yaml:"auth"`
	Database db.Config `yaml:"database"`
}

func (config *Config) TLSConfigured() bool {
	return config.Server.FullchainPath != nil && config.Server.KeyPath != nil
}

func main() {
	var config Config
	initDB := flag.Bool("initdb", false, "Initialize the backend API scheme, default: false")
	_, logger := entrypoint.Entrypoint(&config)

	dbConn, err := db.CreateDBConnection(logger, &config.Database)
	if err != nil {
		logger.Fatal("Could not create connection to DB", zap.Error(err))
	}
	if *initDB {
		logger.Info("Auto-migrating models")
		if err := models.AutoMigrate(dbConn); err != nil {
			logger.Fatal("Error auto-migrating DB", zap.Error(err))
		}
		logger.Info("Populating DB with default data")
		if err := models.CreateDefaultData(dbConn); err != nil {
			logger.Fatal("Error creating default data", zap.Error(err))
		}
		return
	}

	serverIdentity := identity.NewIdentity(config.Identity.Secret, config.Identity.Salt)

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(echozap.ZapLogger(logger))

	// Registering RPC endpoints.
	router := endpoints.NewRouter(logger, dbConn, config.ApiToken.Server, config.ApiToken.Client)
	rpcGroup := e.Group("/rpc", router.ExtractRPCMessage)
	rpcGroup.POST("/server", router.RPCServer)
	rpcGroup.POST("/client", router.RPCClient)

	// Registering auth endpoints.
	authGroup := e.Group("/auth")
	authGroup.POST("/logout", func(c echo.Context) error {
		// TODO: delete session here.
		var session models.Session
		_, err := session.FromEcho(dbConn, serverIdentity, &config.Auth.SessionCookie, c)
		if err == nil {
			// This is a valid session so we need to invalidate it.
			c.SetCookie(web.DeleteCookie(config.Auth.SessionCookie.Name))
			// This might return an error but we have nothing to do here.
			session.Invalidate(dbConn)
		}

		return c.Redirect(http.StatusTemporaryRedirect, config.Auth.FinalRedirect)
	})
	auth.InstallOauth2Provider(authGroup, logger, dbConn, serverIdentity,
		&oauth2.Config{
			RedirectURL:  config.Auth.Oauth2.Config.Discord.RedirectUrl,
			ClientID:     config.Auth.Oauth2.Config.Discord.ClientId,
			ClientSecret: config.Auth.Oauth2.Config.Discord.ClientSecret,
			Scopes:       config.Auth.Oauth2.Config.Discord.Scopes,
			Endpoint:     discord.Endpoint,
		}, hooks.DiscordHooks{
			SessionCookieConf: &config.Auth.SessionCookie,
			Config:            &config.Auth.Oauth2.Hooks.Discord,
		}, &config.Auth.Oauth2.StateCookie, &config.Auth.SessionCookie, config.Auth.FinalRedirect)
	auth.InstallOauth2Provider(authGroup, logger, dbConn, serverIdentity,
		&oauth2.Config{
			RedirectURL:  config.Auth.Oauth2.Config.Github.RedirectUrl,
			ClientID:     config.Auth.Oauth2.Config.Github.ClientId,
			ClientSecret: config.Auth.Oauth2.Config.Github.ClientSecret,
			Scopes:       config.Auth.Oauth2.Config.Github.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://github.com/login/oauth/authorize",
				TokenURL:  "https://github.com/login/oauth/access_token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		}, hooks.GithubHooks{
			SessionCookieConf: &config.Auth.SessionCookie,
		}, &config.Auth.Oauth2.StateCookie, &config.Auth.SessionCookie, config.Auth.FinalRedirect)

	serverConf := &config.Server
	if config.TLSConfigured() {
		logger.Info("Starting HTTPS server",
			zap.String("addr", serverConf.ListenAddr),
			zap.Int("port", serverConf.ListenPort))
		e.Logger.Fatal(
			e.StartTLS(
				fmt.Sprintf("%s:%d", serverConf.ListenAddr, serverConf.ListenPort),
				*serverConf.FullchainPath,
				*serverConf.KeyPath))
	}

	logger.Info("Starting HTTP server",
		zap.String("addr", serverConf.ListenAddr),
		zap.Int("port", serverConf.ListenPort))
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf("%s:%d", serverConf.ListenAddr, serverConf.ListenPort)))
}
