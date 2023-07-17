package endpoints

import (
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/stats"
	"backend/x/ws"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Stats struct {
		LatestStatsPushIntervalSec int `yaml:"latest_stats_push_interval_sec"`
	} `yaml:"stats"`
}

type Router struct {
	logger *zap.Logger
	dbConn *gorm.DB
	config *Config
}

func NewRouter(logger *zap.Logger, dbConn *gorm.DB, config *Config) *Router {
	return &Router{
		logger: logger,
		dbConn: dbConn,
		config: config,
	}
}

func (router *Router) HandleWS(c echo.Context) error {
	ws, err := ws.NewWS(c)
	if err != nil {
		router.logger.Error("Error upgrading WS", zap.Error(err))
		return err
	}
	defer ws.Close()
	wrapped, err := ws.ReadWrappedMessage()
	if err != nil {
		return fmt.Errorf("Error reading initial message: %w", err)
	}
	endpointIntentWrapped, ok := wrapped.Message.(*messages.Wrapper_EndpointIntent)
	if !ok {
		return fmt.Errorf("Bad message type received")
	}
	intent := endpointIntentWrapped.EndpointIntent
	switch i := intent.Endpoint; i {
	case messages.Endpoint_latest_stats_push:
		err := stats.EndpointLatestStatsPush(
			router.logger, router.dbConn, ws,
			time.Second*time.Duration(router.config.Stats.LatestStatsPushIntervalSec))
		if err != nil {
			router.logger.Error("Error in latest_stats_push endpoint", zap.Error(err))
			return err
		}
	default:
		router.logger.Error("Unknown endpoint intent", zap.Any("intent", i))
		return fmt.Errorf("Unknown endpoint intent")
	}
	return nil
}
