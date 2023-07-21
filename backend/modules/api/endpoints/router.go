package endpoints

import (
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/stats"
	"backend/x/web"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const (
	requestMessageKey = "rpc_message"
)

type Router struct {
	logger         *zap.Logger
	dbConn         *gorm.DB
	serverApiToken string
	clientApiToken string
}

func NewRouter(logger *zap.Logger, dbConn *gorm.DB, serverApiToken, clientApiToken string) *Router {
	return &Router{
		logger:         logger,
		dbConn:         dbConn,
		serverApiToken: serverApiToken,
		clientApiToken: clientApiToken,
	}
}

func (router *Router) ExtractRPCMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := io.ReadAll(c.Request().Body)
		if err != nil {
			router.logger.Error("Could not read request data", zap.Error(err))
			return web.StatusBadRequest
		}
		wrappedRequest := &messages.Wrapper{}
		if err := proto.Unmarshal(data, wrappedRequest); err != nil {
			router.logger.Error("Could not unmarshal wrapped message", zap.Error(err))
			return web.StatusBadRequest
		}
		c.Set(requestMessageKey, wrappedRequest)
		return next(c)
	}
}

func (router *Router) RPCClient(c echo.Context) error {
	wrappedRequest, ok := c.Get(requestMessageKey).(*messages.Wrapper)
	if !ok {
		router.logger.Error("Cannot convert message to wrapper")
		return web.StatusBadRequest
	}
	if wrappedRequest.ApiToken != router.clientApiToken {
		router.logger.Error("Bad client API token")
		return web.StatusUnauthorized
	}
	return c.String(http.StatusAccepted, "empty")
}

func (router *Router) RPCServer(c echo.Context) error {
	wrappedRequest, ok := c.Get(requestMessageKey).(*messages.Wrapper)
	if !ok {
		router.logger.Error("Cannot convert message to wrapper")
		return web.StatusBadRequest
	}
	if wrappedRequest.ApiToken != router.serverApiToken {
		router.logger.Error("Bad server API token")
		return web.StatusBadRequest
	}

	wrappedResponse := &messages.Wrapper{}
	switch request := wrappedRequest.Message.(type) {
	case *messages.Wrapper_LatestStatsRequest:
		response := stats.EndpointLatestStats(router.dbConn, router.logger, request.LatestStatsRequest)
		wrappedResponse.Message = &messages.Wrapper_LatestStatsResponse{
			LatestStatsResponse: response,
		}
	default:
		router.logger.Error("Unknown server message type")
		return web.StatusBadRequest
	}
	responseData, err := proto.Marshal(wrappedResponse)
	if err != nil {
		router.logger.Error("Could not marshal message", zap.Error(err))
		return web.StatusInternalServerError
	}
	return c.Blob(http.StatusOK, "application/octet-stream", responseData)
}
