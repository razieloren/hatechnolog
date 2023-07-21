package models

import (
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
	AccessToken  string    `gorm:"NOT NULL"`
	RefreshToken string    `gorm:"NOT NULL"`
	TokenType    string    `gorm:"NOT NULL"`
	TokenExpiry  time.Time `gorm:"NOT NULL"`
}

func (token *OAuth2Token) TableName() string {
	return "api.oauth2_token"
}

func (token *OAuth2Token) SafeAction(dbConn *gorm.DB, conf *oauth2.Config,
	action OAuth2Action) (error, error, error) {
	oldToken := &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Expiry:       token.TokenExpiry,
	}
	tokenSource := conf.TokenSource(oauth2.NoContext, oldToken)
	client := oauth2.NewClient(oauth2.NoContext, tokenSource)
	callbackError := action(client)
	newToken, tokenErr := tokenSource.Token()
	if tokenErr == nil {
		token.AccessToken = newToken.AccessToken
		token.RefreshToken = newToken.RefreshToken
		token.TokenType = newToken.TokenType
		token.TokenExpiry = newToken.Expiry
		if err := dbConn.Save(&token).Error; err != nil {
			return callbackError, nil, err
		}
		return callbackError, nil, nil
	}
	return callbackError, tokenErr, nil
}
