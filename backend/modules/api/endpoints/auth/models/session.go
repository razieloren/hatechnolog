package models

import (
	"backend/x/identity"
	"backend/x/messages"
	"backend/x/web"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const (
	SessionTokenLength = 32
	// A year.
	SessionLength = time.Hour * 24 * 365
)

type Session struct {
	gorm.Model
	UserID uint
	Token  string    `gorm:"UNIQUE;NOT NULL;size:32"`
	Exipry time.Time `gorm:"NOT NULL"`
}

func (session *Session) TableName() string {
	return "api.session"
}

func (session *Session) HasExpired() bool {
	return time.Now().UTC().After(session.Exipry)
}

func (session *Session) Invalidate(dbConn *gorm.DB) error {
	return dbConn.Delete(session).Error
}

func (session *Session) FromEcho(dbConn *gorm.DB, identity *identity.Identity, cookieConf *web.CookieConfig, c echo.Context) (*User, error) {
	sessionCookie, err := c.Cookie(cookieConf.Name)
	if err != nil {
		return nil, fmt.Errorf("request does not contain a session cookie: %w", err)
	}
	value, err := web.ParseEncryptedCookieValue(identity, sessionCookie.Value)
	if err != nil {
		// Malformed cookie, mark to delete.
		c.SetCookie(web.DeleteCookie(cookieConf.Name))
		return nil, fmt.Errorf("error parsing session cookie value: %w", err)
	}
	userSession := messages.UserSessionCookieValue{}
	if err := proto.Unmarshal(value, &userSession); err != nil {
		// Malformed cookie, mark to delete.
		c.SetCookie(web.DeleteCookie(cookieConf.Name))
		return nil, fmt.Errorf("error unmrashaling session value: %w", err)
	}
	if err := dbConn.Where(&Session{
		Token: userSession.Token,
	}).Take(session).Error; err != nil {
		c.SetCookie(web.DeleteCookie(cookieConf.Name))
		return nil, fmt.Errorf("no such session: %w", err)
	}
	if session.HasExpired() {
		// This session os not relevant anymore since it has expired.
		if err := session.Invalidate(dbConn); err != nil {
			return nil, fmt.Errorf("error invalidating session: %w", err)
		}
		c.SetCookie(web.DeleteCookie(cookieConf.Name))
		return nil, fmt.Errorf("session expired")
	}
	var user User
	if err := dbConn.Take(&user, session.UserID).Error; err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	if user.Handle != userSession.Handle {
		return nil, fmt.Errorf("mismatching cookie and session user")
	}
	return &user, nil
}
