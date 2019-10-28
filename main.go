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

//func webhookHandler(c echo.Context) error {
//
//	var bodyBytes []byte
//	if c.Request().Body != nil {
//		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
//	}
//
//	body := string(bodyBytes)
//	log.Println(body)
//
//	var update tgbotapi.Update
//	err := json.Unmarshal([]byte(body), &update)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	handlerMessage(&update)
//
//	// to monitor changes run: heroku logs --tail
//	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
//	return nil
//}
//
//
//func handlerMessage(update *tgbotapi.Update) {
//	message := update.Message.Text
//
//	if message == "" {
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command is required")
//		bot.Send(msg)
//		return
//	}
//	message =  strings.ToLower(message)
//
//	var replyMessage string
//
//	switch message {
//	case "/help":
//		// Show all command
//		replyMessage = showAllCommand()
//
//	case "count history":
//		// Call api count history
//		//executeCountData(model.DBHistoryQueue, &replyMessage)
//		replyMessage = "0"
//	default:
//		replyMessage = "Command isn't defined"
//	}
//
//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyMessage)
//	msg.ReplyToMessageID = update.Message.MessageID
//
//	bot.Send(msg)
//	return
//}
//
//func showAllCommand()string{
//	return `
//			- count {db}
//				+ history
//				+ oms
//				+ oos
//				+ order-tracking
//				+ pptl-history
//				+ transport-package
//				+ update-warehouse
//		`
//}

