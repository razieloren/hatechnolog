package auth

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints/auth/hooks"
	"backend/modules/api/endpoints/auth/models"
	"backend/x/identity"
	"backend/x/messages"
	"backend/x/web"
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type OAuth2Provider struct {
	dbConn       *gorm.DB
	identity     *identity.Identity
	serviceHooks hooks.ServiceHooks
}

func InstallOauth2Provider(
	group *echo.Group,
	dbConn *gorm.DB, identity *identity.Identity,
	serviceHooks hooks.ServiceHooks) {
	provider := &OAuth2Provider{
		dbConn:       dbConn,
		identity:     identity,
		serviceHooks: serviceHooks,
	}
	group.GET(fmt.Sprintf("/login/%s", provider.serviceHooks.Name()), provider.Login)
	group.GET(fmt.Sprintf("/callback/%s", provider.serviceHooks.Name()), provider.Callback)
}

func (provider *OAuth2Provider) generateStateCookie(redirectTo string) (*http.Cookie, error) {
	state := make([]byte, 16)
	if _, err := rand.Read(state); err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	stateMessage := &messages.LoginStateCookieValue{
		State:      state,
		RedirectTo: redirectTo,
	}
	content, err := proto.Marshal(stateMessage)
	if err != nil {
		return nil, fmt.Errorf("marhal: %w", err)
	}
	cookie, err := web.CreateEncryptedCookie(&config.Globals.Auth.OAuth2.StateCookie, provider.identity, content)
	if err != nil {
		return nil, fmt.Errorf("create encrypted cookie: %w", err)
	}
	return cookie, nil
}

func (provider *OAuth2Provider) generateResponseRedirectUrl(redirectTo string) string {
	// Ignoring the "err" since "redirectHost" should always be valid...
	baseUrl, _ := url.Parse(config.Globals.Auth.RedirectHost)
	baseUrl.Path += "users/welcome"
	params := url.Values{}
	params.Add(config.Globals.Auth.RedirectParam, redirectTo)
	baseUrl.RawQuery = params.Encode()
	// Value: scheme://redirectHost/users/welcome?redirect=redirectTo
	return baseUrl.String()
}

func (provider *OAuth2Provider) Login(c echo.Context) error {
	var session models.Session
	_, err := session.FromEcho(provider.dbConn, provider.identity, c)
	if err == nil {
		// Session cookie is valid.
		return c.Redirect(http.StatusTemporaryRedirect, provider.generateResponseRedirectUrl(c.QueryParam(config.Globals.Auth.RedirectParam)))
	}
	// Other cases, FromEcho will delete the current session cookie.
	if err := provider.serviceHooks.OnLoginRequest(provider.dbConn, c); err != nil {
		c.Logger().Error("OnLoginRequest failed: ", err)
		return web.GenerateInternalServerError()
	}
	stateCookie, err := provider.generateStateCookie(c.QueryParam(config.Globals.Auth.RedirectParam))
	if err != nil {
		c.Logger().Error("Generating stste cookie failed: ", err)
		return web.GenerateInternalServerError()
	}
	c.SetCookie(stateCookie)
	authCodeUrl := provider.serviceHooks.GetAuthCodeURL(stateCookie.Value)
	return c.Redirect(http.StatusTemporaryRedirect, authCodeUrl)
}

func (provider *OAuth2Provider) Callback(c echo.Context) error {
	stateFromCookie, err := c.Cookie(config.Globals.Auth.OAuth2.StateCookie.Name)
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
	token, err := provider.serviceHooks.OAuth2Exchange(authCode)
	if err != nil {
		c.Logger().Error("OAuth2 exchange failed: ", err)
		return web.GenerateInternalServerError()
	}
	c.SetCookie(web.DeleteCookie(stateFromCookie.Name))
	// Save the OAuth2 token to the DB to the system user.
	var oauthToken models.OAuth2Token
	if err := oauthToken.FromOAuth2Token(token, provider.identity); err != nil {
		return fmt.Errorf("FromToken: %w", err)
	}
	if err := provider.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&oauthToken).Error; err != nil {
			return fmt.Errorf("oauth2 create: %w", err)
		}
		if err := provider.serviceHooks.OnOAuth2Callback(tx, c, provider.identity, &oauthToken); err != nil {
			return fmt.Errorf("OnOAuth2Callback: %w", err)
		}
		return nil
	}); err != nil {
		c.Logger().Error("OnOAuth2Callback tx failed: ", err)
		return web.GenerateInternalServerError()
	}
	return c.Redirect(http.StatusTemporaryRedirect, provider.generateResponseRedirectUrl(stateUrlMessage.RedirectTo))
}
