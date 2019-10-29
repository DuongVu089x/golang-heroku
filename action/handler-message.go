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
	"time"
)

// WebhookHandler...
func WebhookHandler(c echo.Context) error {

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	body := string(bodyBytes)

	var update tgbotapi.Update
	err := json.Unmarshal([]byte(body), &update)
	if err != nil {
		log.Println(err)
		return err
	}

	handlerMessage(&update)
	return nil
}

// handlerMessage...
func handlerMessage(update *tgbotapi.Update) {
	if update == nil || update.Message == nil {
		return
	}

	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
	message := update.Message.Text

	if message == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command is required")
		config.Bot.Send(msg)
		return
	}
	messageArr := strings.Split(message, " ")
	messageArr[0] = strings.ToLower(messageArr[0])

	// Handler set token
	var replyMessage string
	switch messageArr[0] {
	case "/start":
		replyMessage = "Type /help to more info"
	case "/help":
		// Show all command
		replyMessage = showAllCommand()
	case "/count":
		replyMessage = handlerCount(messageArr[1])
	case "/set-auto":
		replyMessage = handlerSetAuto(update.Message.Chat.ID)
	case "/quit-auto":
		replyMessage = handlerQuitAuto(update.Message.Chat.ID)
		return
	default:
		replyMessage = "Command isn't defined"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyMessage)
	msg.ReplyToMessageID = update.Message.MessageID

	config.Bot.Send(msg)
	return
}

// showAllCommand...
func showAllCommand() string{
	return `
			- /count {db}
				+ history
				+ oms
				+ oos
				+ order-tracking
				+ pptl-history
				+ transport-package
				+ update-warehouse
			- set-auto: Auto count record in DB queue
			- quit-auto: Quit auto count record in DB queue
		`
}

// handlerCount...
func handlerCount(tableName string) string {
	// Call api count history
	req, err := http.NewRequest("GET", config.Config.OutboundURL["pmq-count"] + tableName, nil)
	if err != nil {
		return "Some error!"
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic " + config.Config.Key["basic-token"] )

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

// handlerSetAuto...
func handlerSetAuto(id int64) string {

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	//chanel = &quit

	(*config.UserChanel)[id] = &quit
	go func() {
		for {
			select {
			case <- ticker.C:
				msg := tgbotapi.NewMessage(616257809, time.Now().String())
				config.Bot.Send(msg)
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
	return "Success"
}

// handlerQuitAuto...
func handlerQuitAuto(id int64) string {
	chanel := (*config.UserChanel)[id]
	if chanel != nil {
		close(*(chanel))
		return "Success"
	}
	return "Error"
}