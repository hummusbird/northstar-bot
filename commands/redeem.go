package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RedeemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(201) + 50

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("```Diff\n+%d ScorchBucks\n```", c),
		},
	})

}
