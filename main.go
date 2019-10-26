package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"os"

	//"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

const (
	// TOKEN telegram
	TOKEN = "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs"
	// URL telegram
	URL = "https://api.telegram.org/bot"
	// PORT local
	// PORT = os.Getenv("PORT")
)

var (
	bot      *tgbotapi.BotAPI
	botToken = "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs"
	baseURL  = "https://<YOUR-APP-NAME>.herokuapp.com/"
)

func main() {
	//port := os.Getenv("PORT")
	//
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}
	//
	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//
	//e.POST("/demo", func(c echo.Context) error {
	//	log.Print(c)
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":"+port))

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// gin router
	router := gin.New()
	router.Use(gin.Logger())

	// telegram
	initTelegram()
	router.POST("/" + bot.Token, webhookHandler)

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
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

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	// to monitor changes run: heroku logs --tail
	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
}