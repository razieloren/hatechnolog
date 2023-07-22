package auth

import (
	"backend/modules/api/endpoints/auth/hooks"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/messages"
	"backend/x/web"
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type Oauth2ProviderConf struct {
	RedirectUrl  string   `yaml:"redirect_url"`
	ClientId     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
}

type OAuth2Provider struct {
	dbConn          *gorm.DB
	oauth2Conf      *oauth2.Config
	stateCookieConf *web.CookieConfig
	sessionCookies  *web.SessionCookiesConfig
	identity        *identity.Identity
	serviceHooks    hooks.ServiceHooks
	redirectHost    string
}

func InstallOauth2Provider(group *echo.Group,
	dbConn *gorm.DB, identity *identity.Identity,
	oauth2Conf *oauth2.Config, serviceHooks hooks.ServiceHooks,
	stateCookieConf *web.CookieConfig, sessionCookies *web.SessionCookiesConfig,
	redirectHost string) {
	provider := &OAuth2Provider{
		dbConn:          dbConn,
		identity:        identity,
		oauth2Conf:      oauth2Conf,
		stateCookieConf: stateCookieConf,
		sessionCookies:  sessionCookies,
		redirectHost:    redirectHost,
		serviceHooks:    serviceHooks,
	}
	group.GET(fmt.Sprintf("/login/%s", provider.serviceHooks.Name()), provider.Login)
	group.GET(fmt.Sprintf("/callback/%s", provider.serviceHooks.Name()), provider.Callback)
}

func (provider *OAuth2Provider) generateStateCookie(redirectLocation string) (*http.Cookie, error) {
	state := make([]byte, 16)
	if _, err := rand.Read(state); err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	stateMessage := &messages.LoginStateCookieValue{
		State:      state,
		RedirectTo: redirectLocation,
	}
	content, err := proto.Marshal(stateMessage)
	if err != nil {
		return nil, fmt.Errorf("marhal: %w", err)
	}
	cookie, err := web.CreateEncryptedCookie(provider.stateCookieConf, provider.identity, content)
	if err != nil {
		return nil, fmt.Errorf("create encrypted cookie: %w", err)
	}
	return cookie, nil
}

func (provider *OAuth2Provider) Login(c echo.Context) error {
	redirectLocation := provider.redirectHost + c.QueryParam("return")
	var session models.Session
	_, err := session.FromEcho(provider.dbConn, provider.identity, provider.sessionCookies, c)
	if err == nil {
		// Session cookie is valid.
		return c.Redirect(http.StatusTemporaryRedirect, redirectLocation)
	}
	// Other cases, FromEcho will delete the current session cookie.
	if err := provider.serviceHooks.OnLoginRequest(provider.dbConn, c); err != nil {
		c.Logger().Error("OnLoginRequest failed: ", err)
		return web.GenerateInternalServerError()
	}
	stateCookie, err := provider.generateStateCookie(redirectLocation)
	if err != nil {
		c.Logger().Error("Generating stste cookie failed: ", err)
		return web.GenerateInternalServerError()
	}
	c.SetCookie(stateCookie)
	authCodeUrl := provider.oauth2Conf.AuthCodeURL(stateCookie.Value)
	return c.Redirect(http.StatusTemporaryRedirect, authCodeUrl)
}

func (provider *OAuth2Provider) Callback(c echo.Context) error {
	stateFromCookie, err := c.Cookie(provider.stateCookieConf.Name)
	if err != nil {
		c.Logger().Error("State cookie missing: ", err)
		return web.GenerateUnauthorizedError()
	}
	stateCookieBytes, err := web.ParseEncryptedCookieValue(provider.identity, stateFromCookie.Value)
	if err != nil {
		c.Logger().Error("Malformed state cookie: ", err)
		return web.GenerateUnauthorizedError()
	}
	var stateCookieMessage messages.LoginStateCookieValue
	if err := proto.Unmarshal(stateCookieBytes, &stateCookieMessage); err != nil {
		c.Logger().Error("State cookie umarshal failed: ", err)
		return web.GenerateUnauthorizedError()
	}
	stateFromUrl := c.QueryParam("state")
	stateUrlBytes, err := web.ParseEncryptedCookieValue(provider.identity, stateFromUrl)
	if err != nil {
		c.Logger().Error("Parse encrypted cookie failed: ", err)
		return web.GenerateUnauthorizedError()
	}
	var stateUrlMessage messages.LoginStateCookieValue
	if err := proto.Unmarshal(stateUrlBytes, &stateUrlMessage); err != nil {
		c.Logger().Error("State url umarshal failed: ", err)
		return web.GenerateUnauthorizedError()
	}
	if !bytes.Equal(stateCookieMessage.State, stateUrlMessage.State) {
		c.Logger().Error("Mismatching OAuth state")
		return web.GenerateUnauthorizedError()
	}
	authCode := c.QueryParam("code")
	token, err := provider.oauth2Conf.Exchange(context.Background(), authCode)
	if err != nil {
		c.Logger().Error("OAuth2 exchange failed: ", err)
		return web.GenerateInternalServerError()
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
			return fmt.Errorf("oauth2 create: %w", err)
		}
		if err := provider.serviceHooks.OnOAuth2Callback(tx, c, provider.identity, provider.oauth2Conf, &oauthToken); err != nil {
			return fmt.Errorf("OnOAuth2Callback: %w", err)
		}
		return nil
	}); err != nil {
		c.Logger().Error("OnOAuth2Callback tx failed: ", err)
		return web.GenerateInternalServerError()
	}
	return c.Redirect(http.StatusTemporaryRedirect, stateUrlMessage.RedirectTo)
}
