package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Config struct for toml config file
type Config struct {
}

var config Config

func init() {
	// Read config file
}

func main() {

	// Exiting
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
	log.Println("Shutdown signal received, exiting...")
}

// @see https://github.com/mxssl/tg-captcha-bot/
