package youtube

import "gorm.io/gorm"

type ChannelStats struct {
	gorm.Model
	ChannelName string  `gorm:"NOT NULL;"`
	Success     bool    `gorm:"NOT NULL"`
	Error       *string `gorm:"default:NULL"`
	Subscribers uint64  `gorm:"NOT NULL"`
	Views       uint64  `gorm:"NOT NULL"`
}

func (channelStats *ChannelStats) TableName() string {
	return "query.youtube_channel_stats"
}

func AutoMigrate(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&ChannelStats{},
	)
}
