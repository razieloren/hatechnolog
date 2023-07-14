package youtube

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Work(logger *zap.Logger, dbConn *gorm.DB, config interface{}) {
	youtubeConfig, ok := config.(*Config)
	if !ok {
		logger.Error("Bad config type received")
		return
	}
	youtubeService, err := NewService(youtubeConfig.ApiKey)
	if err != nil {
		logger.Error("Could not create Youtube service", zap.Error(err))
		return
	}
	for _, channelParams := range youtubeConfig.TargetChannels {
		lastErr := ""
		stats, err := youtubeService.QueryChannelStats(&channelParams)
		if err != nil {
			logger.Error("Error querying channel params", zap.String("channel_name", channelParams.Name), zap.Error(err))
			lastErr = err.Error()
		}
		if lastErr == "" {
			stats.Success = true
		} else {
			stats = &ChannelStats{
				ChannelName: channelParams.Name,
				Success:     false,
				Error:       &lastErr,
			}
		}
		if err := dbConn.Create(&stats).Error; err != nil {
			logger.Error("Error creating stats record", zap.String("channel_name", channelParams.Name), zap.Error(err))
		}
	}
}
