package DiscordState

import "github.com/Rivalo/discordgo_cli"

//State is the current state of the attached client
type State struct {
	Guild    *discordgo.Guild
	Channel  *discordgo.Channel
	Channels []*discordgo.Channel
	Members  map[string]*discordgo.Member
	Messages map[string]*discordgo.Message
	Session  *Session
	Enabled  bool //Toggles State for Event handling
}

//Session contains the 'state' of the attached server
type Session struct {
	Username  string
	Password  string
	DiscordGo *discordgo.Session
	Guilds    []*discordgo.Guild
}
