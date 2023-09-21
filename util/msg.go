package util

import "github.com/bwmarrin/discordgo"

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Mentions != nil {
		for _, mention := range m.Mentions {
			if mention.ID == s.State.User.ID {
				s.ChannelMessageSendReply(m.ChannelID, "what", m.Reference())
			}
		}
	}
}
