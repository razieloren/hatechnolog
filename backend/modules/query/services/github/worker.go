package github

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Work(logger *zap.Logger, dbConn *gorm.DB, config interface{}) {
	githubConfig, ok := config.(*Config)
	if !ok {
		logger.Error("Bad config type received")
		return
	}
	githubService := NewService(githubConfig.ApiKey)
	for _, repoParams := range githubConfig.TargetRepos {
		lastErr := ""
		stats, err := githubService.QueryRepoStats(&repoParams)
		if err != nil {
			logger.Error("Error querying repo params", zap.String("repo_name", repoParams.Name), zap.Error(err))
			lastErr = err.Error()
		}
		if lastErr == "" {
			stats.Success = true
		} else {
			stats = &RepoStats{
				RepoName: repoParams.Name,
				Success:  false,
				Error:    &lastErr,
			}
		}
		if err := dbConn.Create(&stats).Error; err != nil {
			logger.Error("Error creating stats record", zap.String("repo_name", repoParams.Name), zap.Error(err))
		}
	}
}
