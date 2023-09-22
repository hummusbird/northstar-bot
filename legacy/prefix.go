package legacy

import (
	"log"
	"northstar-bot/commands"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func PrefixCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		log.Fatal("[FATAL] No prefix loaded")
	}

	statusCmd := prefix + "status"
	redeemCmd := prefix + "redeem"
	githubCmd := prefix + "github"
	hostCmd := prefix + "host"
	wikiCmd := prefix + "wiki"

	if strings.HasPrefix(m.Content, statusCmd) {
		commands.LegacyStatusHandler(s, m)
	} else if strings.HasPrefix(m.Content, redeemCmd) {
		commands.LegacyRedeemHandler(s, m)
	} else if strings.HasPrefix(m.Content, githubCmd) {
		s.ChannelMessageSendReply(m.ChannelID, "https://github.com/R2Northstar", m.Reference())
	} else if strings.HasPrefix(m.Content, hostCmd) {
		s.ChannelMessageSendReply(m.ChannelID, "https://youtu.be/EZ3w2Nl9SZo", m.Reference())
	} else if strings.HasPrefix(m.Content, wikiCmd) {
		s.ChannelMessageSendReply(m.ChannelID, "https://r2northstar.gitbook.io/r2northstar-wiki/", m.Reference())
	}

}
