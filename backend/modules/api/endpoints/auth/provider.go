package auth

import (
	"backend/modules/api/endpoints/auth/hooks"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/web"
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Oauth2ProviderConf struct {
	RedirectUrl  string   `yaml:"redirect_url"`
	ClientId     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
}

type OAuth2Provider struct {
	logger            *zap.Logger
	dbConn            *gorm.DB
	oauth2Conf        *oauth2.Config
	stateCookieConf   *web.CookieConfig
	sessionCookieConf *web.CookieConfig
	identity          *identity.Identity
	serviceHooks      hooks.ServiceHooks
	finalRedirect     string
}

func InstallOauth2Provider(group *echo.Group,
	logger *zap.Logger, dbConn *gorm.DB, identity *identity.Identity,
	oauth2Conf *oauth2.Config, serviceHooks hooks.ServiceHooks,
	stateCookieConf *web.CookieConfig, sessionCookieConf *web.CookieConfig,
	finalRedirect string) {
	provider := &OAuth2Provider{
		logger:            logger,
		dbConn:            dbConn,
		identity:          identity,
		oauth2Conf:        oauth2Conf,
		stateCookieConf:   stateCookieConf,
		sessionCookieConf: sessionCookieConf,
		finalRedirect:     finalRedirect,
		serviceHooks:      serviceHooks,
	}
	group.GET(fmt.Sprintf("/login/%s", provider.serviceHooks.Name()), provider.Login)
	group.GET(fmt.Sprintf("/callback/%s", provider.serviceHooks.Name()), provider.Callback)
}

func (provider *OAuth2Provider) generateStateCookie() (*http.Cookie, error) {
	state := make([]byte, 16)
	if _, err := rand.Read(state); err != nil {
		return nil, fmt.Errorf("error reading random: %w", err)
	}
	cookie, err := web.CreateEncryptedCookie(provider.stateCookieConf, provider.identity, state)
	if err != nil {
		return nil, fmt.Errorf("create cookie: %w", err)
	}
	return cookie, nil
}

func (provider *OAuth2Provider) Login(c echo.Context) error {
	var session models.Session
	_, err := session.FromEcho(provider.dbConn, provider.identity, provider.sessionCookieConf, c)
	if err == nil {
		// Session cookie is valid.
		return c.Redirect(http.StatusTemporaryRedirect, provider.finalRedirect)
	}
	// Other cases, FromEcho will delete the current session cookie.
	if err := provider.serviceHooks.OnLoginRequest(provider.dbConn, c); err != nil {
		provider.logger.Error("Error in OnLoginRequest service hook", zap.Error(err))
		return web.StatusBadRequest
	}
	stateCookie, err := provider.generateStateCookie()
	if err != nil {
		provider.logger.Error("Error generating state cookie", zap.Error(err))
		return web.StatusInternalServerError
	}
	c.SetCookie(stateCookie)
	authCodeUrl := provider.oauth2Conf.AuthCodeURL(stateCookie.Value)
	return c.Redirect(http.StatusTemporaryRedirect, authCodeUrl)
}

func (provider *OAuth2Provider) Callback(c echo.Context) error {
	stateFromCookie, err := c.Cookie(provider.stateCookieConf.Name)
	if err != nil {
		provider.logger.Error("No state cookie in oauth2 callback", zap.Error(err))
		return web.StatusUnauthorized
	}
	safeStateVal, err := web.ParseEncryptedCookieValue(provider.identity, stateFromCookie.Value)
	if err != nil {
		provider.logger.Error("Bad state cookie", zap.Error(err))
		return web.StatusUnauthorized
	}
	stateValUrl := c.QueryParam("state")
	safeValUrl, err := web.ParseEncryptedCookieValue(provider.identity, stateValUrl)
	if err != nil {
		provider.logger.Error("Bad state url", zap.Error(err))
		return web.StatusUnauthorized
	}
	if !bytes.Equal(safeStateVal, safeValUrl) {
		provider.logger.Error("Bad session value", zap.String("from_cookie", stateFromCookie.Value), zap.String("from_url", stateValUrl))
		return web.StatusUnauthorized
	}
	authCode := c.QueryParam("code")
	token, err := provider.oauth2Conf.Exchange(context.Background(), authCode)
	if err != nil {
		provider.logger.Error("Failed exchanging auth code for token", zap.Error(err))
		return web.StatusInternalServerError
	}
	c.SetCookie(web.DeleteCookie(stateFromCookie.Name))
	// Save the OAuth2 token to the DB to the system user.
	oauthToken := models.OAuth2Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		TokenExpiry:  token.Expiry,
	}
	if err := provider.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&oauthToken).Error; err != nil {
			return fmt.Errorf("failed creating oauth token: %w", err)
		}
		if err := provider.serviceHooks.OnOAuth2Callback(tx, c, provider.identity, provider.oauth2Conf, &oauthToken); err != nil {
			return fmt.Errorf("OnOAuth2Callback failed: %w", err)
		}
		return nil
	}); err != nil {
		provider.logger.Error("OnOAuth2Callback tx failed", zap.Error(err))
		return web.StatusInternalServerError
	}
	return c.Redirect(http.StatusTemporaryRedirect, provider.finalRedirect)
}
