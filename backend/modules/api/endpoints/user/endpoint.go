package user

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/user"
	"backend/x/identity"
	"backend/x/web"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndpointServerPublicGetUser(dbConn *gorm.DB, c echo.Context, request *user.GetUserRequest) error {
	handle := request.GetHandle()
	if handle == "" {
		c.Logger().Error("Public GetUser Server requests must define handle")
		return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_NOT_FOUND)
	}
	return GetUser(dbConn, c, handle, false)
}

func EndpointServerPrivateGetUser(dbConn *gorm.DB, c echo.Context, currentUser *models.User, request *user.GetUserRequest) error {
	handle := request.GetHandle()
	if handle != "" {
		handle = currentUser.Handle
	}
	return GetUser(dbConn, c, currentUser.Handle, true)
}

func EndpointClientPrivateUpdateDiscordUser(dbConn *gorm.DB, c echo.Context, identity *identity.Identity, currentUser *models.User, request *user.UpdateDiscordUserRequest) error {
	var discordUser models.DiscordUser
	if err := dbConn.Transaction(func(tx *gorm.DB) error {
		if err := discordUser.FromUser(tx, identity, currentUser); err != nil {
			return fmt.Errorf("FromUser: %w", err)
		}
		if err := currentUser.FromDiscordUser(tx, &discordUser); err != nil {
			return fmt.Errorf("FromDiscordUser: %w", err)
		}
		return nil
	}); err != nil {
		c.Logger().Error("Error in update Discord user TX: ", err)
		return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_GENERAL)
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_UpdateDiscordUserResponse{
			UpdateDiscordUserResponse: &user.UpdateDiscordUserResponse{},
		},
	})
}
