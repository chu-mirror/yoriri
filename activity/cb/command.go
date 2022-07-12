package cb

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chu-mirror/yoriri/activity"
	"github.com/chu-mirror/yoriri/activity/cb/hit"
	"github.com/chu-mirror/yoriri/activity/cb/state"
)

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: msg},
	})
}

func init() {
	activity.RegisterCommand(
		&discordgo.ApplicationCommand{
			Name:        "hit",
			Description: "Hit a boss",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "boss",
					Description:  "which boss to hit, input the number directly(1~5)",
					Type:         discordgo.ApplicationCommandOptionInteger,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "sync",
					Description:  "whether to sync",
					Type:         discordgo.ApplicationCommandOptionBoolean,
					Required:     false,
					Autocomplete: false,
				},
			},
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				ops := i.ApplicationCommandData().Options
				boss := ops[0].IntValue()
				tosync := false
				if len(ops) > 1 {
					tosync = ops[1].BoolValue()
				}
				if boss < 1 || boss > 5 {
					return respond(s, i, "Invalid boss number")
				}
				errNo := hit.Hit(i.Interaction.Member.User.ID, state.IntToBossNo(boss), tosync)
				switch errNo {
				case hit.HitLockedFail:
					return respond(s, i, "The boss is hitting by another people")
				case hit.HitInHittingFail:
					return respond(s, i, "You are already hitting another boss")
				}
				return respond(s, i, "Go Go Go")
			}
			return nil
		},
	)
}
