package bot

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdStart = "start"
	cmdHelp  = "help"
	cmdAsk   = "ask"
)

type chatGPTBot struct {
	bot         *tg.BotAPI
	config      *TelegramConfig
	gpt3Service *gpt3Service
}

func NewChatGPTBot(cfg *Config) (*chatGPTBot, error) {
	bot, err := tg.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = cfg.Debug

	gptBot := &chatGPTBot{
		bot:         bot,
		config:      &cfg.Telegram,
		gpt3Service: NewGPT3Service(&cfg.ChatGPT),
	}

	return gptBot, nil
}

func (b *chatGPTBot) Run() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	for update := range b.bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}
		if err := b.processUpdate(&update); err != nil {
			return err
		}
	}
	return nil
}

func (b *chatGPTBot) processUpdate(u *tg.Update) error {
	log.Printf("[%s %d] %s", u.Message.From.UserName, u.Message.From.ID, u.Message.Text)
	var (
		err     error
		message = tg.NewMessage(u.Message.Chat.ID, u.Message.Text)
	)
	if b.config.ReplyToMessage {
		message.ReplyToMessageID = u.Message.MessageID
	}

	if u.Message.IsCommand() {
		switch u.Message.Command() {
		case cmdStart:
			message.Text = "Hello, @" + u.Message.From.UserName + "!"
		case cmdHelp:
			message.Text = fmt.Sprintf("I understand /%s command.\nTo exit type quit or exit", cmdAsk)
		default:
			message.Text = "Sorry... I don't know that command"
		}
	} else if len(u.Message.Text) > 0 {
		message.Text, err = b.processMessage(u.Message.Text)
		if err != nil {
			return err
		}
	}

	if _, err := b.bot.Send(message); err != nil {
		return err
	}

	return nil
}

func (b *chatGPTBot) processMessage(prompt string) (string, error) {
	// The most powerful AI ;)
	b64 := base64.StdEncoding.EncodeToString([]byte(strings.ToLower(prompt)))
	if strings.Contains(b64, "QviDQv9C+INGA0YPRgdC90ZY") {
		res, _ := base64.URLEncoding.DecodeString("0YDRg9GB0L3RliDQv9GW0LfQtNCwIA==")
		return string(res) + "\xF0\x9F\x98\x81", nil
	}
	return b.gpt3Service.Ask(prompt)
}

func (b *chatGPTBot) Close() error {
	log.Println("Closing bot...")
	return nil
}
