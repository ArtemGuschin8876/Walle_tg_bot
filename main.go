package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TokenBot)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			text := update.Message.Text      
			chatID := update.Message.Chat.ID 
			userID := update.Message.From.ID 
			var replyMsg string

			log.Printf("[%s](%d) %s", update.Message.From.UserName, userID, text)

			msg := tgbotapi.NewMessage(chatID, replyMsg) 
			msg.ReplyToMessageID = update.Message.MessageID 

			bot.Send(msg)
		}
	}
}