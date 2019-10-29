package action

import (
	"encoding/json"
	"fmt"
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
	if update == nil || update.Message == nil {
		return
	}

	message := update.Message.Text

	if message == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command is required")
		config.Bot.Send(msg)
		return
	}
	message =  strings.ToLower(message)
	messageArr := strings.Split(message, " ")

	// Handler set token
	var replyMessage string
	switch messageArr[0] {
	case "/start":
		replyMessage = "Type /help to more info"
	case "/help":
		// Show all command
		replyMessage = showAllCommand()
		return
	case "/set-token":
		fmt.Println(messageArr)
		fmt.Println(messageArr[1])
		handlerSetToken(update.Message.Chat.ID, messageArr[1])
		replyMessage = "Set token success"
	case "count":
		replyMessage = handlerCount(update.Message.Chat.ID)
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

func handlerSetToken(id int64, token string) {
	m := *config.UserToken
	m[id] = token
	fmt.Println("Token: " + token)
}

func handlerCount(id int64) string {
	var token string
	m := *config.UserToken
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
		if key == id {
			token = value
		}
	}

	//if token == "" {
	//	return "Token is required"
	//}

	// Call api count history
	req, err := http.NewRequest("GET","http://35.247.150.56/pmq/v1/count?tableName=history", nil)
	if err != nil {
		return "Something wrong!"
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)
	fmt.Println("req: " + req.Header.Get("Authorization"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
