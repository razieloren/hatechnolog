package user

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/user"
	"backend/x/web"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func GetUser(dbConn *gorm.DB, c echo.Context, handle string, withPrivate bool) error {
	var requestedUser models.User
	var response user.GetUserResponse
	if err := requestedUser.FromHandle(dbConn, handle); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Logger().Error("No such handle: ", handle)
			return web.GenerateError(c, http.StatusBadRequest, messages.ErrorCode_NOT_FOUND)
		}
		c.Logger().Error("User from handle failed: ", err)
		return web.GenerateInternalServerError()
	}
	// Public info accessible to anyone.
	response.Handle = requestedUser.Handle
	response.Karma = requestedUser.Karma
	response.Plan = requestedUser.Plan.Name
	response.AvatarUrl = requestedUser.DiscordUser.GetAvatar()
	response.PlanSince = &timestamppb.Timestamp{
		Seconds: requestedUser.PlanGrantedAt.Unix(),
	}
	response.MemberSince = &timestamppb.Timestamp{
		Seconds: requestedUser.CreatedAt.Unix(),
	}
	response.IsSupporter = requestedUser.DiscordUser.IsSupporter
	response.IsVip = requestedUser.DiscordUser.IsVIP
	if requestedUser.GithubUserID != nil {
		response.GithubUsername = &requestedUser.GithubUser.Username
	}

	if withPrivate {
		if requestedUser.TACAcceptedAt != nil {
			response.TacAcceptanceDate.Seconds = requestedUser.TACAcceptedAt.Unix()
		}
		response.AllowMarketingEmails = &requestedUser.AllowMarketingEmails
		response.AllowWeeklyDigestEmails = &requestedUser.AllowWeeklyDigest
		response.MfaEnabled = &requestedUser.DiscordUser.MfaEnabled
		response.EmailVerified = &requestedUser.DiscordUser.EmailVerified
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetUserResponse{
			GetUserResponse: &response,
		},
	})
}
