package auth

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InstallLogoutHandler(authGroup *echo.Group, dbConn *gorm.DB, identity *identity.Identity) {
	authGroup.GET("/logout", func(c echo.Context) error {
		var session models.Session
		_, err := session.FromEcho(dbConn, identity, c)
		if err == nil {
			config.Globals.Auth.SessionCookies.Delete(c)
			// This might return an error but we have nothing to do here.
			session.Invalidate(dbConn)
		}
		redirectTo := config.Globals.Auth.RedirectHost + c.QueryParam(config.Globals.Auth.RedirectParam)
		return c.Redirect(http.StatusTemporaryRedirect, redirectTo)
	})
}
