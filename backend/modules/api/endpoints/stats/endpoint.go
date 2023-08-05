package stats

import (
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/stats"
	"backend/x/web"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndpointLatestStats(dbConn *gorm.DB, c echo.Context, request *stats.GetLatestStatsRequest) error {
	discordStats, err := fetchDiscordLatestStats(dbConn, request.DiscordGuild)
	if err != nil {
		c.Logger().Error("Error fetching Discord latest stats: ", err)
	}
	youtubeStats, err := fetchYoutubeLatestStats(dbConn, request.YoutubeChannel)
	if err != nil {
		c.Logger().Error("Error fetching Youtube latest stats: ", err)
	}
	githubStats, err := fetchGithubLatestStats(dbConn, request.GithubRepo)
	if err != nil {
		c.Logger().Error("Error fetching Github latest stats: ", err)
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetLatestStatsResponse{
			GetLatestStatsResponse: &stats.GetLatestStatsResponse{
				DiscordStats: discordStats,
				YoutubeStats: youtubeStats,
				GithubStats:  githubStats,
			},
		},
	})
}
