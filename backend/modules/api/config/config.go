package config

import (
	"backend/x/db"
	"backend/x/web"

	"golang.org/x/oauth2"
)

var (
	GithubOAuth2Endpoint = oauth2.Endpoint{
		AuthURL:   "https://github.com/login/oauth/authorize",
		TokenURL:  "https://github.com/login/oauth/access_token",
		AuthStyle: oauth2.AuthStyleInParams,
	}
)

type DiscordConsts struct {
	GuildID   string `yaml:"guild_id"`
	VIPRoleID string `yaml:"vip_role_id"`
}

type OAuth2ProviderConfig struct {
	RedirectUrl  string   `yaml:"redirect_url"`
	ClientId     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
}

func (conf *OAuth2ProviderConfig) ToOAuth2Config(endpoint oauth2.Endpoint) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  conf.RedirectUrl,
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Scopes:       conf.Scopes,
		Endpoint:     endpoint,
	}
}

type Config struct {
	Server struct {
		ListenAddr    string  `yaml:"listenAddr"`
		ListenPort    int     `yaml:"listenPort"`
		KeyPath       *string `yaml:"key_path"`
		FullchainPath *string `yaml:"fullchain_path"`
		AllowedOrigin string  `yaml:"allowed_origin"`
	} `yaml:"server"`
	Identity struct {
		Salt   string `yaml:"salt"`
		Secret string `yaml:"secret"`
	} `yaml:"identity"`
	Api struct {
		ClientToken string `yaml:"client_token"`
		ServerToken string `yaml:"server_token"`
	} `yaml:"api"`
	Auth struct {
		RedirectParam  string                   `yaml:"redurect_param"`
		RedirectHost   string                   `yaml:"redirect_host"`
		SessionCookies web.SessionCookiesConfig `yaml:"session_cookies"`
		OAuth2         struct {
			StateCookie web.CookieConfig `yaml:"state_cookie"`
			Config      struct {
				Discord OAuth2ProviderConfig `yaml:"discord"`
				Github  OAuth2ProviderConfig `yaml:"github"`
			} `yaml:"config"`
		} `yaml:"oauth2"`
	} `yaml:"auth"`
	Consts struct {
		Discord DiscordConsts `yaml:"discord"`
	} `yaml:"consts"`
	Database db.Config `yaml:"database"`
}

func (config *Config) TLSConfigured() bool {
	return config.Server.FullchainPath != nil && config.Server.KeyPath != nil
}

var Globals Config
