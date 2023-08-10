package endpoints

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/content"
	"backend/modules/api/endpoints/courses"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/sitemap"
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
	dbConn   *gorm.DB
	identity *identity.Identity
}

func NewRouter(dbConn *gorm.DB, identity *identity.Identity) *Router {
	return &Router{
		dbConn:   dbConn,
		identity: identity,
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
		user, err := session.FromEcho(router.dbConn, router.identity, c)
		if err != nil {
			c.Logger().Error("Failed validating session: ", err)
			return web.GenerateError(c, http.StatusUnauthorized, messages.ErrorCode_GENERAL)
		}
		c.Set(sessionUserKey, user)
		return next(c)
	}
}

func (router *Router) Sitemap(c echo.Context) error {
	return sitemap.EndpointSitemap(router.dbConn, c)
}

func (router *Router) PublicRPCClient(c echo.Context) error {
	c.Logger().Error("Unknown public RPC client message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PublicRPCServer(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	switch request := request.Message.(type) {
	case *messages.Wrapper_GetLatestStatsRequest:
		return stats.EndpointLatestStats(router.dbConn, c, request.GetLatestStatsRequest)
	case *messages.Wrapper_GetUserRequest:
		return user.EndpointServerPublicGetUser(router.dbConn, c, request.GetUserRequest)
	case *messages.Wrapper_GetCoursesTeasersRequest:
		return courses.EndpointGetCoursesTeasers(router.dbConn, c, request.GetCoursesTeasersRequest)
	case *messages.Wrapper_GetCourseRequest:
		return courses.EndpointGetCourse(router.dbConn, c, request.GetCourseRequest)
	case *messages.Wrapper_GetPageRequest:
		return content.EndpointGetPageRequest(router.dbConn, c, request.GetPageRequest)
	case *messages.Wrapper_GetPostsTeasersRequest:
		return content.EndpointGetPostsTeasersRequest(router.dbConn, c, request.GetPostsTeasersRequest)
	case *messages.Wrapper_GetPostRequest:
		return content.EndpointGetPostRequest(router.dbConn, c, request.GetPostRequest)
	case *messages.Wrapper_GetCategoriesTeasersRequest:
		return content.EndpointGetCategoriesTeasersRequest(router.dbConn, c, request.GetCategoriesTeasersRequest)
	case *messages.Wrapper_GetCategoryRequest:
		return content.EndpointGetCategoryRequest(router.dbConn, c, request.GetCategoryRequest)
	}
	c.Logger().Error("Unknown public RPC server message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PrivateRPCClient(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	sessionUser := c.Get(sessionUserKey).(*models.User)
	switch request := request.Message.(type) {
	case *messages.Wrapper_UpdateDiscordUserRequest:
		return user.EndpointClientPrivateUpdateDiscordUser(router.dbConn, c, router.identity, sessionUser, request.UpdateDiscordUserRequest)
	}
	c.Logger().Error("Unknown private RPC client message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}

func (router *Router) PrivateRPCServer(c echo.Context) error {
	request := c.Get(messageRequestKey).(*messages.Wrapper)
	sessionUser := c.Get(sessionUserKey).(*models.User)
	switch request := request.Message.(type) {
	case *messages.Wrapper_GetUserRequest:
		return user.EndpointServerPrivateGetUser(router.dbConn, c, sessionUser, request.GetUserRequest)
	}
	c.Logger().Error("Unknown private RPC server message type")
	return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
}
