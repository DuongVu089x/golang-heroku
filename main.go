package main

import (
	"encoding/json"
	"github.com/DuongVu089x/golang-heroku/config"
	"github.com/labstack/echo"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	port := "3000"
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
	e.POST("/"  + config.Bot.Token, webhookHandler)
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

func webhookHandler(c echo.Context) error {

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	body := string(bodyBytes)
	log.Println(body)

	var update tgbotapi.Update
	err := json.Unmarshal([]byte(body), &update)
	if err != nil {
		log.Println(err)
		return err
	}

	handlerMessage(&update)

	// to monitor changes run: heroku logs --tail
	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
	return nil
}


func handlerMessage(update *tgbotapi.Update) {
	message := update.Message.Text

	if message == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command is required")
		config.Bot.Send(msg)
		return
	}
	message =  strings.ToLower(message)

	var replyMessage string

	switch message {
	case "/help":
		// Show all command
		replyMessage = showAllCommand()

	case "count history":
		// Call api count history
		//executeCountData(model.DBHistoryQueue, &replyMessage)
		replyMessage = "0"
	default:
		replyMessage = "Command isn't defined"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyMessage)
	msg.ReplyToMessageID = update.Message.MessageID

	config.Bot.Send(msg)
	return
}

func showAllCommand()string{
	return `
			- count {db}
				+ history
				+ oms
				+ oos
				+ order-tracking
				+ pptl-history
				+ transport-package
				+ update-warehouse
		`
}

