package stats

import (
	"backend/modules/api/endpoints/messages/stats"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func handleLatestStatsPushRequest(logger *zap.Logger, dbConn *gorm.DB, message *stats.LatestStatsPushRequest) *stats.LatestStatsPushResponse {
	discordStats, err := fetchDiscordLatestStats(dbConn, message.DiscordGuild)
	if err != nil {
		logger.Error("Error fetching Discord latest stats", zap.String("guild_name", message.DiscordGuild), zap.Error(err))
	}
	youtubeStats, err := fetchYoutubeLatestStats(dbConn, message.YoutubeChannel)
	if err != nil {
		logger.Error("Error fetching Youtube latest stats", zap.String("channel_name", message.YoutubeChannel), zap.Error(err))
	}
	githubStats, err := fetchGithubLatestStats(dbConn, message.GithubRepo)
	if err != nil {
		logger.Error("Error fetching Github latest stats", zap.String("repo_name", message.GithubRepo), zap.Error(err))
	}
	return &stats.LatestStatsPushResponse{
		DiscordStats: discordStats,
		YoutubeStats: youtubeStats,
		GithubStats:  githubStats,
	}
}
