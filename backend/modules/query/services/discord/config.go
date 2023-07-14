package discord

type GuildQueryParams struct {
	Name                  string `yaml:"name"`
	NewMemberPeriodDays   int    `yaml:"newMemberPeriodDays"`
	AvgJoinTimePeriodDays int    `yaml:"avgJoinTimePeriodDays"`
}

type Config struct {
	BotToken     string             `yaml:"botToken"`
	TargetGuilds []GuildQueryParams `yaml:"targetGuilds"`
}
