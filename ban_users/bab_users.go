package banusers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var err error

func DeleteMessagesUser(bot *tgbotapi.BotAPI, chatID int64, messageID int, update *tgbotapi.Update){
		deleteMessage := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
		_, err := bot.Send(deleteMessage)
		if err != nil {
			log.Printf("Error deleting message: %v", err)
		}
}

func SendOtherMessage(bot *tgbotapi.BotAPI, chatID int64,update *tgbotapi.Update, text string){
	newMessage := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		_, err = bot.Send(newMessage)
		if err != nil {
			log.Printf("Error sending new message: %v", err)
		}
}