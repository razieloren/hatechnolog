package user

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/user"
	"backend/x/web"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndpointClientGetUser(dbConn *gorm.DB, c echo.Context, currentUser *models.User, request *user.GetUserRequest) error {
	if request.GetHandle() != "" {
		// Client cannot request for a specific user.
		return web.GenerateError(c, http.StatusUnauthorized, messages.ErrorCode_GENERAL)
	}
	return GetUser(dbConn, c, currentUser.Handle, true)
}

func EndpointServerGetUser(dbConn *gorm.DB, c echo.Context, currentUser *models.User, request *user.GetUserRequest) error {
	handle := request.GetHandle()
	if handle == "" {
		handle = currentUser.Handle
	}
	return GetUser(dbConn, c, handle, handle == currentUser.Handle)
}
