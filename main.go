package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var icons = map[string]string{
	"success":   "✅",
	"failure":   "❌",
	"cancelled": "❕",
}

func main() {

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
		chat                = os.Getenv("INPUT_CHAT_ID")
		message             = os.Getenv("INPUT_MESSAGE")
		parseMode           = os.Getenv("INPUT_PARSE_MODE")
		disableLinksPreview = os.Getenv("INPUT_DISABLE_LINKS_PREVIEW")
		status              = os.Getenv("INPUT_STATUS")
	)

	log.Printf("%#v", os.Environ())

	if chat == "" {
		return nil, fmt.Errorf("chat_id is required")
	}

	chatID, err := strconv.ParseInt(chat, 10, 64)
	if err != nil {
		return nil, err
	}

	var msg = tgbotapi.NewMessage(chatID, fmt.Sprintf("%s %s", icons[status], message))
	switch parseMode {
	case "markdown":
		msg.ParseMode = "Markdown"
	case "markdown2":
		msg.ParseMode = "MarkdownV2"
	case "html":
		msg.ParseMode = "HTML"
	}

	if strings.ToLower(disableLinksPreview) == "true" {
		msg.DisableWebPagePreview = true
	}

	return &msg, nil
}
