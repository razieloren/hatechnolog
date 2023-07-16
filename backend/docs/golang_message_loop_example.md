```go
func (router *Router) messageRouter(request *messages.MessageWrapper) (*messages.MessageWrapper, error) {
	wrapper := &messages.MessageWrapper{}
	switch message := request.Message.(type) {
	case *messages.MessageWrapper_DiscordLatestStatsRequest:
		response, err := router.handleDiscordLatestStatsRequest(message.DiscordLatestStatsRequest)
		if err != nil {
			return nil, fmt.Errorf("handle_discord_latest_stats_request: %w", err)
		}
		wrapper.Message = &messages.MessageWrapper_DiscordLatestStatsResponse{
			DiscordLatestStatsResponse: response,
		}
	case *messages.MessageWrapper_YoutubeLatestStatsRequest:
		response, err := router.handleYoutubeLatestStatsRequest(message.YoutubeLatestStatsRequest)
		if err != nil {
			return nil, fmt.Errorf("handle_youtube_latest_stats_request: %w", err)
		}
		wrapper.Message = &messages.MessageWrapper_YoutubeLatestStatsResponse{
			YoutubeLatestStatsResponse: response,
		}
	case *messages.MessageWrapper_GithubLatestStatsRequest:
		response, err := router.handleGithubLatestStatsRequest(message.GithubLatestStatsRequest)
		if err != nil {
			return nil, fmt.Errorf("handle_github_latest_stats_request: %w", err)
		}
		wrapper.Message = &messages.MessageWrapper_GithubLatestStatsResponse{
			GithubLatestStatsResponse: response,
		}
	default:
		return nil, fmt.Errorf("Unknown message revceived")
	}
	return wrapper, nil
}
```