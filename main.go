package main

import (
	"fmt"
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

	var token = os.Getenv("INPUT_TOKEN")

	if token == "" {
		log.Fatal("token input is required")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err.Error())
	}

	msg, err := newMessage()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func newMessage() (*tgbotapi.MessageConfig, error) {
	var (
		chat      = os.Getenv("INPUT_CHAT_ID")
		message   = os.Getenv("INPUT_MESSAGE")
		parseMode = os.Getenv("INPUT_PARSE_MODE")
	)

	if chat == "" {
		return nil, fmt.Errorf("chat_id is required")
	}

	chatID, err := strconv.ParseInt(chat, 10, 64)
	if err != nil {
		return nil, err
	}

	if message == "" {
		message = "Pushed" //TODO:
	}
	var msg = tgbotapi.NewMessage(chatID, message)
	switch parseMode {
	case "markdown":
		msg.ParseMode = "Markdown"
	case "markdown2":
		msg.ParseMode = "MarkdownV2"
	case "html":
		msg.ParseMode = "HTML"
	}

	return &msg, nil
}
