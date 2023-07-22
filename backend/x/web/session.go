package web

import (
	"backend/x/identity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SessionCookiesConfig struct {
	Session     CookieConfig `yaml:"session"`
	SessionFlag CookieConfig `yaml:"session_flag"`
}

func (config *SessionCookiesConfig) Get(c echo.Context) (*http.Cookie, error) {
	return c.Cookie(config.Session.Name)
}

func (config *SessionCookiesConfig) Set(c echo.Context, identity *identity.Identity, sessionBytes []byte) error {
	flagCookie, err := CreateCookie(&config.SessionFlag, []byte("1"))
	if err != nil {
		return fmt.Errorf("create cookie: %w", err)
	}
	sessionCookie, err := CreateEncryptedCookie(&config.Session, identity, sessionBytes)
	if err != nil {
		return fmt.Errorf("create encrypted cookie: %w", err)
	}
	c.SetCookie(flagCookie)
	c.SetCookie(sessionCookie)
	return nil
}

func (config *SessionCookiesConfig) Delete(c echo.Context) {
	c.SetCookie(DeleteCookie(config.Session.Name))
	c.SetCookie(DeleteCookie(config.SessionFlag.Name))
}
