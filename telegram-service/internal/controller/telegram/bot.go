package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBot структура для бота
type TelegramBot struct {
	API      *tgbotapi.BotAPI
	Commands CommandHandler
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	telegramBot := &TelegramBot{
		API: bot,
	}

	// Назначение стандартного обработчика команд
	telegramBot.Commands = &TestCommandHandler{}

	return telegramBot, nil
}

// Run запускает обработку сообщений
func (tb *TelegramBot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tb.API.GetUpdatesChan(u)

	log.Println("Telegram bot is running...")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go tb.Commands.HandleCommand(update.Message)
	}
}
