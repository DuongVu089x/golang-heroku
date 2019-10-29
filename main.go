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

var bot *tgbotapi.BotAPI

func main() {

	port := os.Getenv("PORT")
	config.Init()

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// telegram
	initTelegram()

	if bot == nil {
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/"  + bot.Token, action.WebhookHandler)
	e.Logger.Fatal(e.Start(":"+port))
}

func initTelegram() {
	var err error

	bot, err = tgbotapi.NewBotAPI(config.Config.Key["bot-token"])
	if err != nil {
		log.Println(err)
		return
	}

	url := config.Config.OutboundURL["base-url"] + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
	config.Bot = bot
}

