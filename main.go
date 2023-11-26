package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TokenBot")
	if token == "" {
        log.Fatal("Токен бота не найден в переменной окружения")
    }

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println("Ошибка при создании бота", err)
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