package cb

import (
	"fmt"
	"strconv"
	
	"github.com/bwmarrin/discordgo"
	"github.com/chu-mirror/yoriri/activity"
	"github.com/chu-mirror/yoriri/activity/cb/boss"
	"github.com/chu-mirror/yoriri/activity/cb/member"
)

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: msg},
	})
}

/* Command: hit boss [sync]

   Declare hitting a boss, the optional option said whether this hit is to sync or not.
 */

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
				target := ops[0].IntValue()
				tosync := false
				if len(ops) > 1 {
					tosync = ops[1].BoolValue()
				}
				if target < 1 || target > 5 {
					return respond(s, i, "Invalid boss number")
				}
				errNo := member.Hit(i.Interaction.Member.User.ID, boss.IntToNo(target), tosync)
				switch errNo {
				case member.HitLockedFail:
					return respond(s, i, "The boss is hitting by another people")
				case member.HitInHittingFail:
					return respond(s, i, "You are already hitting another boss")
				}
				return respond(s, i, fmt.Sprintf("Ok, go to hit B%d", target))
			}
			return nil
		},
	)
}

/* Command: done dmg

   Declare how much have done.
   The format of dmg, number + M/K.
 */

func damageNumber(dmg string) (int, bool) {
	d, e := strconv.ParseFloat(dmg[:len(dmg)-1], 32)
	if e != nil {
		return 0, false
	}
	var n int
	switch dmg[len(dmg)-1] {
	case 'M':
		n = int(d*1000)
	case 'K':
		n = int(d)
	default:
		return 0, false
	}
	return n, true
}	

func init() {
	activity.RegisterCommand(
		&discordgo.ApplicationCommand{
			Name:        "done",
			Description: "Record the damage you have done",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "damage",
					Description:  "The format, for example, 22.1M, 50K",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
			},
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				ops := i.ApplicationCommandData().Options
				dmg := ops[0].StringValue()
				d, ok := damageNumber(dmg)
				if !ok {
					return respond(s, i, "Unable to parse input")
				}
				errNo := member.Done(i.Interaction.Member.User.ID, d)
				switch errNo {
				case member.HitNoHittingFail:
					return respond(s, i, "You are not hitting now")
				case member.HitIllegalHpFail:
					return respond(s, i, "Please check your input number")
				}
				return respond(s, i, "Ok, good job")
			}
			return nil
		},
	)
}
