package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

/* Initialization */

// 1. Get configuration.

var (
	confFile string // the path of configuration file
)

func init() {
	flag.StringVar(&confFile, "c", "./bot.json", "the path of configuratioin file")
	flag.Parse()
}

type config struct {
	AppId string
	PublicKey string
	ClientId string
	BotToken string
	GuildId string
}

var conf config

func init() {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &conf); err != nil {
		log.Fatal(err)
	}
}

// 2. Instantialize bot.

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot token: %v", err)
	}
}

/* Commands */

var (
	commands = []*discordgo.ApplicationCommand {
		{
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
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"ido": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		},
	}
)


/* Event Handlers */

func init() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("Prepared!") })
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	_, err = s.ApplicationCommandBulkOverwrite(s.State.User.ID, conf.GuildId, commands)
	if err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")
}
