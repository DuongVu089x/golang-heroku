package config

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"os"
)

type config struct {
	OutboundURL map[string]string
	Key         map[string]string
	Number      map[string]int64
}

// Config main config object
var Config config
var Bot *tgbotapi.BotAPI

// Init config
func Init() error {
	env := os.Getenv("env")

	switch env {
	case "dev":
		Config = config{
			OutboundURL: map[string]string{
				"base-url": "https://desolate-falls-71497.herokuapp.com/",
			},

			Key: map[string]string{
				"bot-token": "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs",
			},
		}
	}
	return nil
}
