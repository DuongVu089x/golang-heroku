package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

const (
	// TOKEN telegram
	TOKEN = "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs"
	// URL telegram
	URL = "https://api.telegram.org/bot"
)

var (
	bot      *tgbotapi.BotAPI
	botToken = "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs"
	baseURL  = "https://desolate-falls-71497.herokuapp.com/"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// telegram
	initTelegram()

	e := echo.New()
	e.POST("/"  + bot.Token, webhookHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}

func initTelegram() {
	var err error

	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Println(err)
		return
	}

	// this perhaps should be conditional on GetWebhookInfo()
	// only set webhook if it is not set properly
	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
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

	var update tgbotapi.Update
	err := json.Unmarshal([]byte(body), &update)
	if err != nil {
		log.Println(err)
		return err
	}

	// to monitor changes run: heroku logs --tail
	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
	return nil
}