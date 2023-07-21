package hooks

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/messages"
	"backend/x/web"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type DiscordHookConfig struct {
	GuildID         string `yaml:"guild_id"`
	SupporterRoleID string `yaml:"supporter_role_id"`
	VIPRoleID       string `yaml:"vip_role_id"`
}

type DiscordHooks struct {
	SessionCookieConf *web.CookieConfig
	Config            *DiscordHookConfig
}

func (discordHooks DiscordHooks) loadUser(dbConn *gorm.DB, config *oauth2.Config, token *models.OAuth2Token) (*models.User, error) {
	var user models.User
	var discordUser models.DiscordUser
	if err := discordUser.FromAPI(
		dbConn, config, token, discordHooks.Config.GuildID,
		discordHooks.Config.SupporterRoleID, discordHooks.Config.VIPRoleID); err != nil {
		return nil, fmt.Errorf("load discord from api: %w", err)
	}
	if err := user.FromDiscordUser(
		dbConn, &discordUser); err != nil {
		return nil, fmt.Errorf("fetch or create user discord: %w", err)
	}
	return &user, nil
}

func (discordHooks DiscordHooks) Name() string {
	return models.ServiceDiscord
}

func (discordHooks DiscordHooks) OnLoginRequest(dbConn *gorm.DB, c echo.Context) error {
	return nil
}

func (discordHooks DiscordHooks) OnOAuth2Callback(dbConn *gorm.DB, c echo.Context, identity *identity.Identity, config *oauth2.Config, token *models.OAuth2Token) error {
	user, err := discordHooks.loadUser(dbConn, config, token)
	if err != nil {
		return fmt.Errorf("fetch or create user: %w", err)
	}
	session, err := user.GetSession(dbConn)
	sessionMessage := messages.UserSessionCookieValue{
		Handle: user.Handle,
		Token:  session.Token,
	}
	sessionBytes, err := proto.Marshal(&sessionMessage)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	sessionCookie, err := web.CreateEncryptedCookie(discordHooks.SessionCookieConf, identity, sessionBytes)
	if err != nil {
		return fmt.Errorf("create cookie: %w", err)
	}
	c.SetCookie(sessionCookie)
	return nil
}
