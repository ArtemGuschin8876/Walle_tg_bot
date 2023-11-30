package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func AddStandartKeyboard(bot *tgbotapi.BotAPI, chatID int64){	
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
		 	tgbotapi.NewKeyboardButton("Блокнуть Олега"),
			tgbotapi.NewKeyboardButton("Вторая"),
			tgbotapi.NewKeyboardButton("Третья"),
		),	
	)

		replyKeyboard.OneTimeKeyboard = true
		replyKeyboard.ResizeKeyboard = true
	msg := tgbotapi.NewMessage(chatID, "Пока варианты допиливаются")

	msg.ReplyMarkup = replyKeyboard
	bot.Send(msg)
}