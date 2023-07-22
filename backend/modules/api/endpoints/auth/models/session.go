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

func (session *Session) FromEcho(dbConn *gorm.DB, identity *identity.Identity, cookies *web.SessionCookiesConfig, c echo.Context) (*User, error) {
	sessionCookie, err := cookies.Get(c)
	if err != nil {
		return nil, fmt.Errorf("no session cookie: %w", err)
	}
	value, err := web.ParseEncryptedCookieValue(identity, sessionCookie.Value)
	if err != nil {
		// Malformed cookie, mark to delete.
		cookies.Delete(c)
		return nil, fmt.Errorf("parse encrypted cookie: %w", err)
	}
	userSession := messages.UserSessionCookieValue{}
	if err := proto.Unmarshal(value, &userSession); err != nil {
		// Malformed cookie, mark to delete.
		cookies.Delete(c)
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	if err := dbConn.Where(&Session{
		Token: userSession.Token,
	}).Take(session).Error; err != nil {
		cookies.Delete(c)
		return nil, fmt.Errorf("no such session: %w", err)
	}
	if session.HasExpired() {
		// This session os not relevant anymore since it has expired.
		if err := session.Invalidate(dbConn); err != nil {
			return nil, fmt.Errorf("invalidate: %w", err)
		}
		cookies.Delete(c)
		return nil, fmt.Errorf("session expired")
	}
	var user User
	if err := dbConn.Take(&user, session.UserID).Error; err != nil {
		return nil, fmt.Errorf("no such user: %w", err)
	}
	if user.Handle != userSession.Handle {
		return nil, fmt.Errorf("bad session handle")
	}
	return &user, nil
}
