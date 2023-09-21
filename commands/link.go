package commands

import "github.com/bwmarrin/discordgo"

var (
	LinkOpts = []*discordgo.ApplicationCommandOption{
		{
			Name:        "url",
			Description: "URL to send",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Northstar Github",
					Value: "github",
				},
				{
					Name:  "Northstar Wiki",
					Value: "wiki",
				},
				{
					Name:  "hummusbird Hosting Tutorial",
					Value: "host",
				},
			},
		},
	}
)

func LinkCmdHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var sendString string
	o := i.ApplicationCommandData().Options[0].StringValue()

	if o == "github" {
		sendString = "https://github.com/R2Northstar"
	} else if o == "wiki" {
		sendString = "https://r2northstar.gitbook.io/r2northstar-wiki/"
	} else if o == "host" {
		sendString = "https://youtu.be/EZ3w2Nl9SZo"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: sendString,
		},
	})

}
