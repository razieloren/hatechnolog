package hooks

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/web"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type GithubHooks struct {
	SessionCookieConf *web.CookieConfig
}

func (githubHooks GithubHooks) Name() string {
	return models.ServiceGithub
}

func (githubHooks GithubHooks) OnLoginRequest(dbConn *gorm.DB, c echo.Context) error {
	return nil
}

func (githubHooks GithubHooks) OnOAuth2Callback(dbConn *gorm.DB, c echo.Context, identity *identity.Identity, config *oauth2.Config, token *models.OAuth2Token) error {
	return nil
}
