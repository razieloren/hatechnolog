package hooks

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/web"
	"context"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type GithubHooks struct {
	SessionCookies *web.SessionCookiesConfig
}

func (githubHooks GithubHooks) Name() string {
	return models.ServiceGithub
}

func (githubHooks GithubHooks) OnLoginRequest(dbConn *gorm.DB, c echo.Context) error {
	return nil
}

func (githubHooks GithubHooks) GetAuthCodeURL(state string) string {
	return config.Globals.Auth.OAuth2.Config.Github.ToOAuth2Config(config.GithubOAuth2Endpoint).AuthCodeURL(state)
}

func (githubHooks GithubHooks) OAuth2Exchange(authCode string) (*oauth2.Token, error) {
	return config.Globals.Auth.OAuth2.Config.Github.ToOAuth2Config(config.GithubOAuth2Endpoint).Exchange(context.Background(), authCode)
}

func (githubHooks GithubHooks) OnOAuth2Callback(dbConn *gorm.DB, c echo.Context, identity *identity.Identity, token *models.OAuth2Token) error {
	return nil
}
