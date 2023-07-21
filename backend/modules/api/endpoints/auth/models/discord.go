package models

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
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
	IsSupporter       bool `gorm:"NOT NULL"`
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

func (discordUser *DiscordUser) checkGuildMembership(client *discordgo.Session, guildID string) error {
	for {
		afterId := ""
		apiUsersGuilds, err := client.UserGuilds(maxUserGuildsPerRequest, "", afterId)
		if err != nil {
			return fmt.Errorf("error fetching user guilds, after='%s', err=%w", afterId, err)
		}
		for _, apiUserGuild := range apiUsersGuilds {
			if apiUserGuild.ID == guildID {
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

func (discordUser *DiscordUser) checkRoleStatus(client *discordgo.Session, guildID, supporterRoleID, VIPRoleID string) error {
	memberApi, err := client.UserGuildMember(guildID)
	if err != nil {
		return fmt.Errorf("error fetching guild member: %w", err)
	}
	if memberApi.Avatar != "" {
		avatarUrl := memberApi.AvatarURL(avatarSize)
		discordUser.GuildAvatarURL = &avatarUrl
	}
	for _, roleID := range memberApi.Roles {
		if roleID == supporterRoleID {
			discordUser.IsSupporter = true
		} else if roleID == VIPRoleID {
			discordUser.IsVIP = true
		}
	}
	return nil
}

func (discordUser *DiscordUser) FromAPI(dbConn *gorm.DB, config *oauth2.Config, token *OAuth2Token, guildID, supporterRoleID, VIPRoleID string) error {
	cbError, newTokenError, saveTokenError := token.SafeAction(
		dbConn, config, func(client *http.Client) error {
			discordClient, err := discordgo.New("")
			if err != nil {
				return fmt.Errorf("new discord client: %w", err)
			}
			discordClient.Client = client
			apiUser, err := discordClient.User("@me")
			if err != nil {
				return fmt.Errorf("get discord @me: %w", err)
			}
			if err := dbConn.Where(&DiscordUser{
				DiscordID: apiUser.ID,
			}).Take(discordUser).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("error fetching old discord user: %w", err)
				}
			} else {
				// This user has been here before, we need to delete his old token.
				if err := dbConn.Delete(&OAuth2Token{}, discordUser.OAuth2TokenID).Error; err != nil {
					return fmt.Errorf("error deleting olad oauth2 token: %w", err)
				}
				// All the other values will be updated from the API.
			}
			discordUser.OAuth2TokenID = token.ID
			discordUser.DiscordID = apiUser.ID
			discordUser.Username = apiUser.Username
			discordUser.Email = apiUser.Email
			discordUser.AvatarURL = apiUser.AvatarURL(avatarSize)
			discordUser.MfaEnabled = apiUser.MFAEnabled
			discordUser.EmailVerified = apiUser.Verified
			discordUser.HatechnologMember = false
			discordUser.IsSupporter = false
			discordUser.IsVIP = false
			if err := discordUser.checkGuildMembership(discordClient, guildID); err != nil {
				return fmt.Errorf("check guild memebership: %w", err)
			}
			if discordUser.HatechnologMember {
				// Only if member if Hatechnolog server, check for special roles.
				if err := discordUser.checkRoleStatus(discordClient, guildID, supporterRoleID, VIPRoleID); err != nil {
					return fmt.Errorf("check role status: %w", err)
				}
			}
			if err := dbConn.Save(discordUser).Error; err != nil {
				return fmt.Errorf("error saving discord user to the DB: %w", err)
			}
			return nil
		})
	if cbError != nil {
		return fmt.Errorf("main discord cb failed: %w", cbError)
	}
	if newTokenError != nil {
		return fmt.Errorf("new discord token error: %w", newTokenError)
	}
	if saveTokenError != nil {
		return fmt.Errorf("save discord token error: %w", newTokenError)
	}
	return nil
}
