package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	API      *tgbotapi.BotAPI
	Commands TelegramHandler
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	telegramBot := &TelegramBot{
		API: bot,
	}

	telegramBot.Commands = TelegramHandler{
		Bot: bot,
	}

	return telegramBot, nil
}

func (tb *TelegramBot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tb.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go tb.Commands.HandleCommand(update)
	}
}
