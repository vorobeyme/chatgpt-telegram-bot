package bot

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type chatGPTService struct {
	gpt3   *gogpt.Client
	config *ChatGPTConfig
}

type chatgptOptions struct {
	model            string
	maxTokens        int
	temperature      float64
	topP             int
	frequencyPenalty float64
	presencePenalty  float64
}

// NewChatGPTService returns a new service instance.
func NewChatGPTService(cfg *ChatGPTConfig) *chatGPTService {
	return &chatGPTService{
		gpt3:   gogpt.NewClient(cfg.APIKey),
		config: cfg,
	}
}

func (c *chatGPTService) Ask(prompt string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 100,
		Prompt:    prompt,
	}
	res, err := c.gpt3.CreateCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	fmt.Println("GPT3 completion choices:", res.Choices)

	return res.Choices[0].Text, nil
}
