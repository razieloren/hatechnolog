package discord

import "time"

const (
	maxGuildsPerCall       = 100
	maxGuildMembersPerCall = 1000
	oneDay                 = time.Hour * 24
	botTokenFormat         = "Bot %s"
)
