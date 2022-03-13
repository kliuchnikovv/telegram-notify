package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// icons := map[string]string{
	// 	"failure":   "❗",
	// 	"cancelled": "❕",
	// 	"success":   "✅",
	// }

	var (
		token   = os.Getenv("TELEGRAM_TOKEN")
		chat    = os.Getenv("TELEGRAM_CHAT")
		message = os.Getenv("TELEGRAM_MESSAGE")

		chatID int64
		err    error
	)

	if token == "" {
		log.Fatal("token input is required")
	}

	if chat == "" {
		log.Fatal("chat input is required")
	} else {
		chatID, err = strconv.ParseInt(chat, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if message == "" {
		message = "Pushed" //TODO:
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = bot.Send(tgbotapi.NewMessage(chatID, message))
	if err != nil {
		log.Fatal(err.Error())
	}
}
