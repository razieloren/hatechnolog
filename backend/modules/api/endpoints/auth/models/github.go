package models

import "gorm.io/gorm"

type GithubUser struct {
	gorm.Model
	OAuth2TokenID uint `gorm:"NOT NULL"`
	OAuth2Token   OAuth2Token
	Username      string `gorm:"UNIQUE;NOT NULL"`
	Email         string `gorm:"UNIQUE;NOT NULL"`
}

func (githubUser *GithubUser) TableName() string {
	return "api.github_user"
}
