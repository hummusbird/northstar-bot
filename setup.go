package main

import (
	"log"
	"northstar-bot/commands"
	"northstar-bot/util"

	"github.com/bwmarrin/discordgo"
)

var (
	commandsList = []*discordgo.ApplicationCommand{
		{
			Name:        "status", // Implemented
			Description: "A general overview of northstar.tf",
		},
		{
			Name:        "search", // Implemented
			Description: "Search all running Northstar servers",
			Options:     commands.SearchOpts,
		},
		{
			Name:        "host", // Implemented
			Description: "Link to hummusbird's server tutorial",
		},
		{
			Name:        "wiki", // Implemented
			Description: "Link to Northstar Wiki",
		},
		{
			Name:        "github", // Implemented
			Description: "Link to Northstar Github",
		},
		{
			Name:        "info", // Implemented
			Description: "Display information about the bot",
		},
		{
			Name:        "redeem",
			Description: "Redeem your ScorchBucks",
		},
		{
			Name:        "list", // Implemented
			Description: "List various things about Northstar",
			Options:     commands.ListOpts,
		},
		{
			Name:        "link", // Implemented
			Description: "Easy send URL of various Northstar things",
			Options:     commands.LinkOpts,
		},
		{
			Name:        "redeem",
			Description: "Redeem ScorchBucks",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"status": commands.StatusHandler,
		"info":   commands.InfoHandler,
		"list":   commands.ListCmdHandler,
		"link":   commands.LinkCmdHandler,
		"search": commands.SearchHandler,
		"redeem": commands.RedeemHandler,
	}
)

func registerCommands(s *discordgo.Session) {
	regCmd := make([]*discordgo.ApplicationCommand, len(commandsList))
	for i, v := range commandsList {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, s.State.Application.GuildID, v)
		if err != nil {
			log.Fatal(err)
		}
		regCmd[i] = cmd
	}
	log.Println("[INFO] All commands registered")
}

func setupAllHandlers(s *discordgo.Session) {

	// Attach type-specific handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		// Attach command handlers for slash commands
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
			/* Attach component handlers, such as handlers for buttons
			case discordgo.InteractionMessageComponent:
				if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
					h(s, i)
				}
			*/
		}
	})

	s.AddHandler(util.MessageCreateHandler)

	log.Println("[INFO] All handlers attached")
}

func Setup(s *discordgo.Session) {
	registerCommands(s)
	setupAllHandlers(s)
}
