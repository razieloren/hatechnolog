package stats

import (
	"backend/modules/api/endpoints/messages/stats"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func EndpointLatestStats(dbConn *gorm.DB, logger *zap.Logger, request *stats.LatestStatsRequest) *stats.LatestStatsResponse {
	discordStats, err := fetchDiscordLatestStats(dbConn, request.DiscordGuild)
	if err != nil {
		logger.Error("Error fetching Discord latest stats", zap.String("guild_name", request.DiscordGuild), zap.Error(err))
	}
	youtubeStats, err := fetchYoutubeLatestStats(dbConn, request.YoutubeChannel)
	if err != nil {
		logger.Error("Error fetching Youtube latest stats", zap.String("channel_name", request.YoutubeChannel), zap.Error(err))
	}
	githubStats, err := fetchGithubLatestStats(dbConn, request.GithubRepo)
	if err != nil {
		logger.Error("Error fetching Github latest stats", zap.String("repo_name", request.GithubRepo), zap.Error(err))
	}
	return &stats.LatestStatsResponse{
		DiscordStats: discordStats,
		YoutubeStats: youtubeStats,
		GithubStats:  githubStats,
	}
}
