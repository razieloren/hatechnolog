package models

import (
	"backend/modules/api/config"
	"backend/x/identity"
	"errors"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/ravener/discord-oauth2"
	"gorm.io/gorm"
)

const (
	avatarSize              = "512"
	maxUserGuildsPerRequest = 100
)

type DiscordUser struct {
	gorm.Model
	OAuth2TokenID     uint `gorm:"NOT NULL"`
	OAuth2Token       OAuth2Token
	DiscordID         string `gorm:"UNIQUE;NOT NULL"`
	Username          string `gorm:"UNIQUE;NOT NULL"`
	Email             string `gorm:"UNIQUE;NOT NULL"`
	AvatarURL         string `gorm:"NOT NULL"`
	GuildAvatarURL    *string
	MfaEnabled        bool `gorm:"NOT NULL"`
	EmailVerified     bool `gorm:"NOT NULL"`
	HatechnologMember bool `gorm:"NOT NULL"`
	IsVIP             bool `gorm:"NOT NULL"`
}

func (discordUser *DiscordUser) TableName() string {
	return "api.discord_user"
}

func (discordUser *DiscordUser) GetAvatar() string {
	if discordUser.GuildAvatarURL != nil {
		return *discordUser.GuildAvatarURL
	}
	return discordUser.AvatarURL
}

func (discordUser *DiscordUser) loadGuildMembershipFromAPI(client *discordgo.Session) error {
	for {
		afterId := ""
		apiUsersGuilds, err := client.UserGuilds(maxUserGuildsPerRequest, "", afterId)
		if err != nil {
			return fmt.Errorf("user guilds: %w", err)
		}
		for _, apiUserGuild := range apiUsersGuilds {
			if apiUserGuild.ID == config.Globals.Consts.Discord.GuildID {
				discordUser.HatechnologMember = true
				break
			}
		}
		if discordUser.HatechnologMember || len(apiUsersGuilds) < maxUserGuildsPerRequest {
			break
		}
		afterId = apiUsersGuilds[len(apiUsersGuilds)-1].ID
	}
	return nil
}

func (discordUser *DiscordUser) loadRoleStatusFromAPI(client *discordgo.Session) error {
	memberApi, err := client.UserGuildMember(config.Globals.Consts.Discord.GuildID)
	if err != nil {
		return fmt.Errorf("user guild member: %w", err)
	}
	if memberApi.Avatar != "" {
		avatarUrl := memberApi.AvatarURL(avatarSize)
		discordUser.GuildAvatarURL = &avatarUrl
	}
	for _, roleID := range memberApi.Roles {
		if roleID == config.Globals.Consts.Discord.VIPRoleID {
			discordUser.IsVIP = true
		}
	}
	return nil
}

func (discordUser *DiscordUser) updateFromAPI(dbConn *gorm.DB, identity *identity.Identity, token *OAuth2Token) error {
	cbError, newTokenError, saveTokenError := token.SafeAction(
		dbConn, identity,
		config.Globals.Auth.OAuth2.Config.Discord.ToOAuth2Config(discord.Endpoint),
		func(client *http.Client) error {
			discordClient, err := discordgo.New("")
			if err != nil {
				return fmt.Errorf("new discord: %w", err)
			}
			discordClient.Client = client
			apiUser, err := discordClient.User("@me")
			if err != nil {
				return fmt.Errorf("user: %w", err)
			}
			if err := dbConn.Where(&DiscordUser{
				DiscordID: apiUser.ID,
			}).Take(discordUser).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("take discord user: %w", err)
				}
				discordUser.OAuth2TokenID = token.ID
			} else {
				// A new token has been used (new login) we need to delete the old token.
				if discordUser.OAuth2TokenID != token.ID {
					if err := dbConn.Delete(&OAuth2Token{}, discordUser.OAuth2TokenID).Error; err != nil {
						return fmt.Errorf("delete oauth2: %w", err)
					}
					discordUser.OAuth2TokenID = token.ID
				}
				// All the other values will be updated from the API.
			}
			discordUser.DiscordID = apiUser.ID
			discordUser.Username = apiUser.Username
			discordUser.Email = apiUser.Email
			discordUser.AvatarURL = apiUser.AvatarURL(avatarSize)
			discordUser.MfaEnabled = apiUser.MFAEnabled
			discordUser.EmailVerified = apiUser.Verified
			discordUser.HatechnologMember = false
			discordUser.IsVIP = false
			if err := discordUser.loadGuildMembershipFromAPI(discordClient); err != nil {
				return fmt.Errorf("check guild memebership: %w", err)
			}
			if discordUser.HatechnologMember {
				// Only if member if Hatechnolog server, check for special roles.
				if err := discordUser.loadRoleStatusFromAPI(discordClient); err != nil {
					return fmt.Errorf("check role status: %w", err)
				}
			}
			if err := dbConn.Save(discordUser).Error; err != nil {
				return fmt.Errorf("save discord user: %w", err)
			}
			return nil
		})
	if cbError != nil {
		return fmt.Errorf("cb: %w", cbError)
	}
	if newTokenError != nil {
		return fmt.Errorf("new token: %w", newTokenError)
	}
	if saveTokenError != nil {
		return fmt.Errorf("save token: %w", newTokenError)
	}
	return nil
}

func (discordUser *DiscordUser) FromOAuth2Token(dbConn *gorm.DB, identity *identity.Identity, token *OAuth2Token) error {
	return discordUser.updateFromAPI(dbConn, identity, token)
}

func (discordUser *DiscordUser) FromUser(dbConn *gorm.DB, identity *identity.Identity, user *User) error {
	var oauthToken OAuth2Token
	if err := dbConn.Take(&oauthToken, user.DiscordUser.OAuth2TokenID).Error; err != nil {
		return fmt.Errorf("take oauth2 token: %w", err)
	}
	return discordUser.updateFromAPI(dbConn, identity, &oauthToken)
}
