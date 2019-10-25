package main

import (
	"encoding/json"
	"fmt"
	"github.com/DuongVu089x/golang-heroku/model"
	"log"
	"net/http"
	"os"

	//"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

const (
	// TOKEN telegram
	TOKEN = "YOUR-TOKEN"
	// URL telegram
	URL = "https://api.telegram.org/bot"
	// PORT local
)

func update(w http.ResponseWriter, r *http.Request) {

	message := &model.ReceiveMessage{}

	chatID := 0
	msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
	}

	// if private or group
	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	} else {
		// if channel
		fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
		chatID = message.ChannelPost.Chat.ID
		msgText = message.ChannelPost.Text
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", URL, TOKEN, chatID, msgText)

	// send echo resp
	_, err = http.Get(respMsg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	port :=  os.Getenv("PORT")

	http.HandleFunc("/api/v1/update", update)

	fmt.Println("Listenning on port", port, ".")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}