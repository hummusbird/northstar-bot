package messages

import (
	"github.com/bwmarrin/discordgo"
)

func helpHandler(s *discordgo.Session, m *discordgo.MessageCreate, p string) {
	helpString := "```Diff\n" +
		"+ Here is a list of all available commands: \n\n" +
		"-[HELP]\n" +
		p + "help                   - displays this message\n" +
		//p + "prefix [prefix]        - allows an admin to set the prefix\n" +
		"-[Northstar]\n" +
		p + "status                 - a general overview of northstar.tf\n" +
		p + "search title [string]  - searches server titles\n" +
		p + "search mode [gamemode] - searches all servers running that mode\n" +
		p + "search map [map]       - searches all servers running that map\n" +
		"-[Titanfall 2 Lists]\n" +
		p + "playlistvars           - lists some useful playlist vars\n" +
		p + "modes                  - lists all Titanfall 2 gamemodes\n" +
		p + "maps                   - lists all Titanfall 2 maps\n" +
		"-[Links]\n" +
		p + "info                   - display info about the bot\n" +
		p + "host                   - links hummusbird's server tutorial\n" +
		p + "git                    - links the github\n" +
		p + "wiki                   - links the wiki\n\n" +
		"+ ## This bot also has slash commands ##\n" +
		"```"

	s.ChannelMessageSendReply(m.ChannelID, helpString, m.Reference())
}
