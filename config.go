package bot

type Config struct {
	Debug bool `yaml:"debug" env:"DEBUG"`

	Telegram TelegramConfig `yaml:"telegram"`
	ChatGPT  ChatGPTConfig  `yaml:"chatGPT"`
}

type TelegramConfig struct {
	Token string `env:"TELEGRAM_TOKEN"`

	ReplyToMessage bool `yaml:"reply_to_message"`
}

type ChatGPTConfig struct {
	Token string `env:"CHATGPT_TOKEN"`

	Model            string  `yaml:"model"`
	MaxTokens        int     `yaml:"max_tokens"`
	Temperature      float64 `yaml:"temperature"`
	TopP             int     `yaml:"top_p"`
	FrequencyPenalty float64 `yaml:"frequency_penalty"`
	PresencePenalty  float64 `yaml:"presence_penalty"`
}

func NewConfig() (*Config, error) {
	var config *Config

	return config, nil
}
