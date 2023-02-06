package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiBaseURL = "https://api.openai.com/v1"

var (
	ctx = context.Background()

	completionsEndpoint = "completions"
)

type completionRequest struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt,omitempty"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float64        `json:"temperature,omitempty"`
	TopP             float64        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	LogProbs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float64        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64        `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

type errorResponse struct {
	Error *struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Param   string `json:"param,omitempty"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type completionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
	LogProbs     struct {
		Tokens        []string             `json:"tokens"`
		TokenLogprobs []float32            `json:"token_logprobs"`
		TopLogprobs   []map[string]float32 `json:"top_logprobs"`
		TextOffset    []int                `json:"text_offset"`
	} `json:"logprobs"`
}

type GPT3Service interface {
	Ask() (string, error)
}

type gpt3Service struct {
	baseURL    string
	config     *ChatGPTConfig
	httpClient *http.Client
}

func NewGPT3Service(cfg *ChatGPTConfig) *gpt3Service {
	return &gpt3Service{
		baseURL:    apiBaseURL,
		config:     cfg,
		httpClient: &http.Client{},
	}
}

func (s *gpt3Service) Ask(prompt string) (string, error) {
	res, err := s.createCompletion(completionRequest{
		Prompt:           prompt,
		Model:            s.config.Model,
		MaxTokens:        s.config.MaxTokens,
		Temperature:      float64(s.config.Temperature),
		TopP:             float64(s.config.TopP),
		PresencePenalty:  s.config.PresencePenalty,
		FrequencyPenalty: s.config.FrequencyPenalty,
	})
	if err != nil {
		return "", err
	}
	return res.Choices[0].Text, nil
}

func (s *gpt3Service) createCompletion(req completionRequest) (*completionResponse, error) {
	var reqBytes []byte
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return &completionResponse{}, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", s.baseURL, completionsEndpoint), bytes.NewBuffer(reqBytes))
	if err != nil {
		return &completionResponse{}, err
	}

	request = request.WithContext(ctx)
	response := &completionResponse{}
	if err := s.makeRequest(request, response); err != nil {
		return &completionResponse{}, err
	}

	return response, nil
}

func (s *gpt3Service) makeRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.config.APIKey))

	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			return fmt.Errorf("error, status code: %d", res.StatusCode)
		}
		return fmt.Errorf("error, status code: %d, message: %s", res.StatusCode, errRes.Error.Message)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return err
		}
	}

	return nil
}
