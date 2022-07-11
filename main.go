package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/chu-mirror/yoriri/activity"
	_ "github.com/chu-mirror/yoriri/activity/cb"
)

func main() {
	bot, err := activity.Birth(os.Getenv("BOTTOKEN"))
	if err != nil {
		log.Fatalf("Cannot give bot birth: %v", err)
	}
	err = bot.Start()
	if err != nil {
		log.Fatalf("Cannot start the bot: %v", err)
	}
	defer bot.End()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")
}
