package discord

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Work(logger *zap.Logger, dbConn *gorm.DB, config interface{}) {
	discordConfig, ok := config.(*Config)
	if !ok {
		logger.Error("Bad config type received")
		return
	}
	discordService, err := NewService(discordConfig.BotToken)
	if err != nil {
		logger.Error("Could not create Discord service", zap.Error(err))
		return
	}
	for _, guildParams := range discordConfig.TargetGuilds {
		lastErr := ""
		stats, err := discordService.QueryGuildStats(&guildParams)
		if err != nil {
			logger.Error("Error querying guild params", zap.String("guild_name", guildParams.Name), zap.Error(err))
			lastErr = err.Error()
		}
		if lastErr == "" {
			stats.Success = true
		} else {
			stats = &GuildStats{
				GuildName: guildParams.Name,
				Success:   false,
				Error:     &lastErr,
			}
		}
		if err := dbConn.Create(&stats).Error; err != nil {
			logger.Error("Error creating stats record", zap.String("guild_name", guildParams.Name), zap.Error(err))
		}
	}
}
