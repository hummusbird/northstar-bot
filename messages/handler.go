package messages

import (
	"log"
	"northstar-bot/commands"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Put anything related to handling individual user messages here. Avoids needing multiple handlers for this
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from self
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to pings with "what" (important feature)
	if m.Mentions != nil {
		for _, mention := range m.Mentions {
			if mention.ID == s.State.User.ID {
				s.ChannelMessageSendReply(m.ChannelID, "what", m.Reference())
				return
			}
		}
	}

	// Legacy prefix command handling
	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		log.Fatal("[FATAL] No prefix loaded")
	}

	var (
		statusCmd = prefix + "status"
		//redeemCmd = prefix + "redeem"
		githubCmd = prefix + "github"
		hostCmd   = prefix + "host"
		wikiCmd   = prefix + "wiki"
		helpCmd   = prefix + "help"
		infoCmd   = prefix + "info"
		plvCmd    = prefix + "playlistvars"
		mapsCmd   = prefix + "maps"
		modesCmd  = prefix + "modes"
		searchCmd = prefix + "search"
	)

	switch {
	case strings.HasPrefix(m.Content, statusCmd):
		commands.LegacyStatusHandler(s, m)
		return
	/*
		case strings.HasPrefix(m.Content, redeemCmd):
			commands.LegacyRedeemHandler(s, m)
			return
	*/
	case strings.HasPrefix(m.Content, githubCmd):
		s.ChannelMessageSendReply(m.ChannelID, "https://github.com/R2Northstar", m.Reference())
		return
	case strings.HasPrefix(m.Content, hostCmd):
		s.ChannelMessageSendReply(m.ChannelID, "https://youtu.be/EZ3w2Nl9SZo", m.Reference())
		return
	case strings.HasPrefix(m.Content, wikiCmd):
		s.ChannelMessageSendReply(m.ChannelID, "https://r2northstar.gitbook.io/r2northstar-wiki/", m.Reference())
		return
	case strings.HasPrefix(m.Content, helpCmd):
		helpHandler(s, m, prefix)
		return
	case strings.HasPrefix(m.Content, infoCmd):
		commands.LegacyInfoHandler(s, m)
		return
	case strings.HasPrefix(m.Content, plvCmd):
		s.ChannelMessageSendReply(m.ChannelID, commands.AllPlaylistVarString, m.Reference())
		return
	case strings.HasPrefix(m.Content, mapsCmd):
		s.ChannelMessageSendReply(m.ChannelID, commands.AllMapsString, m.Reference())
		return
	case strings.HasPrefix(m.Content, modesCmd):
		s.ChannelMessageSendReply(m.ChannelID, commands.AllModesString, m.Reference())
		return
	case strings.HasPrefix(m.Content, searchCmd):
		commands.LegacySearchHandler(s, m, m.Content)
		return
	}

}
