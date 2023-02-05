[![Go](https://github.com/vorobeyme/chatgpt-telegram-bot/actions/workflows/ci.yml/badge.svg)](https://github.com/vorobeyme/chatgpt-telegram-bot/actions/workflows/ci.yml)

# :robot: GPT Telegram Bot
Simple telegram bot written in Go that uses the OpenAI completions API to have conversations with the GPT-3 (Generative Pretrained Transformer 3) model.

## Run
- Copy `.env.dist` to `.env`
- Fill `TELEGRAM_TOKEN` and `OPENAI_API_KEY` with your values
- Run the bot `go run cmd/gptbot/main.go`
