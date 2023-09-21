package commands

import (
	"fmt"
	"northstar-bot/util"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func createStatusString() string {
	var (
		data           []util.ServerInfo
		currentPlayers = 0
		maxPlayers     = 0
		serverCount    int
		protectedCount = 0
		population     int
		wg             sync.WaitGroup
	)

	wg.Add(1)

	go func() {
		defer wg.Done()
		data = util.ProcessAPIResp()
		serverCount = len(data)
		for _, server := range data {
			currentPlayers += server.Players
			maxPlayers += server.MaxPlayers
			if server.HasPassword {
				protectedCount++
			}
		}

		population = (currentPlayers * 100) / maxPlayers
	}()

	wg.Wait()

	messageString := fmt.Sprintf("```diff\n## NORTHSTAR.TF STATUS: ##\n\n+ Servers Online: %d\n\n- Password Protected Servers: %d\n\n+ Players in-game: %d/%d (%d%%)\n```", serverCount, protectedCount, currentPlayers, maxPlayers, population)

	return messageString
}

func StatusHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: createStatusString(),
		},
	})

}
