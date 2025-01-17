package core

import "os"

type OAuthConfig struct {
	ClientID string
	Secret   string
}

type Config struct {
	WebDev                 bool
	BaseURL                string
	Name                   string
	Owner                  string
	Key                    string
	GithubOAuthConfig      *OAuthConfig
	DiscordOAuthConfig     *OAuthConfig
	MesehubOAuthConfig     *OAuthConfig
	GoogleSiteVerification string
	CookiePath             string
	CookieDomain           string
	CookieSecure           bool
	CookieName             string
	RedisHost              string
	RedisPort              string
}

func CreateConfig() (*Config, error) {
	cfg := &Config{
		Name:                   os.Getenv("BLOCKEXCHANGE_NAME"),
		Owner:                  os.Getenv("BLOCKEXCHANGE_OWNER"),
		WebDev:                 os.Getenv("WEBDEV") == "true",
		BaseURL:                os.Getenv("BASE_URL"),
		CookiePath:             os.Getenv("BLOCKEXCHANGE_COOKIE_PATH"),
		CookieDomain:           os.Getenv("BLOCKEXCHANGE_COOKIE_DOMAIN"),
		CookieSecure:           os.Getenv("BLOCKEXCHANGE_COOKIE_SECURE") == "true",
		CookieName:             "blockexchange",
		GoogleSiteVerification: os.Getenv("GOOGLE_SITE_VERIFICATION"),
		RedisHost:              os.Getenv("REDIS_HOST"),
		RedisPort:              os.Getenv("REDIS_PORT"),
	}

	if os.Getenv("DISCORD_APP_ID") != "" {
		cfg.DiscordOAuthConfig = &OAuthConfig{
			ClientID: os.Getenv("DISCORD_APP_ID"),
			Secret:   os.Getenv("DISCORD_APP_SECRET"),
		}
	}

	if os.Getenv("GITHUB_APP_ID") != "" {
		cfg.GithubOAuthConfig = &OAuthConfig{
			ClientID: os.Getenv("GITHUB_APP_ID"),
			Secret:   os.Getenv("GITHUB_APP_SECRET"),
		}
	}

	if os.Getenv("MESEHUB_APP_ID") != "" {
		cfg.MesehubOAuthConfig = &OAuthConfig{
			ClientID: os.Getenv("MESEHUB_APP_ID"),
			Secret:   os.Getenv("MESEHUB_APP_SECRET"),
		}
	}

	return cfg, nil
}
