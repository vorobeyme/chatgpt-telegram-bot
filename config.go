package bot

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config represents structure that hold the YAML data
type Config struct {
	Debug    bool           `yaml:"debug" env:"APP_DEBUG"`
	Telegram TelegramConfig `yaml:"telegram"`
	ChatGPT  ChatGPTConfig  `yaml:"chatGPT"`
}

// TelegramConfig represents telegram configurations.
type TelegramConfig struct {
	Token          string
	ReplyToMessage bool `yaml:"reply_to_message"`
}

// ChatGPTConfig represents ChatGPT configurations.
type ChatGPTConfig struct {
	APIKey           string
	Model            string  `yaml:"model"`
	MaxTokens        int     `yaml:"maxTokens"`
	Temperature      float64 `yaml:"temperature"`
	TopP             int     `yaml:"topP"`
	FrequencyPenalty float64 `yaml:"frequencyPenalty"`
	PresencePenalty  float64 `yaml:"presencePenalty"`
}

func NewConfig() (*Config, error) {
	var config = &Config{}
	if err := parseConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func parseConfig(cfg *Config) error {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return fmt.Errorf("error reading .yaml file")
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return fmt.Errorf("error parsing .yaml file")
	}

	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("error reading .env file")
	}

	cfg.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	cfg.Telegram.Token = os.Getenv("TELEGRAM_TOKEN")
	cfg.ChatGPT.APIKey = os.Getenv("OPENAI_API_KEY")

	return nil
}
