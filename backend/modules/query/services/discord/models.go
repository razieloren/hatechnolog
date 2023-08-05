package discord

import (
	"time"

	"gorm.io/gorm"
)

type GuildStats struct {
	gorm.Model
	GuildName           string        `gorm:"NOT NULL;"`
	Success             bool          `gorm:"NOT NULL"`
	Error               *string       `gorm:"default:NULL"`
	NewHumans           int           `gorm:"NOT NULL"`
	TotalHumans         int           `gorm:"NOT NULL"`
	TotalMembers        int           `gorm:"NOT NULL"`
	NewMemberPeriodDays int           `gorm:"NOT NULL"`
	JoinAvgSec          float64       `gorm:"NOT NULL"`
	JoinUserCount       int           `gorm:"-"`
	JoinDeltasSum       time.Duration `gorm:"-"`
	LastTimeReference   time.Time     `gorm:"-"`
}

func (guildStats *GuildStats) TableName() string {
	return "query.discord_guild_stats"
}

func AutoMigrate(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&GuildStats{},
	)
}
