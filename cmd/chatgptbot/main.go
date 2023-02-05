package main

import (
	"log"

	bot "github.com/vorobeyme/chatgpt-telegram-bot"
)

func main() {
	config, err := bot.NewConfig()
	if err != nil {
		log.Fatalf("[ERROR] configuration file: %v", err)
	}

	bot, err := bot.NewChatGPTBot(config)
	if err != nil {
		log.Fatalf("[ERROR] telegram API - an authentication error occurs: %v", err)
	}
	defer bot.Close()

	log.Fatal(bot.Run())
}
