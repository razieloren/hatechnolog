package stats

import (
	"backend/modules/api/endpoints/messages/stats"
	"backend/modules/query/services/discord"
	"backend/modules/query/services/github"
	"backend/modules/query/services/youtube"
	"fmt"

	"gorm.io/gorm"
)

func fetchDiscordLatestStats(dbConn *gorm.DB, guildName string) (*stats.LatestDiscordStats, error) {
	var record discord.GuildStats
	if err := dbConn.Where(&discord.GuildStats{
		GuildName: guildName,
		Success:   true,
	}).Last(&record).Error; err != nil {
		return &stats.LatestDiscordStats{
			Valid: false,
		}, fmt.Errorf("Select discord_guild_stats: %w", err)
	}
	return &stats.LatestDiscordStats{
		TotalMembers:         uint32(record.TotalHumans),
		NewMembers:           uint32(record.NewHumans),
		NewMembersPeriodDays: uint32(record.NewMemberPeriodDays),
		JoinAvgSec:           float32(record.JoinAvgSec),
		Valid:                true,
	}, nil
}

func fetchYoutubeLatestStats(dbConn *gorm.DB, channelName string) (*stats.LatestYoutubeStats, error) {
	var record youtube.ChannelStats
	if err := dbConn.Where(&youtube.ChannelStats{
		ChannelName: channelName,
		Success:     true,
	}).Last(&record).Error; err != nil {
		return &stats.LatestYoutubeStats{
			Valid: false,
		}, fmt.Errorf("Select youtube_channel_stats: %w", err)
	}
	return &stats.LatestYoutubeStats{
		Subscribers: record.Subscribers,
		Views:       record.Views,
		Valid:       true,
	}, nil
}

func fetchGithubLatestStats(dbConn *gorm.DB, repoName string) (*stats.LatestGithubStats, error) {
	var record github.RepoStats
	if err := dbConn.Where(&github.RepoStats{
		RepoName: repoName,
		Success:  true,
	}).Last(&record).Error; err != nil {
		return &stats.LatestGithubStats{
			Valid: false,
		}, fmt.Errorf("Select github_repo_stats: %w", err)
	}
	return &stats.LatestGithubStats{
		Contributors: record.Contributors,
		Commits:      record.Commits,
		Valid:        true,
	}, nil
}
