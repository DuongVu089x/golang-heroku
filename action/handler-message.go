package action

import (
	"encoding/json"
	"github.com/DuongVu089x/golang-heroku/config"
	"github.com/labstack/echo"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func WebhookHandler(c echo.Context) error {

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
	//var resp *http.Response
	switch message {
	case "/start":
		replyMessage = "Type /help to more info"
	case "/set-token":
		// Handler set token
		handlerSetToken();
	case "/help":
		// Show all command
		replyMessage = showAllCommand()

	case "count history":
		// Call api count history
		resp, err := http.Get("http://35.247.150.56/pmq/v1/count?tableName=history")
		if err != nil {
			replyMessage = "Something wrong!"
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil{
			replyMessage = "Something wrong!"
		}

		replyMessage = string(body)
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

func handlerSetToken() {

}
