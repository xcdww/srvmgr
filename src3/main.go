package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TOKEN = "5803513827:AAE7jEV8JkUiDCFY-OXH-MuzmbcPF_dFi7Y"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	check(err)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			comm := update.Message.Command()
			log.Printf("com: %s", comm)
			switch comm {
			case "help":
				msg.Text = "type /sayhi or /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}

	}
}

func rCommand(command string) {
	co := strings.Split(command, "_")
	fmt.Println(co)
}
