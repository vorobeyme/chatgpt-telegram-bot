package main

import (
	"log"

	bot "github.com/vorobeyme/chatgpt-telegram-bot"
)

func main() {
	var config = &bot.Config{}
	config.Telegram.Token = "tgToken"
	config.ChatGPT.Token = "chatgptToken"

	bot, err := bot.NewChatGPTBot(config)
	if err != nil {
		log.Fatal(err)
	}
	defer bot.Close()

	log.Fatal(bot.Run())
}
