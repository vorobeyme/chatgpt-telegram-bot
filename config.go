package bot

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config represents structure that hold the YAML data
type Config struct {
	Debug    bool           `yaml:"debug"`
	Telegram TelegramConfig `yaml:"telegram"`
	ChatGPT  ChatGPTConfig  `yaml:"chatGPT"`
}

// TelegramConfig represents telegram configurations.
type TelegramConfig struct {
	Token          string `yaml:"token"`
	ReplyToMessage bool   `yaml:"reply_to_message"`
}

// ChatGPTConfig represents ChatGPT configurations.
type ChatGPTConfig struct {
	APIKey           string  `yaml:"apiKey"`
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

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error reading .env file")
	}

	data = []byte(os.ExpandEnv(string(data)))

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("error parsing .yaml file")
	}

	return nil
}
expand env variables in yaml file