package discord

import (
	"fmt"
	"sort"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	session *discordgo.Session
}

func NewService(botToken string) (*Service, error) {
	session, err := discordgo.New(fmt.Sprintf(botTokenFormat, botToken))
	if err != nil {
		return nil, fmt.Errorf("new_session: %w", err)
	}
	return &Service{
		session: session,
	}, nil
}

func (service *Service) ListGuildMembers(guildId string) ([]*discordgo.Member, error) {
	lastMemberId := ""
	var allMembers []*discordgo.Member
	for {
		members, err := service.session.GuildMembers(
			guildId, lastMemberId, maxGuildMembersPerCall)
		if err != nil {
			return nil, fmt.Errorf("guild_members: %w", err)
		}
		allMembers = append(allMembers, members...)
		if len(members) < maxGuildMembersPerCall {
			break
		}
		lastMemberId = members[maxGuildMembersPerCall-1].User.ID
	}
	sort.Slice(allMembers, func(i, j int) bool {
		return allMembers[i].JoinedAt.Before(allMembers[j].JoinedAt)
	})
	return allMembers, nil
}

func (service *Service) QueryGuildStats(params *GuildQueryParams) (*GuildStats, error) {
	lastGuildId := ""
	for {
		userGuilds, err := service.session.UserGuilds(maxGuildsPerCall, "", lastGuildId)
		if err != nil {
			return nil, fmt.Errorf("user_guilds: %w", err)
		}
		for _, userGuild := range userGuilds {
			if params.Name == userGuild.Name {
				creationTime, err := discordgo.SnowflakeTimestamp(userGuild.ID)
				if err != nil {
					return nil, fmt.Errorf("snowflake_timestamp: %w", err)
				}
				stats := GuildStats{
					GuildName:           params.Name,
					LastTimeReference:   creationTime,
					NewMemberPeriodDays: params.NewMemberPeriodDays,
				}
				members, err := service.ListGuildMembers(userGuild.ID)
				if err != nil {
					return nil, fmt.Errorf("list_guild_members: %w", err)
				}
				for _, member := range members {
					stats.TotalMembers += 1
					if !member.User.Bot {
						stats.TotalHumans += 1
						joinedDelta := time.Now().UTC().Sub(member.JoinedAt)
						if joinedDelta <
							oneDay*time.Duration(params.NewMemberPeriodDays) {
							stats.NewHumans += 1
						}
						if joinedDelta <
							oneDay*time.Duration(params.AvgJoinTimePeriodDays) {
							stats.JoinDeltasSum += member.JoinedAt.Sub(stats.LastTimeReference)
							stats.JoinUserCount += 1
						}
						stats.LastTimeReference = member.JoinedAt
					}
				}
				if stats.JoinUserCount > 0 {
					stats.JoinAvgSec = stats.JoinDeltasSum.Seconds() / float64(stats.JoinUserCount)
				}
				return &stats, nil
			}
		}
		if len(userGuilds) < maxGuildsPerCall {
			return nil, fmt.Errorf("Could not find guild")
		}
		lastGuildId = userGuilds[maxGuildsPerCall-1].ID
	}

}
