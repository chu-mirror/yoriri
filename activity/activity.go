// Package activity defines the behaviour of bot, basically a simplified version
// of a typical discord bot with limited types of interaction.
package activity

import (
	"os"
	"log"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	commands []*discordgo.ApplicationCommand
	commandHandlers = make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate)error)

	initializations []func(*discordgo.Session)error
)

func RegisterCommand(cmd *discordgo.ApplicationCommand, handler func(*discordgo.Session, *discordgo.InteractionCreate)error) {
	commands = append(commands, cmd)
	commandHandlers[cmd.Name] = handler
}

func RegisterInitialization(it func(*discordgo.Session)error) {
	initializations = append(initializations, it)
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
			err := h(s, i)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	err = l.session.Open()
	if err != nil {
		return
	}


	for _, it := range initializations {
		err = it(l.session)
		if err != nil {
			return
		}
	}

	registered, err := l.session.ApplicationCommandBulkOverwrite(l.session.State.User.ID, os.Getenv("GUILDID"), commands)
	if err != nil {
		return
	}

	if len(registered) != len(commands) {
		return fmt.Errorf("didn't register all commands successfully")
	}

	return
}

func (l Life) End() {
	l.session.Close()
}
