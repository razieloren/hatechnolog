package github

import (
	"context"
	"fmt"
	"net/http"
	"time"

	gogithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Service struct {
	session *gogithub.Client
}

func NewService(apiKey string) *Service {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &Service{
		session: gogithub.NewClient(tc),
	}
}

func (service *Service) getContributosStats(owner, repo string, retry bool) ([]*gogithub.ContributorStats, error) {
	ctx := context.Background()
	stats, resp, err := service.session.Repositories.ListContributorsStats(ctx, owner, repo)
	if resp.StatusCode == http.StatusAccepted {
		if retry {
			return nil, fmt.Errorf("stats_timed_out")
		}
		// This means GitHub started processing the request, and it will be ready in a few seconds.
		time.Sleep(statsRetryAfter)
		return service.getContributosStats(owner, repo, true)
	}
	if err != nil {
		return nil, fmt.Errorf("list_contributors_stats: %w", err)
	}
	return stats, nil
}

func (service *Service) QueryRepoStats(params *RepoQueryParams) (*RepoStats, error) {
	ctx := context.Background()
	repos, _, err := service.session.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, fmt.Errorf("repo_list: %w", err)
	}
	for _, repo := range repos {
		if repo.Name == nil {
			continue
		}
		if *repo.Name == params.Name {
			stats, err := service.getContributosStats(params.Owner, params.Name, false)
			if err != nil {
				return nil, fmt.Errorf("get_contrib_stats: %w", err)
			}
			repoStats := RepoStats{
				RepoName: params.Name,
			}
			for _, stat := range stats {
				if stat.Total == nil {
					continue
				}
				repoStats.Contributors += 1
				repoStats.Commits += uint64(*stat.Total)
			}
			return &repoStats, nil
		}
	}
	return nil, fmt.Errorf("Could not find repo %s", params.Name)
}
