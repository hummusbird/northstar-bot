package commands

import (
	"fmt"
	"northstar-bot/util"
	"sort"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	maxResults = 10
	SearchOpts = []*discordgo.ApplicationCommandOption{
		{
			Name:        "region",
			Description: "Search by Region",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "US East",
					Value: "US East",
				},
				{
					Name:  "US West",
					Value: "US West",
				},
				{
					Name:  "US_Central",
					Value: "US Central",
				},
				{
					Name:  "US South",
					Value: "US South",
				},
				{
					Name:  "EU North",
					Value: "EU North",
				},
				{
					Name:  "EU South",
					Value: "EU South",
				},
				{
					Name:  "EU East",
					Value: "EU East",
				},
				{
					Name:  "EU West",
					Value: "EU West",
				},
				{
					Name:  "AUS",
					Value: "AUS",
				},
				{
					Name:  "Africa",
					Value: "Africa",
				},
				{
					Name:  "Americas",
					Value: "Americas",
				},
				{
					Name:  "JPN",
					Value: "Asia East",
				},
			},
		},
		{
			Name:        "name",
			Description: "Search by server name",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
		{
			Name:        "map",
			Description: "Search by Map",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: func() []*discordgo.ApplicationCommandOptionChoice {
				var MapChoices []*discordgo.ApplicationCommandOptionChoice
				for _, r2map := range util.Maps[:25] {
					MapChoices = append(MapChoices, &discordgo.ApplicationCommandOptionChoice{
						Name:  r2map.ReadableName,
						Value: r2map.Name,
					})
				}
				return MapChoices
			}(),
		},
		{
			Name:        "map_sp",
			Description: "Search by Singleplayer Map",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: func() []*discordgo.ApplicationCommandOptionChoice {
				var MapChoices []*discordgo.ApplicationCommandOptionChoice
				for _, r2map := range util.Maps[25:] {
					MapChoices = append(MapChoices, &discordgo.ApplicationCommandOptionChoice{
						Name:  r2map.ReadableName,
						Value: r2map.Name,
					})
				}
				return MapChoices
			}(),
		},
		{
			Name:        "gamemode",
			Description: "Search by Vanilla Gamemodes",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: func() []*discordgo.ApplicationCommandOptionChoice {
				var ModeChoices []*discordgo.ApplicationCommandOptionChoice
				for _, r2map := range util.Playlists[:19] {
					ModeChoices = append(ModeChoices, &discordgo.ApplicationCommandOptionChoice{
						Name:  r2map.ReadableName,
						Value: r2map.Name,
					})
				}
				return ModeChoices
			}(),
		},
		{
			Name:        "gamemode_other",
			Description: "Search by gamemodes included in Northstar and Vanilla Featured",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: func() []*discordgo.ApplicationCommandOptionChoice {
				var ModeChoices []*discordgo.ApplicationCommandOptionChoice
				for _, r2map := range util.Playlists[19:] {
					ModeChoices = append(ModeChoices, &discordgo.ApplicationCommandOptionChoice{
						Name:  r2map.ReadableName,
						Value: r2map.Name,
					})
				}
				return ModeChoices
			}(),
		},
	}
)

func createSearchMessage(i *discordgo.InteractionCreate) string {
	var (
		messageString = "```Diff"
		data          []util.ServerInfo
		wg            sync.WaitGroup
	)

	data = util.ProcessAPIResp()

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Players > data[j].Players
	})

	if i.ApplicationCommandData().Options != nil {

		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, option := range i.ApplicationCommandData().Options {

				if option.Name == "region" {
					for i := len(data) - 1; i >= 0; i-- {
						if data[i].Region != option.StringValue() {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
				if option.Name == "name" {
					for i := len(data) - 1; i >= 0; i-- {
						if strings.Contains(strings.ToLower(data[i].Name), strings.ToLower(option.StringValue())) {
							continue
						} else {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
				if option.Name == "map" {
					for i := len(data) - 1; i >= 0; i-- {
						if data[i].Map != option.StringValue() {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
				if option.Name == "map_sp" {
					for i := len(data) - 1; i >= 0; i-- {
						if data[i].Map != option.StringValue() {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
				if option.Name == "gamemode" {
					for i := len(data) - 1; i >= 0; i-- {
						if data[i].Playlist != option.StringValue() {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
				if option.Name == "gamemode_other" {
					for i := len(data) - 1; i >= 0; i-- {
						if data[i].Playlist != option.StringValue() {
							data = append(data[:i], data[i+1:]...)
						}
					}
				}
			}
		}()

	} else {
		return "No options passed"
	}

	wg.Wait()

	messageString = messageString + fmt.Sprintf("\n+ %d servers were found - Showing up to %d results", len(data), maxResults)

	for index, result := range data {
		if index < maxResults {
			messageString = messageString + fmt.Sprintf("\n\n%s\n+ %d / %d players connected\n+ Playing %s on %s", result.Name, result.Players, result.MaxPlayers, result.Playlist, result.Map)
			if result.HasPassword {
				messageString = messageString + "\n- PASSWORD PROTECTED"
			}
		}
	}

	messageString = messageString + "\n```"

	return messageString
}

func SearchHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: createSearchMessage(i),
		},
	})
}
