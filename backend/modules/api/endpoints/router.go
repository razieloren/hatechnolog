package endpoints

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/stats"
	"backend/modules/api/endpoints/user"
	"backend/x/identity"
	"backend/x/web"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const (
	sessionUserKey    = "session_user"
	messageRequestKey = "message_request"
)

type Router struct {
	dbConn         *gorm.DB
	identity       *identity.Identity
	sessionCookies *web.SessionCookiesConfig
}

func NewRouter(dbConn *gorm.DB, identity *identity.Identity,
	sessionCookies *web.SessionCookiesConfig) *Router {
	return &Router{
		dbConn:         dbConn,
		identity:       identity,
		sessionCookies: sessionCookies,
	}
}

func (router *Router) RPCCommunicationWrapper(apiToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Reading wrapped request.
			data, err := io.ReadAll(c.Request().Body)
			if err != nil {
				c.Logger().Error("Reading request data failed: ", err)
				return web.GenerateInternalServerError()
			}
			wrappedRequest := &messages.Wrapper{}
			if err := proto.Unmarshal(data, wrappedRequest); err != nil {
				c.Logger().Warn("Unmarshaling request message failed: ", err)
				return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
			}

			// Validating API token.
			if wrappedRequest.ApiToken == nil || wrappedRequest.GetApiToken() != apiToken {
				c.Logger().Error("API token mismatch in client request")
				return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
			}
			c.Set(messageRequestKey, wrappedRequest)
			return next(c)
		}
	}
}

func (router *Router) SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var session models.Session
		user, err := session.FromEcho(router.dbConn, router.identity, router.sessionCookies, c)
		if err != nil {
			c.Logger().Error("Failed validating session: ", err)
			return web.GenerateError(c, http.StatusUnauthorized, messages.ErrorCode_GENERAL)
		}
		c.Set(sessionUserKey, user)
		return next(c)
	}
}

func (router *Router) PublicRPCClient(c echo.Context) error {
	c.Logger().Error("Unknown public RPC client message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PublicRPCServer(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	switch request := request.Message.(type) {
	case *messages.Wrapper_LatestStatsRequest:
		return stats.EndpointLatestStats(router.dbConn, c, request.LatestStatsRequest)
	}
	c.Logger().Error("Unknown public RPC server message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PrivateRPCClient(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	sessionUser := c.Get(sessionUserKey).(*models.User)
	switch request := request.Message.(type) {
	case *messages.Wrapper_GetUserRequest:
		return user.EndpointClientGetUser(router.dbConn, c, sessionUser, request.GetUserRequest)
	}
	c.Logger().Error("Unknown private RPC client message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PrivateRPCServer(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	sessionUser := c.Get(sessionUserKey).(*models.User)
	switch request := request.Message.(type) {
	case *messages.Wrapper_GetUserRequest:
		return user.EndpointServerGetUser(router.dbConn, c, sessionUser, request.GetUserRequest)
	}
	c.Logger().Error("Unknown private RPC server message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}
