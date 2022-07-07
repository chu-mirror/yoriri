package cb

import (
	"github.com/chu-mirror/yoriri/activity"
)

func init() {
	RegisterCommand(
		*discordgo.ApplicationCommand {
			Name: "ido",
			Description: "Register your hit",
			Type: discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "boss",
					Description: "Which boss to hit",
					Type: discordgo.ApplicationCommandOptionString,
					Required: true,
					Autocomplete: true,
				},
				{
					Name: "damage",
					Description: "how much you can do",
					Type: discordgo.ApplicationCommandOptionString,
					Required: true,
					Autocomplete: false,
				},
			},
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				data := i.ApplicationCommandData()
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData {
						Content: fmt.Sprintf(
							"You registered to hit %q",
							data.Options[0].StringValue(),
						),
					},
				})
				if err !=  nil {
					panic(err)
				}
			case discordgo.InteractionApplicationCommandAutocomplete:
				choices := []*discordgo.ApplicationCommandOptionChoice {
					{
						Name: "B1",
						Value: "1",
					},
					{
						Name: "B2",
						Value: "2",
					},
					{
						Name: "B3",
						Value: "3",
					},
				}
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
					Type: discordgo.InteractionApplicationCommandAutocompleteResult,
					Data: &discordgo.InteractionResponseData {
						Choices: choices,
					},
				})
				if err != nil {
					panic(err)
				}
			}
		}
	)
}
