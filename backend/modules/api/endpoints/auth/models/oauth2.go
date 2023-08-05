package models

import (
	"backend/x/identity"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

const (
	ServiceDiscord = "discord"
	ServiceGithub  = "github"
)

type OAuth2Action func(*http.Client) error

type OAuth2Token struct {
	gorm.Model
	AccessIV        []byte    `gorm:"NOT NULL"`
	EncAccessToken  []byte    `gorm:"NOT NULL"`
	RefreshIV       []byte    `gorm:"NOT NULL"`
	EncRefreshToken []byte    `gorm:"NOT NULL"`
	TokenType       string    `gorm:"NOT NULL"`
	TokenExpiry     time.Time `gorm:"NOT NULL"`
}

func (token *OAuth2Token) TableName() string {
	return "api.oauth2_token"
}

func (token *OAuth2Token) FromOAuth2Token(oauth *oauth2.Token, identity *identity.Identity) error {
	accessIV, encAccess, err := identity.GCMAESEncrypt([]byte(oauth.AccessToken))
	if err != nil {
		return fmt.Errorf("GCMAESEncrypt access: %w", err)
	}
	refreshIV, encRefresh, err := identity.GCMAESEncrypt([]byte(oauth.RefreshToken))
	if err != nil {
		return fmt.Errorf("GCMAESEncrypt refresh: %w", err)
	}
	token.AccessIV = accessIV
	token.EncAccessToken = encAccess
	token.RefreshIV = refreshIV
	token.EncRefreshToken = encRefresh
	token.TokenType = oauth.TokenType
	token.TokenExpiry = oauth.Expiry
	return nil
}

func (token *OAuth2Token) SafeAction(dbConn *gorm.DB, identity *identity.Identity, conf *oauth2.Config,
	action OAuth2Action) (error, error, error) {
	decAccessToken, err := identity.GCMAESDecrypt(token.AccessIV, token.EncAccessToken)
	if err != nil {
		return fmt.Errorf("GCMAESDecrypt access: %w", err), nil, nil
	}
	decRefreshToken, err := identity.GCMAESDecrypt(token.RefreshIV, token.EncRefreshToken)
	if err != nil {
		return fmt.Errorf("GCMAESDecrypt refresh: %w", err), nil, nil
	}
	oldToken := &oauth2.Token{
		AccessToken:  string(decAccessToken),
		RefreshToken: string(decRefreshToken),
		TokenType:    token.TokenType,
		Expiry:       token.TokenExpiry,
	}
	tokenSource := conf.TokenSource(oauth2.NoContext, oldToken)
	client := oauth2.NewClient(oauth2.NoContext, tokenSource)
	callbackError := action(client)
	newToken, tokenErr := tokenSource.Token()
	if tokenErr == nil {
		if err := token.FromOAuth2Token(newToken, identity); err != nil {
			return callbackError, nil, fmt.Errorf("FromOAuth2Token: %w", err)
		}
		if err := dbConn.Save(&token).Error; err != nil {
			return callbackError, nil, err
		}
		return callbackError, nil, nil
	}
	return callbackError, tokenErr, nil
}
