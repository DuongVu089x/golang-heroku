package main

import (
	"github.com/DuongVu089x/golang-heroku/action"
	"github.com/DuongVu089x/golang-heroku/config"
	"github.com/labstack/echo"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("env")
	config.Init()

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// telegram
	initTelegram()

	if config.Bot == nil {
		return
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/"  + config.Bot.Token, action.WebhookHandler)
	e.Logger.Fatal(e.Start(":"+port))
}

func initTelegram() {
	var err error

	config.Bot, err = tgbotapi.NewBotAPI(config.Config.Key["bot-token"])
	if err != nil {
		log.Println(err)
		return
	}

	url := config.Config.OutboundURL["base-url"] + config.Bot.Token
	_, err = config.Bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

