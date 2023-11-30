package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddInlineKeyboards(bot *tgbotapi.BotAPI, chatID int64, text, replyMsg string) {
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Прогноз погоды","button 1"),
				tgbotapi.NewInlineKeyboardButtonData("Мем дня","button 2"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Блокнуть Олега на 1 минуту", "Button 3"),
			),
		)
		msg := tgbotapi.NewMessage(chatID, "Дратути, выбирайте: ")
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)
		
	}
