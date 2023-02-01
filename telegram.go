// Package bot has functions and types used for interacting withthe Telegram Bot API
package bot

// HTTPClient is the type needed for the bot to perform HTTP requests.
type HTTPClient interface {
	Do()
}

// BotAPI allows you to interact with the Telegram Bot API.
type BotAPI struct {
	Token  string
	Client HTTPClient

	apiEndpoint string
}

// NewBotAPI creates a new BotAPI instance.
func NewBotAPI(token, apiEndpoint string, client HTTPClient) *BotAPI {
	return &BotAPI{
		Token:       token,
		Client:      client,
		apiEndpoint: apiEndpoint,
	}
}

// SetAPIEndpoint changes the Telegram Bot API endpoint used by the instance.
func (b *BotAPI) SetAPIEndpoint(apiEndpoint string) {
	b.apiEndpoint = apiEndpoint
}

// @see https://github.com/go-telegram-bot-api/telegram-bot-api
