package hooks

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/messages"
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo/v4"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type DiscordHooks struct{}

func (discordHooks DiscordHooks) loadUser(dbConn *gorm.DB, identity *identity.Identity, token *models.OAuth2Token) (*models.User, error) {
	var user models.User
	var discordUser models.DiscordUser
	if err := discordUser.FromOAuth2Token(dbConn, identity, token); err != nil {
		return nil, fmt.Errorf("discord from api: %w", err)
	}
	if err := user.FromDiscordUser(
		dbConn, &discordUser); err != nil {
		return nil, fmt.Errorf("from discord user: %w", err)
	}
	return &user, nil
}

func (discordHooks DiscordHooks) Name() string {
	return models.ServiceDiscord
}

func (discordHooks DiscordHooks) OnLoginRequest(dbConn *gorm.DB, c echo.Context) error {
	return nil
}

func (discordHooks DiscordHooks) GetAuthCodeURL(state string) string {
	return config.Globals.Auth.OAuth2.Config.Discord.ToOAuth2Config(discord.Endpoint).AuthCodeURL(state)
}

func (discordHooks DiscordHooks) OAuth2Exchange(authCode string) (*oauth2.Token, error) {
	return config.Globals.Auth.OAuth2.Config.Discord.ToOAuth2Config(discord.Endpoint).Exchange(context.Background(), authCode)
}

func (discordHooks DiscordHooks) OnOAuth2Callback(dbConn *gorm.DB, c echo.Context, identity *identity.Identity, token *models.OAuth2Token) error {
	user, err := discordHooks.loadUser(dbConn, identity, token)
	if err != nil {
		return fmt.Errorf("load user: %w", err)
	}
	session, sessionToken, err := user.GetSession(dbConn, identity)
	sessionMessage := messages.UserSessionCookieValue{
		Handle:   user.Handle,
		Token:    sessionToken,
		PublicId: session.PublicID,
	}
	sessionBytes, err := proto.Marshal(&sessionMessage)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := config.Globals.Auth.SessionCookies.Set(c, identity, sessionBytes); err != nil {
		return fmt.Errorf("set: %w", err)
	}
	return nil
}
