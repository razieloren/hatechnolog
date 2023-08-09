package user

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/modules/api/endpoints/content"
	content_models "backend/modules/api/endpoints/content/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/user"
	"backend/x/web"
	"errors"
	"fmt"
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
	response.IsVip = requestedUser.DiscordUser.IsVIP
	if requestedUser.GithubUserID != nil {
		response.GithubUsername = &requestedUser.GithubUser.Username
	}

	var contents []content_models.ContentTeaser
	if err := content.QueryPosts(dbConn).Where(&content_models.Content{
		UserID: requestedUser.ID,
	}).Find(&contents).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.Logger().Error("Contents from user: ", err)
			return web.GenerateInternalServerError()
		}
	}
	fmt.Println(requestedUser.ID, len(contents))
	for _, item := range contents {
		response.Contents = append(response.Contents, content.ContentToTeaser(&item))
	}

	if withPrivate {
		response.IsHatechnologMember = &requestedUser.DiscordUser.HatechnologMember
		response.MfaEnabled = &requestedUser.DiscordUser.MfaEnabled
		response.EmailVerified = &requestedUser.DiscordUser.EmailVerified
		response.State = user.UserState(requestedUser.State)
		response.Me = true
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetUserResponse{
			GetUserResponse: &response,
		},
	})
}
