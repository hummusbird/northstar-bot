package commands

import "github.com/bwmarrin/discordgo"

func InfoHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Northstar Servers Bot",
					Description: "Made by hummusbird\nRemade in Rust by H0L0\nRemade (again) in Go by cyrv6737\n\nA discord bot that displays the status of the northstar.tf servers\n\n**hummusbird**\nhttps://birb.cc/\n**H0L0**\nhttps://h0l0.cc/\n**cyrv6737**\nhttps://github.com/cyrv6737/",
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "https://northstar.tf/assets/logo_1k.png",
					},
					URL:   "https://github.com/hummusbird/northstar-bot",
					Color: 15999081,
				},
			},
		},
	})
}
