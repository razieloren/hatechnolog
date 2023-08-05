package hooks

import (
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type ServiceHooks interface {
	Name() string
	OnLoginRequest(*gorm.DB, echo.Context) error
	GetAuthCodeURL(state string) string
	OAuth2Exchange(authCode string) (*oauth2.Token, error)
	OnOAuth2Callback(*gorm.DB, echo.Context, *identity.Identity, *models.OAuth2Token) error
}
