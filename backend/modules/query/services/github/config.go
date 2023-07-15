package github

type RepoQueryParams struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}

type Config struct {
	ApiKey      string            `yaml:"apiKey"`
	TargetRepos []RepoQueryParams `yaml:"targetRepos"`
}
