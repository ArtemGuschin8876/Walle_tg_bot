package callback

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleCallBack(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery){
	switch callbackQuery.Data {
	case "Блокнуть Олега": 
		handleButton1(bot, callbackQuery)
	default:
		handleUnknownButton(bot, callbackQuery)
	}
}

func handleButton1(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery){
	replyMsg := "Ты нажал на Олега"
	msg := tgbotapi.NewMessage(callbackQuery.From.ID, replyMsg)
	bot.Send(msg)
}

func handleUnknownButton(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
    replyMsg := "Неизвестная кнопка!"
    msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, replyMsg)
    bot.Send(msg)
}

