package main

import (
	"fmt"
	"log"
	"os"
	callback "projects/Walley_tg_bot/call_back"
	"projects/Walley_tg_bot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TokenBot")
	if token == "" {
		log.Fatal("Bot token not found in environment variable")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Error when creating a bot")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//Upd
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Logger
	logFile, err := os.Create("chat_logs.txt")
	if err != nil {
		log.Fatal("Creation error logger")
	}
	defer logFile.Close()
	
	fmt.Println("Logger working...")
	logger := log.New(logFile, "", log.LstdFlags)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback.HandleCallBack(bot, update.CallbackQuery)
	
		}

		if update.Message != nil {
			text := update.Message.Text
			chatID := update.Message.Chat.ID
			userID := update.Message.From.ID
			
			var replyMsg string

			log.Printf("[%s](%d) %s", update.Message.From.UserName, userID, text)

			switch text {
			case "/start", "start", "Start", "Запуск", "1":	
				replyMsg = fmt.Sprintf("Дарова мой самый любимый кентик %s, воспользуйся командой /menu, или нажми на меню справа внизу.", update.Message.From.UserName)

			case "/menu", "/menu@WallerBroBot":
				keyboards.AddStandartKeyboard(bot, chatID)
				
			case "phonk", "Phonk":
				replyMsg = fmt.Sprintf("А ты харош %s, я всегда знал что только ты мой единственный и настоящий друг", update.Message.From.UserName)
			
			default: 
				replyMsg = "Нет такой команды дурик, тебе же сказали /menu"
			
		}
			
		
			msg := tgbotapi.NewMessage(chatID, replyMsg)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

		logger.Printf("[%s] [%d]\n %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
			
		}
		
	}

}



