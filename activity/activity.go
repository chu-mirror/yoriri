package activity

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	commands []*discordgo.ApplicationCommand
	commandHandlers = make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
)

func RegisterCommand(cmd *discordgo.ApplicationCommand, handler func(*discordgo.Session, *discordgo.InteractionCreate)) {
	commands = append(commands, cmd)
	commandHandlers[cmd.Name] = handler
}

type Life struct {
	session *discordgo.Session
}

func Birth(token string) (l Life, err error) {
	l.session, err = discordgo.New("Bot " + token)
	return
}

func (l Life) Start() (err error){
	l.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err = l.session.Open()
	if err != nil {
		return
	}

	_, err = l.session.ApplicationCommandBulkOverwrite(l.session.State.User.ID, os.Getenv("GUILDID"), commands)
	return
}

func (l Life) End() {
	l.session.Close()
}
