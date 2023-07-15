package github

import "gorm.io/gorm"

type RepoStats struct {
	gorm.Model
	RepoName     string  `gorm:"NOT NULL;size:40"`
	Success      bool    `gorm:"NOT NULL"`
	Error        *string `gorm:"default:NULL"`
	Contributors uint64  `gorm:"NOT NULL"`
	Commits      uint64  `gorm:"NOT NULL"`
}

func (repoStats *RepoStats) TableName() string {
	return "query.github_repo_stats"
}

func AutoMigrate(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&RepoStats{},
	)
}
